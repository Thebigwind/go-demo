package main

import (
	//"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"io/ioutil"
	"log"
)

//这个示例代码将一个 PEM 格式的证书文件加载到内存中，并将其发布到 OpenLDAP 服务器上的 ou=certificates,dc=example,dc=com 节点下。
//请注意，这个示例代码只处理单个证书，如果您需要批量发布证书，需要对代码进行修改。另外，您可能需要根据您的 LDAP 服务器设置调整代码中的服务器地址和绑定参数。
func publish() error {
	// 加载证书
	certData, err := ioutil.ReadFile("root.crt")
	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	//data := string(certData)[27:1149]
	// 解码证书
	block, _ := pem.Decode(certData)
	if block == nil {
		log.Printf("Failed to decode certificate")
		return err
	}

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	// 准备证书对象
	entry := ldap.NewAddRequest("cn="+cert.Subject.CommonName+",ou=certs,dc=zdlz,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson"})
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Attribute("sn", []string{"01"}) //cert.Subject.SerialNumber
	entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})

	l := GetLdapConn()
	// 发布证书
	err = l.Add(entry)
	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	fmt.Println("Certificate added successfully")
	return nil
}
