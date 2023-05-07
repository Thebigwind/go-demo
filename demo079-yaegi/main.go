package main

import (
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"os"
	"time"
)

func main() {
	// 创建一个解释器实例
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	// 加载并运行程序
	run(i)

	// 监听文件变化，如果发生了变化则重新加载程序
	for {
		time.Sleep(1 * time.Second)
		fileInfo, err := os.Stat("program.go")
		if err == nil {
			if fileInfo.ModTime().After(lastModTime) {
				lastModTime = fileInfo.ModTime()
				i.Reset()
				i.Use(stdlib.Symbols)
				run(i)
			}
		}
	}
}

var lastModTime time.Time

func run(i *interp.Interpreter) {
	// 加载并运行程序
	_, err := i.RunFile("program.go")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
