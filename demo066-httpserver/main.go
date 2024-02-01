package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// UpgradeRequest 结构定义了升级请求的参数
type UpgradeRequest struct {
	IP  string `json:"ip"`
	Dir string `json:"dir"`
}

// UpgradeResponse 结构定义了升级响应的参数
type UpgradeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SlaveUpgrade 是处理升级请求的处理器函数
func SlaveUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("xxxxxxxxxx")
		time.Sleep(time.Second)
		fmt.Println("oooooooooo")
	}()

	go asyncUpgrade()
	// 解析请求JSON数据
	var upgradeRequest UpgradeRequest
	err := json.NewDecoder(r.Body).Decode(&upgradeRequest)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("upgradeRequest:%+v", upgradeRequest)
	// 在实际应用中，你可以根据 upgradeRequest 的内容执行相应的升级逻辑
	// 这里简单返回一个成功的响应
	upgradeResponse := UpgradeResponse{
		Code: 0,
		Msg:  "Upgrade request received successfully",
	}

	// 将响应转换为JSON格式并发送给客户端
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(upgradeResponse)
}

func main() {
	// 注册处理器函数
	http.HandleFunc("/v1/slave_upgrade", SlaveUpgrade)

	// 启动HTTP服务器监听端口
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8888", nil)
}

func asyncUpgrade() {
	time.Sleep(time.Second * 10)
	fmt.Println("xxxxxxxxxxxxx")
	time.Sleep(time.Second * 2)
	fmt.Println("i'm last haha")
}
