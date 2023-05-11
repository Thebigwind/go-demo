package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
)

func main() {
	cert, key, err := getCertAndKey()
	if err != nil {
		fmt.Printf("err:%v", err.Error())
		return
	}
	p12Str, pass, err := certToP12(cert, key)
	if err != nil {
		fmt.Printf("err:%v", err.Error())
		return
	}
	fmt.Printf("p12str:%v,pass:%v\n", p12Str, pass)
	//

	p12Bytres, err := base64.StdEncoding.DecodeString(p12Str)
	if err != nil {
		fmt.Printf("DecodeString err:%v", err.Error())
		return
	}
	//decode
	certbase, certsn, certafter, err := getMDMCertCont(p12Bytres, pass)
	if err != nil {
		fmt.Printf("err:%v", err.Error())
		return
	}
	fmt.Printf("certbase:%s\n,certsn:%s\n,certafter:%v\n", certbase, certsn, certafter)
}

func getCertAndKey() ([]byte, []byte, error) {
	certBytes, err := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo108-p12/node.crt")
	if err != nil {
		return nil, nil, err
	}
	pemBytes, err := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo108-p12/node.key")
	if err != nil {
		return nil, nil, err
	}
	return certBytes, pemBytes, nil
}

//========
func createCertAndKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Failed to generate private key: %v\n", err)
		return nil, nil, err
	}

	// 将私钥编码为 PEM 格式
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pemBytes := pem.EncodeToMemory(pemBlock)

	// 生成 X.509 证书
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "example.com"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		fmt.Printf("Failed to create certificate: %v\n", err)
		return nil, nil, err
	}

	// 将证书编码为 PEM 格式
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	}
	certBytes := pem.EncodeToMemory(certBlock)

	return certBytes, pemBytes, nil
}

func Parse(p12Data []byte, password []byte) {
	// 读取 P12 文件内容
	//p12Data, err := ioutil.ReadFile("cert.p12")
	//if err != nil {
	//	fmt.Println("read P12 file error:", err)
	//	return
	//}

	// 解析 P12 文件
	p12, err := tls.X509KeyPair(p12Data, password)
	if err != nil {
		fmt.Println("parse P12 file error:", err)
		return
	}

	// 解析证书内容
	cert, err := x509.ParseCertificate(p12.Certificate[0])
	if err != nil {
		fmt.Println("parse certificate error:", err)
		return
	}

	fmt.Println("Certificate Subject:", cert.Subject)
	fmt.Println("Certificate Issuer:", cert.Issuer)
	fmt.Println("Certificate Serial Number:", cert.SerialNumber)
	fmt.Println("Certificate Not Before:", cert.NotBefore)
	fmt.Println("Certificate Not After:", cert.NotAfter)
}
