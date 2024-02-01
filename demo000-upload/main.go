package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const uploadPath = "./uploads/"

func main() {
	http.HandleFunc("/v1/ops/upgrade_package", uploadHandler)
	fmt.Println("Server listening on :8888...")
	http.ListenAndServe(":8888", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("token")
	fmt.Printf("token:%s\n", token)
	// 解析请求参数
	err := r.ParseMultipartForm(100 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 获取文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("handler.Filename:%s", handler.Filename)

	// 创建上传目录
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	// 创建文件
	dst, err := os.Create(uploadPath + handler.Filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将文件内容复制到目标文件
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"code": 200, "msg": "File uploaded successfully"}`))
}
