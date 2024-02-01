package main

import (
	"fmt"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	Cors1(w, r)
	fmt.Fprintf(w, "hello world,apiHandler")
}

func Cors1(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	fmt.Println("============Cors1========================", origin)
	if origin != "" {
		//接收客户端发送的origin （重要！）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		w.Header().Set("Access-Control-Allow-Methods", "*")
		//允许跨域设置可以返回其他子段，可以自定义字段
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Length,X-CSRF-Token,token,Token,session,Content-Type,content-type,Access-Token")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		w.Header().Set("Access-Control-Expose-Headers", "*")
		//设置缓存时间
		w.Header().Set("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		//content-type
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	}
}

func main() {
	fmt.Println("===================main 8003========================")
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8003", nil)
}
