package main

import (
	"fmt"
	"golang.org/x/crypto/sm4"
)

func test1() {
	// 定义密钥
	key := []byte("1234567890ABCDEF")
	// 定义要加密的数据
	data := []byte("身份证号码")
	// 创建一个SM4的加密器
	cipher, _ := sm4.NewCipher(key)
	// 加密数据
	encrypted := make([]byte, len(data))
	cipher.Encrypt(encrypted, data)
	// 输出加密后的数据
	fmt.Println(encrypted)
}
