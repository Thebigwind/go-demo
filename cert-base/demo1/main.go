package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func test() {
	// 读取 PEM 格式的证书文件
	certBytes, err := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/cert-base/demo1/example.crt")
	if err != nil {
		panic(err)
	}

	// 解析 PEM 格式的证书
	block, _ := pem.Decode(certBytes)
	if block == nil {
		panic("failed to decode PEM block")
	}

	// 解析 ASN.1 格式的证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 输出证书信息
	fmt.Printf("Subject: %s\n", cert.Subject)
	fmt.Printf("Issuer: %s\n", cert.Issuer)
	fmt.Printf("Serial Number: %v\n", cert.SerialNumber)
	fmt.Printf("Not Before: %v\n", cert.NotBefore)
	fmt.Printf("Not After: %v\n", cert.NotAfter)
	fmt.Printf("Signature Algorithm: %v\n", cert.SignatureAlgorithm)
	fmt.Printf("Public Key Algorithm: %v\n", cert.PublicKeyAlgorithm)
	fmt.Printf("Public Key: %v\n", cert.PublicKey)
}
