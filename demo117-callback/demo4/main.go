package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CallbackRequest 结构定义了回调请求的参数
type CallbackRequest struct {
	Version string `json:"version"`
	Status  int    `json:"status"`
}

// CallbackResponse 结构定义了回调响应的参数
type CallbackResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// upgradeSlave 函数模拟了Master服务调用Slave服务的升级接口
func upgradeSlave() (string, int) {
	// 模拟升级操作，返回升级结果
	// 这里简单起见，直接返回成功，实际情况根据业务逻辑处理
	return "1.0.1", 1
}

// handleUpgradeCallback 函数处理Slave的升级结果回调
func handleUpgradeCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 解析回调请求的JSON数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var callbackRequest CallbackRequest
	err = json.Unmarshal(body, &callbackRequest)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// 处理升级结果
	// 这里简单起见，直接输出回调信息，实际情况根据业务逻辑处理
	fmt.Printf("Received upgrade callback: Version %s, Status %d\n", callbackRequest.Version, callbackRequest.Status)

	// 构造回调响应
	callbackResponse := CallbackResponse{
		Code: 0,
		Msg:  "Callback processed successfully",
	}

	// 将响应转换为JSON格式并发送给Slave
	responseJSON, err := json.Marshal(callbackResponse)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	// 模拟Master服务调用Slave服务的升级接口
	slaveVersion, upgradeStatus := upgradeSlave()

	// 构造回调请求的JSON数据
	callbackRequest := CallbackRequest{
		Version: slaveVersion,
		Status:  upgradeStatus,
	}

	// 将回调请求的JSON数据转换为字节流
	callbackRequestBody, err := json.Marshal(callbackRequest)
	if err != nil {
		fmt.Println("Error encoding callback request body:", err)
		return
	}

	// 发送回调请求给Slave的回调接口
	resp, err := http.Post("http://slave-service-host:port/v1/slave_upgrade_callback", "application/json", bytes.NewBuffer(callbackRequestBody))
	if err != nil {
		fmt.Println("Error sending callback request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取并解析回调响应的JSON数据
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading callback response body:", err)
		return
	}

	var callbackResponse CallbackResponse
	err = json.Unmarshal(responseBody, &callbackResponse)
	if err != nil {
		fmt.Println("Error decoding callback response body:", err)
		return
	}

	// 输出回调响应信息
	fmt.Printf("Callback Response: Code %d, Msg %s\n", callbackResponse.Code, callbackResponse.Msg)
}
