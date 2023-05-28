package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	// 连接 OpenLDAP 服务器
	l, err := ldap.Dial("tcp", "10.10.10.125:389")
	if err != nil {
		fmt.Println("连接 OpenLDAP 服务器失败：", err)
		return
	}
	defer l.Close()

	// 绑定管理员账户
	err = l.Bind("cn=admin,dc=zdlz,dc=com", "123456")
	if err != nil {
		fmt.Println("绑定管理员账户失败：", err)
		return
	}

	// 准备要发布的用户信息
	dn := "uid=johndoe1,ou=cert,dc=zdlz,dc=com"

	entry := ldap.NewAddRequest(dn, nil)
	entry.Attribute("objectClass", []string{"top", "certificationAuthority", "pkiCA"}) // "inetOrgPerson",
	entry.Attribute("cn", []string{"John Doe1"})
	entry.Attribute("sn", []string{"Doe1"})
	entry.Attribute("givenName", []string{"John1"})
	entry.Attribute("uid", []string{"johndoe1"})
	entry.Attribute("userPassword", []string{"{SSHA}i9y43hf8ygrf49y2h38fh298g"})
	entry.Attribute("mail", []string{"johndoe1@example.com"})

	// 发布用户信息
	//entry := ldap.NewEntry(dn, attributes)
	err = l.Add(entry)
	if err != nil {
		fmt.Println("发布用户信息失败：", err)
		return
	}

	fmt.Println("用户信息已成功发布到 OpenLDAP。")
}

func user() {
	// 连接 OpenLDAP 服务器
	l, err := ldap.Dial("tcp", "10.10.10.125:389")
	if err != nil {
		fmt.Println("连接 OpenLDAP 服务器失败：", err)
		return
	}
	defer l.Close()

	// 绑定管理员账户
	err = l.Bind("cn=admin,dc=zdlz,dc=com", "123456")
	if err != nil {
		fmt.Println("绑定管理员账户失败：", err)
		return
	}

	// 准备要发布的用户信息
	dn := "uid=johndoe1,ou=people,dc=zdlz,dc=com"

	entry := ldap.NewAddRequest(dn, nil)
	entry.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"}) // "inetOrgPerson",
	entry.Attribute("cn", []string{"John Doe1"})
	entry.Attribute("sn", []string{"Doe1"})
	entry.Attribute("givenName", []string{"John1"})
	entry.Attribute("uid", []string{"johndoe1"})
	entry.Attribute("userPassword", []string{"{SSHA}i9y43hf8ygrf49y2h38fh298g"})
	entry.Attribute("mail", []string{"johndoe1@example.com"})

	// 发布用户信息
	//entry := ldap.NewEntry(dn, attributes)
	err = l.Add(entry)
	if err != nil {
		fmt.Println("发布用户信息失败：", err)
		return
	}

	fmt.Println("用户信息已成功发布到 OpenLDAP。")
}

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
