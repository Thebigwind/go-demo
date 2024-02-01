package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个通道用于接收信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动一些工作
	go doSomeWork()

	// 等待信号
	sigReceived := <-sigChan
	fmt.Printf("Received signal: %v\n", sigReceived)

	// 执行清理操作
	cleanup()

	// 优雅退出
	os.Exit(0)
}

func doSomeWork() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("Working...")
		}
	}
}

func cleanup() {
	// 在这里执行清理操作，例如关闭数据库连接、保存状态等
	fmt.Println("Cleaning up...")
}
