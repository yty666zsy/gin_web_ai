## 我们将会使用go语言的gin库来搭建一个属于自己的网页版GPT

### 一、准备工作

我们需要使用到ollama，如何下载和使用[ollama]([Ollama完整教程：本地LLM管理、WebUI对话、Python/Java客户端API应用 - 老牛啊 - 博客园](https://www.cnblogs.com/obullxl/p/18295202/NTopic2024071001))请看这个文档

有过gin环境的直接运行就可以，如果没有就根据文档内容去下载相关配置库

### 二、使用步骤

```shell
git clone https://github.com/yty666zsy/gin_web_ai.git
cd gin_web_ai
ollama run "大模型的名称"
```

这里需要注意的是要在chat.html文件中修改模型的名称，要不然找不到模型，在这个位置![image-20241126164437662](C:\Users\yuzai\AppData\Roaming\Typora\typora-user-images\image-20241126164437662.png)

然后运行代码，如下图所示

![image-20241126164139716](https://github.com/yty666zsy/gin_web_ai/raw/master/image/image-20241126164437662.png)

```shell
"然后开启一个新的终端"
go run main.go
```

这里需要注意的是端口号可以适当的进行修改，防止某些端口被占用的情况

![image-20241126164139716](https://github.com/yty666zsy/gin_web_ai/raw/master/image/image-20241126164234439.png)

然后本地访问127.0.0.1:8088就能打开网址进行愉快的聊天啦

![image-20241126164139716](https://github.com/yty666zsy/gin_web_ai/raw/master/image/image-20241126164631039.png)

同时后台同步打印信息以便日志管理

![image-20241126164139716](https://github.com/yty666zsy/gin_web_ai/raw/master/image/image-20241126164717202.png)
