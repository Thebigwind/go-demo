package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

//hmac

func GenerateHMAC(text string, key string) string {

	textBytes := []byte(text) //源数据
	keyBytes := []byte(key)   //加密密钥

	hash := hmac.New(sha256.New, keyBytes)

	hash.Write(textBytes)

	result := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
	//WQXWaNzwM92rDABEsmmFsWP2W7GCuou+civrsJeued0=
}

func VerifyHMAC(HMAC string, text string, key string) bool {

	HMACBytes := []byte(HMAC)

	nowHMAC := GenerateHMAC(text, key)
	nowHMACBytes := []byte(nowHMAC)

	return hmac.Equal(HMACBytes, nowHMACBytes)
}

func main() {
	a := GenerateHMAC("123456", "xxxxxx")
	fmt.Println(a)

	result := VerifyHMAC("WQXWaNzwM92rDABEsmmFsWP2W7GCuou+civrsJeued0=", "123456", "xxxxxx")
	fmt.Println(result)
}
