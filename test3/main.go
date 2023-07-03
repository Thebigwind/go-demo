package main

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
)

func main() {
	pub := `-----BEGIN PUBLIC KEY-----
	MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEz/F/5r5xdI3b4KuHj3u//Dnd+OIA
	jO2oMFer6qr1Ikpe/+J3cbyeT4YPUKzljvlcWpY4PfhzDA5MqGlCSzwn0A==
	-----END PUBLIC KEY-----`
	// 公钥字节切片
	pubKeyBytes := []byte(pub)

	// 解析公钥
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		return
	}

	// 将公钥转换为DER编码
	derBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		fmt.Println("Error converting public key to DER:", err)
		return
	}

	// 将DER编码的公钥转换为不带04前缀的HEX格式
	hexBytes := hex.EncodeToString(derBytes)[2:]
	fmt.Println(hexBytes)
}
