package main

import (
	"fmt"
	"net/http"
	"time"
)

// SSE 处理函数，向客户端推送事件
func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 Content-Type 为 text/event-stream，表示服务器推送事件
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 使用长期连接向客户端发送消息
	for i := 0; ; i++ {
		// 构造每个消息
		message := fmt.Sprintf("data: Message %d from server\n\n", i)
		// 写入消息
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
		// 强制刷新缓冲区，确保实时推送
		flusher, ok := w.(http.Flusher)
		if ok {
			flusher.Flush()
		}

		// 每 2 秒发送一个新消息
		time.Sleep(2 * time.Second)
	}
}

// HTML 和 JavaScript 嵌入到 Go 代码中
const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OpenAI SSE Example</title>
    <style>
        #messages {
            margin-top: 20px;
            font-family: Arial, sans-serif;
        }
    </style>
</head>
<body>

    <h1>OpenAI SSE Example</h1>
    <div id="messages"></div>

    <script>
        // 创建一个 EventSource 实例，连接到服务器的 /events 路径
        const eventSource = new EventSource('/events');

        // 监听 'message' 事件，接收来自服务器的每条消息
        eventSource.onmessage = function(event) {
            const messagesDiv = document.getElementById('messages');
            // 显示收到的消息
            const newMessage = document.createElement('p');
            newMessage.textContent = event.data;
            messagesDiv.appendChild(newMessage);
        };

        // 错误处理
        eventSource.onerror = function(event) {
            console.error('Error occurred:', event);
        };
    </script>

</body>
</html>
`

// 主页处理函数，返回嵌入的 HTML 内容
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

func mainSSE() {
	// 设置路由
	http.HandleFunc("/", indexHandler)     // 根路由提供 HTML 页面
	http.HandleFunc("/events", sseHandler) // SSE 路由

	// 启动服务器
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
