package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
)

//我们首先将证书内容解码为 PEM 编码的数据块，然后使用 `x509.ParseCertificate()` 函数将 PEM 编码的数据块解析为 `x509.Certificate` 结构体。
//最后，我们输出了该证书的一些基本信息，包括主题、颁发者、有效期、序列号、签名算法、公钥算法以及公钥。
func certStrToCertStruct() {
	// 假设证书内容已经读取到了字符串变量 certStr 中
	//	certStr := `
	//-----BEGIN CERTIFICATE-----
	//MIIC6DCCAdCgAwIBAgIUDVfmEHz/6U/2gUfrj6UHv6cRfjQwCgYIKoEcz1UBg3Uw
	//czELMAkGA1UEBhMCQ04xEDAOBgNVBAgMB1p1cmljaDETMBEGA1UEBwwKSG9ja2lu
	//d2F5MRIwEAYDVQQKDAlDTElFTlRBTCBOMRowGAYDVQQLDBFDTkVkZHJlc3MgQ0Eg
	//MzIwNjE4ODMxMDAwMRYwFAYDVQQDDA1sb2NhbGhvc3Q6ODA4MB4XDTE5MDQwMjIz
	//NTg0NVoXDTI5MDQwMjIzNTg0NVowfTELMAkGA1UEBhMCQ04xEDAOBgNVBAgMB1p1
	//cmljaDETMBEGA1UEBwwKSG9ja2lud2F5MRIwEAYDVQQKDAlDTElFTlRBTCBOMRow
	//GAYDVQQLDBFDTkVkZHJlc3MgQ0EgMzIwNjE4ODMxMDAwMRYwFAYDVQQDDA1sb2Nh
	//bGhvc3Q6ODA4MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEV7DFdDdjfuYXe+nO
	//jGrDv7zMhHiJlxlV7aQZrjKdqSDI+oA4G4f+xnpQsM4T/v1bU23KjU0w6JxOG6A
	//jKOBqzCBqDAfBgNVHSMEGDAWgBQa+gfvLZDxHIBPMRrz8JvzELEjWTAKBggqhkjO
	//PQQDAgNIADBFAiBj+L5x/SW8zoKFJn/Hn+FiD+W7jusgE8Nv7c25r/j+LQIhAKW
	//uEWdeV0rh0o/dy7XDJjGMeu8OXoMfoMw0Z7VfNgW
	//-----END CERTIFICATE-----`

	certStr, _ := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo107-sm2-protect/node.crt")

	// 将证书内容解码为 PEM 编码的数据块
	block, _ := pem.Decode([]byte(certStr))
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("failed to parse certificate")
		return
	}

	// 将 PEM 编码的数据块解析为 x509.Certificate 结构体
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse certificate:", err)
		return
	}

	// 输出证书信息
	fmt.Println("Subject:", cert.Subject.CommonName)
	fmt.Println("Issuer:", cert.Issuer.CommonName)
	fmt.Println("Not Before:", cert.NotBefore)
	fmt.Println("Not After:", cert.NotAfter)
	fmt.Println("Serial Number:", cert.SerialNumber)
	fmt.Println("Signature Algorithm:", cert.SignatureAlgorithm)
	fmt.Println("Public Key Algorithm:", cert.PublicKeyAlgorithm)
	fmt.Println("Public Key:", cert.PublicKey)
}

//我们首先将私钥字符串解码为 PEM 编码的数据块，然后使用 x509.ParseECPrivateKey() 函数将 PEM 编码的数据块解析为 sm2.PrivateKey 结构体。
//最后，我们输出了该私钥的 D 值。需要注意的是，在解析私钥之前，需要先将私钥字符串解码为 PEM 编码的数据块，否则会出现解析失败的情况
func keyStrToKeyStruct() {

	// SM2 私钥字符串
	//priKeyStr := "308193020100301306072a8648ce3d020106082a811ccf5501822d047930770201010420c11d82b7a10ce6039f92a285556a44d1a06c6e13a5a7fbc39ec055fe7d19198a00a06082a811ccf5501822da14403420004c69455b5a5b0a102d5d29f2e3a3b0c4c90f0a5b5f031b4303e2153d6ab8f27d7fbce47806df6e61d2e1e305d7a2c1b05f3920d8a05572055d47200efc44f15da97"

	priKeyStr, _ := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo107-sm2-protect/node.key")
	// 将私钥字符串解码为 PEM 编码的数据块
	block, _ := pem.Decode([]byte(priKeyStr))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	// 使用 x509.ParseECPrivateKey() 函数将 PEM 编码的数据块解析为 sm2.PrivateKey 结构体
	priKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}
	// 检查是否为 sm2 椭圆曲线
	//if priKey.Curve != elliptic.Curve {
	//	panic("invalid private key curve")
	//}
	// 打印私钥信息
	fmt.Println("Private Key:", priKey.D)

}

func pubkeyStrToKeyStruct() {

	// SM2 私钥字符串
	pubKeyStr := "3059301306072a8648ce3d020106082a811ccf5501822d03420004c69455b5a5b0a102d5d29f2e3a3b0c4c90f0a5b5f031b4303e2153d6ab8f27d7fbce47806df6e61d2e1e305d7a2c1b05f3920d8a05572055d47200efc44f15da97"
	// 将私钥字符串解码为 PEM 编码的数据块
	block, _ := pem.Decode([]byte(pubKeyStr))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	// 使用 x509.ParsePKIXPublicKey() 函数将 PEM 编码的数据块解析为 sm2.PublicKey 结构体
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}
	// 转换为 sm2.PublicKey 结构体
	sm2PubKey, ok := pubKey.(*sm2.PublicKey)
	if !ok {
		panic("invalid public key type")
	}
	// 打印公钥信息
	fmt.Println("Public Key X:", sm2PubKey.X)
	fmt.Println("Public Key Y:", sm2PubKey.Y)
}

/*
import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sm2"
	"fmt"
	"io/ioutil"
)

func main() {
	// 生成随机密钥对
	key, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("generate key error:", err)
		return
	}

	// 加密数据
	plaintext := []byte("Hello, world!")
	ciphertext, err := sm2.Encrypt(rand.Reader, &key.PublicKey, plaintext, nil)
	if err != nil {
		fmt.Println("encrypt error:", err)
		return
	}

	// 解密数据
	decrypted, err := key.Decrypt(ciphertext, nil)
	if err != nil {
		fmt.Println("decrypt error:", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)

	// 保存密钥对到文件
	err = ioutil.WriteFile("private.key", key.D.Bytes(), 0644)
	if err != nil {
		fmt.Println("write private key error:", err)
		return
	}
	err = ioutil.WriteFile("public.key", elliptic.Marshal(&key.PublicKey.Curve, key.PublicKey.X, key.PublicKey.Y), 0644)
	if err != nil {
		fmt.Println("write public key error:", err)
		return
	}
}
*/
