package main

//https://juejin.cn/post/7176220436714225721

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateAppKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

func generateAppSecret() string {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(secret)
}

func main() {
	appKey := generateAppKey()
	appSecret := generateAppSecret()

	fmt.Println("App Key:", appKey)
	fmt.Println("App Secret:", appSecret)

	// Use the app key and secret in your application
	// ...
}
