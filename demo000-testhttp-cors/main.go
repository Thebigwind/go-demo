package main

import (
	"fmt"
	"net/http"
)

// 定义HTML网页字符串常量。
const html = `
<!DOCTYPE html> 
<html lang="en"> 
<head>
    <meta charset="UTF-8"> 
</head>
<body>
    <h1>Simple CORS</h1> 
    <div id="output"></div> 
    <script>
        document.addEventListener('DOMContentLoaded', function() { 
            fetch("http://localhost:8003/api").then(
                function (response) { 
                    response.text().then(function (text) {
                        document.getElementById("output").innerHTML = text; 
                    });
                }, 
                function(err) {
                    document.getElementById("output").innerHTML = err; 
                }
            ); 
        });
    </script> 
</body>
</html>`

// https://pkg.go.dev/github.com/tmc/scp?utm_source=godoc

type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Cors(w, r)
	//fmt.Fprintf(w, ih.content)
	w.Write([]byte(html))
}

func main() {
	fmt.Println("===================main========================")
	http.Handle("/", &indexHandler{content: "hello world!"})
	http.ListenAndServe(":8001", nil)
}

func Cors(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	fmt.Println("====================================", origin)
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
