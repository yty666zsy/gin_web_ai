package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"bufio"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Message Message `json:"message"`
}

// 添加新的结构体用于处理流式响应
type StreamResponse struct {
	Model      string  `json:"model"`
	CreatedAt  string  `json:"created_at"`
	Message    Message `json:"message"`
	Done       bool    `json:"done"`
	DoneReason string  `json:"done_reason,omitempty"`
}

// 在main函数前添加一个新的函数
func findAvailablePort(startPort int) int {
	for port := startPort; port < startPort+100; port++ {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			listener.Close()
			return port
		}
	}
	return startPort // 如果没找到可用端口，返回初始端口
}

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("templates/*")

	// 设置静态文件路径
	r.Static("/static", "./static")

	// 首页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	// 处理聊天请求的API
	r.POST("/chat", func(c *gin.Context) {
		var req ChatRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		prettyJSON, _ := json.MarshalIndent(req, "", "    ")
		fmt.Printf("发送到Ollama的请求:\n%s\n", string(prettyJSON))

		jsonData, _ := json.Marshal(req)
		resp, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("调用Ollama API错误: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// 使用scanner来读取流式响应
		scanner := bufio.NewScanner(resp.Body)
		var fullContent string

		for scanner.Scan() {
			line := scanner.Text()
			var streamResp StreamResponse
			if err := json.Unmarshal([]byte(line), &streamResp); err != nil {
				fmt.Printf("解析流式响应行错误: %v\n", err)
				continue
			}

			// 累积内容
			fullContent += streamResp.Message.Content

			// 如果是最后一条消息
			if streamResp.Done {
				response := ChatResponse{
					Message: Message{
						Role:    "assistant",
						Content: fullContent,
					},
				}
				c.JSON(http.StatusOK, response)
				return
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("读取流式响应错误: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	// 从环境变量获取端口，如果未设置则使用默认值
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	r.Run(":" + port)
}
