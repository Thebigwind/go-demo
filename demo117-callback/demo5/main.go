package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// CallbackRequest 结构定义了回调请求的参数
type CallbackRequest struct {
	Message string `json:"message"`
}

// upgradeCallbackHandler 处理回调请求的处理器
func upgradeCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 解析回调请求的JSON数据
	// 在实际应用中，你可能需要根据实际情况定义更复杂的结构体
	var callbackRequest CallbackRequest
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing request form", http.StatusBadRequest)
		return
	}

	callbackRequest.Message = r.FormValue("message")

	// 处理回调请求
	// 在实际应用中，这里应该根据回调内容执行相应的逻辑
	fmt.Printf("Received callback: %s\n", callbackRequest.Message)

	// 返回回调响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Callback processed successfully"))
}

// asyncTask 模拟异步任务
func asyncTask(callbackURL string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟异步任务执行
	time.Sleep(5 * time.Second)

	// 构造回调请求的参数
	callbackRequest := CallbackRequest{
		Message: "Async task completed",
	}
	fmt.Printf("%v", callbackRequest)

	// 发送回调请求给指定的回调接口
	resp, err := http.Post(callbackURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error sending callback request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取并输出回调响应
	body, err := http.ReadResponse(resp, nil)
	if err != nil {
		fmt.Println("Error reading callback response:", err)
		return
	}

	fmt.Printf("Callback Response: %s\n", body.Status)
}

// main 函数启动 HTTP 服务器
func main() {
	// 启动 HTTP 服务器，处理异步任务
	http.HandleFunc("/upgrade_callback", upgradeCallbackHandler)
	go http.ListenAndServe(":8080", nil)

	// 处理 RESTful 接口请求
	http.HandleFunc("/process_request", func(w http.ResponseWriter, r *http.Request) {
		// 立即返回 200 OK 给客户端
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request received, processing asynchronously"))

		// 启动异步任务
		var wg sync.WaitGroup
		wg.Add(1)
		go asyncTask("http://localhost:8080/upgrade_callback", &wg)

		// 在这里你可以继续处理请求的其他逻辑
		// ...

		// 等待异步任务完成
		wg.Wait()
	})

	// 启动 HTTP 服务器监听端口
	fmt.Println("Server listening on :8081")
	http.ListenAndServe(":8081", nil)
}
