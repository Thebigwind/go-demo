package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	// LDAP 服务器地址和端口
	ldapURL := "ldap://10.10.10.125:389"
	//server := "10.10.10.125:389"
	bindDN := "cn=admin,dc=xyz,dc=com"
	bindPassword := "123456"

	// 要发布的证书
	cert := `-----BEGIN CERTIFICATE-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr3f6lV3UjN8p3IUKFSg0
...
-----END CERTIFICATE-----`

	// 创建 LDAP 连接
	l, err := ldap.DialURL(ldapURL)
	//l, err := ldap.Dial("tcp", server)
	if err != nil {
		log.Fatalf("Failed to connect to LDAP server: %v", err)
	}
	defer l.Close()

	// 使用 TLS
	//err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	//if err != nil {
	//	log.Fatalf("Failed to start TLS: %v", err)
	//}

	// 绑定到 LDAP
	err = l.Bind(bindDN, bindPassword)
	if err != nil {
		log.Fatalf("Failed to bind to LDAP server: %v", err)
	}

	// 创建或更新条目
	dn := "cn=exampleUser,dc=example,dc=com"
	addReq := ldap.NewAddRequest(dn, nil)
	addReq.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute("cn", []string{"exampleUser"})
	addReq.Attribute("sn", []string{"User"})
	addReq.Attribute("userCertificate;binary", []string{cert})

	err = l.Add(addReq)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultEntryAlreadyExists) {
			// 条目已存在，则更新证书
			modifyReq := ldap.NewModifyRequest(dn, nil)
			modifyReq.Replace("userCertificate;binary", []string{cert})

			err = l.Modify(modifyReq)
			if err != nil {
				log.Fatalf("Failed to modify entry: %v", err)
			} else {
				fmt.Println("Certificate updated successfully")
			}
		} else {
			log.Fatalf("Failed to add entry: %v", err)
		}
	} else {
		fmt.Println("Entry added successfully")
	}
}
