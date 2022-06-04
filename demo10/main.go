package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 在这里，我把桶设置成了每一毫秒投放一次令牌，桶容量大小为 10，起一个 http 的服务，模拟后台 API。
	r := Every(1 * time.Millisecond)
	limit := NewLimiter(r, 10)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if limit.Allow() {
			fmt.Printf("请求成功，当前时间：%s\n", time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("请求成功，但是被限流了。。。\n")
		}
	})

	_ = http.ListenAndServe(":8081", nil)
}
