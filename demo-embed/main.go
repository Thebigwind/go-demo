package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed my_cert.pm
var fs embed.FS

func main() {
	// 读取嵌入的文件内容
	data, err := fs.ReadFile("my_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("证书内容：", string(data))
}
