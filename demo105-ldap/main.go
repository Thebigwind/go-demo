package main

//
//import (
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"crypto/x509/pkix"
//	"encoding/pem"
//	"fmt"
//	"golang.org/x/crypto/ocsp"
//	"io/ioutil"
//	"math/big"
//	"net/http"
//	"time"
//)
//
//func main() {
//	// 读取证书和颁发者证书
//	certFile, _ := ioutil.ReadFile("path/to/cert.pem")
//	certBlock, _ := pem.Decode(certFile)
//	cert, _ := x509.ParseCertificate(certBlock.Bytes)
//
//	issuerCertFile, _ := ioutil.ReadFile("path/to/issuer_cert.pem")
//	issuerCertBlock, _ := pem.Decode(issuerCertFile)
//	issuerCert, _ := x509.ParseCertificate(issuerCertBlock.Bytes)
//
//	// 构建OCSP请求
//	ocspRequest, _ := ocsp.CreateRequest(cert, issuerCert, nil)
//
//	// 发送OCSP请求
//	ocspURL := "http://ocsp.example.com/"
//	client := &http.Client{Timeout: time.Second * 10}
//	resp, err := client.Post(ocspURL, "application/ocsp-request", ocspRequest)
//	if err != nil {
//		panic(err)
//	}
//
//	// 解析OCSP响应
//	ocspBytes, _ := ioutil.ReadAll(resp.Body)
//	ocspResponse, _ := ocsp.ParseResponse(ocspBytes, issuerCert)
//
//	// 验证OCSP响应
//	ocspResponse.Check(cert, issuerCert, time.Now())
//
//	// 检查证书状态
//	if ocspResponse.Status != ocsp.Good {
//		panic("Certificate is not valid")
//	}
//
//	fmt.Println("Certificate is valid")
//}
//
//// 生成测试证书
//func generateTestCert() (*x509.Certificate, *rsa.PrivateKey, error) {
//	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
//
//	template := x509.Certificate{
//		SerialNumber: big.NewInt(1),
//		Subject: pkix.Name{
//			CommonName: "example.com",
//		},
//		NotBefore: time.Now(),
//		NotAfter:  time.Now().Add(time.Hour * 24 * 365),
//		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
//		ExtKeyUsage: []x509.ExtKeyUsage{
//			x509.ExtKeyUsageServerAuth,
//			x509.ExtKeyUsageClientAuth,
//		},
//		BasicConstraintsValid: true,
//		IsCA:                  false,
//	}
//
//	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	cert, err := x509.ParseCertificate(certBytes)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return cert, privateKey, nil
//}
