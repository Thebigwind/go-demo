package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	// LDAP服务器地址和端口
	server := "10.10.10.125:389"

	// 连接LDAP服务器
	conn, err := ldap.Dial("tcp", server)
	if err != nil {
		log.Fatalf("无法连接到LDAP服务器: %s", err)
	}
	defer conn.Close()

	// 绑定到LDAP服务器
	err = conn.Bind("cn=admin,dc=zdlz,dc=com", "123456")
	if err != nil {
		log.Fatalf("LDAP绑定失败: %s", err)
	}

	//certBytes, err := ioutil.ReadFile("./demo105-ldap/root.crt")
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//
	//block, _ := pem.Decode(certBytes)
	//if block == nil {
	//	fmt.Println("Failed to decode certificate")
	//	return
	//}
	//
	//cert, err := x509.ParseCertificate(block.Bytes)
	//
	//if err != nil {
	//	log.Printf("err:%s", err)
	//	return
	//}
	fmt.Println("-----------")
	entry := ldap.NewAddRequest("uid=root,ou=certs,dc=zdlz,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson", "person", "pkiUser"})
	entry.Attribute("cn", []string{"cn"})
	entry.Attribute("sn", []string{"01"})
	//entry.Attribute("userCertificate;binary", []string{string(block.Bytes)})
	//entry.Attribute("userCertificate;binary", []string{string(cert.Raw)})
	data := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"
	entry.Attribute("userCertificate;binary", []string{data})

	err = conn.Add(entry)
	if err != nil {
		fmt.Println("Error adding entry:", err)
		return
	}

	fmt.Println("Certificate added successfully.")
}
func user5() {
	// LDAP服务器地址和端口
	server := "10.10.10.125:389"

	// 连接LDAP服务器
	conn, err := ldap.Dial("tcp", server)
	if err != nil {
		log.Fatalf("无法连接到LDAP服务器: %s", err)
	}
	defer conn.Close()

	// 绑定到LDAP服务器
	err = conn.Bind("cn=admin,dc=zdlz,dc=com", "123456")
	if err != nil {
		log.Fatalf("LDAP绑定失败: %s", err)
	}

	//certBytes, err := ioutil.ReadFile("./demo105-ldap/root.crt")
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//
	//block, _ := pem.Decode(certBytes)
	//if block == nil {
	//	fmt.Println("Failed to decode certificate")
	//	return
	//}
	//
	//cert, err := x509.ParseCertificate(block.Bytes)
	//
	//if err != nil {
	//	log.Printf("err:%s", err)
	//	return
	//}
	fmt.Println("-----------")
	entry := ldap.NewAddRequest("cn=root,ou=certs,dc=zdlz,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson", "person", "pkiUser"})
	entry.Attribute("cn", []string{"cn"})
	entry.Attribute("sn", []string{"01"})
	//entry.Attribute("userCertificate;binary", []string{string(block.Bytes)})
	//entry.Attribute("userCertificate;binary", []string{string(cert.Raw)})
	data := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"
	entry.Attribute("userCertificate;binary", []string{data})

	err = conn.Add(entry)
	if err != nil {
		fmt.Println("Error adding entry:", err)
		return
	}

	fmt.Println("Certificate added successfully.")
}

func user3() {
	// LDAP服务器地址和端口
	server := "10.10.10.125:389"

	// 连接LDAP服务器
	conn, err := ldap.Dial("tcp", server)
	if err != nil {
		log.Fatalf("无法连接到LDAP服务器: %s", err)
	}
	defer conn.Close()

	// 绑定到LDAP服务器
	err = conn.Bind("cn=admin,dc=zdlz,dc=com", "123456")
	if err != nil {
		log.Fatalf("LDAP绑定失败: %s", err)
	}

	// 加载要发布的证书文件
	//cert, err := tls.LoadX509KeyPair("/Users/me/Thebigwind/go-demo/demo105-ldap/root.crt", "/Users/me/Thebigwind/go-demo/demo105-ldap/root.key")
	//if err != nil {
	//	log.Fatalf("无法加载证书: %s", err)
	//}
	// 将证书内容进行base64编码
	//certData := base64.StdEncoding.EncodeToString(cert.Certificate[0])

	// 加载要发布的证书文件
	//cert, err := x509.ParseCertificate([]byte("root.crt"))
	//if err != nil {
	//	log.Fatalf("无法加载证书: %s", err)
	//}

	// 创建LDAP条目
	entry := ldap.NewAddRequest("cn=mycert,ou=certificates,dc=example,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "person"})
	entry.Attribute("cn", []string{"mycert"})
	//entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	data := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"
	entry.Attribute("userCertificate;binary", []string{data})
	// 发布证书到LDAP
	err = conn.Add(entry)
	if err != nil {
		log.Fatalf("证书发布失败: %s", err)
	}

	fmt.Println("证书已成功发布到LDAP服务器")
}

func user2() {
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
