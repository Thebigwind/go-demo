package main

import _ "github.com/zeromicro/go-zero/core/stat"

func main() {
	select {}
}

/*
import (
	//"encoding/asn1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/go-ldap/ldap/v3"

func main() {
	// 读取数字证书文件
	certBytes, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析数字证书
	certBlock, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing certificate:", err)
		return
	}

	// 将数字证书转换为LDAP条目
	entry := ldap.NewEntry(fmt.Sprintf("cn=%s,ou=certificates,o=example", cert.Subject.CommonName), []ldap.Attribute{
		ldap.Attribute{
			Type: "objectClass",
			Vals: []string{"top", "inetOrgPerson"},
		},
		ldap.Attribute{
			Type: "cn",
			Vals: []string{cert.Subject.CommonName},
		},
		ldap.Attribute{
			Type: "userCertificate;binary",
			Vals: []string{base64.StdEncoding.EncodeToString(cert.Raw)},
		},
	})

	// 将LDAP条目发布到OpenLDAP服务器
	conn, err := ldap.Dial("tcp", "ldap.example.com:389")
	if err != nil {
		fmt.Println("Error connecting to LDAP server:", err)
		return
	}
	defer conn.Close()

	err = conn.Bind("cn=admin,o=example", "password")
	if err != nil {
		fmt.Println("Error binding to LDAP server:", err)
		return
	}

	err = conn.Add(entry)
	if err != nil {
		fmt.Println("Error adding entry:", err)
		return
	}

	fmt.Println("Certificate added successfully.")
}
*/
