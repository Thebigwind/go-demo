package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// 加密私钥
	password := []byte("my-password")
	encryptedPrivateKey, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(privateKey), password, x509.PEMCipherAES256)
	if err != nil {
		panic(err)
	}

	// 将加密后的私钥保存到文件
	pemFile, err := os.Create("private_key.pem")
	if err != nil {
		panic(err)
	}
	defer pemFile.Close()

	err = pem.Encode(pemFile, encryptedPrivateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Private key generated and encrypted successfully!")
}
