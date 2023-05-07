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

	data := string(certData)[27:1149]
	fmt.Println(data)
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
	fmt.Println("------------")
	fmt.Printf("subject:%+v\n", cert.Subject)
	fmt.Printf("Issuer:%+v\n", cert.Issuer)
	//fmt.Printf("Issuer raw:%+v\n", string(cert.RawIssuer))
	fmt.Printf("SerialNumber:%+v\n", cert.SerialNumber)
	fmt.Printf("Raw:%+v\n", cert.Raw)
	fmt.Printf("NotBefore:%+v\n", cert.NotBefore)
	fmt.Printf("NotAfter:%+v\n", cert.NotAfter)
	fmt.Printf("Extensions:%+v\n", cert.Extensions)
	fmt.Println("------------")
	// 准备证书对象
	entry := ldap.NewAddRequest("cn="+cert.Subject.CommonName+",ou=certs,dc=zdlz,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson"})
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Attribute("sn", []string{"01"}) //cert.Subject.SerialNumber
	//entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})

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

func publish2() error {
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
	entry.Attribute("objectClass", []string{"top", "pkiCA"}) //ipsecCert
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Attribute("sn", []string{"01"}) //cert.Subject.SerialNumber
	entry.Attribute("userCertificate:", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})

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

func publishmodify() error {
	// 加载证书
	certData, err := ioutil.ReadFile("root.crt")
	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	data := string(certData)[27:1149]
	fmt.Println(data)
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
	fmt.Println("------------")
	fmt.Printf("subject:%+v\n", cert.Subject)
	fmt.Printf("Issuer:%+v\n", cert.Issuer)
	//fmt.Printf("Issuer raw:%+v\n", string(cert.RawIssuer))
	fmt.Printf("SerialNumber:%+v\n", cert.SerialNumber)
	fmt.Printf("Raw:%+v\n", cert.Raw)
	fmt.Printf("NotBefore:%+v\n", cert.NotBefore)
	fmt.Printf("NotAfter:%+v\n", cert.NotAfter)
	fmt.Printf("Extensions:%+v\n", cert.Extensions)
	fmt.Printf("rawcert:%+v\n", base64.StdEncoding.EncodeToString(cert.Raw))
	fmt.Println("------------")
	// 准备证书对象
	entry := ldap.NewModifyRequest("cn="+cert.Subject.CommonName+",ou=certs,dc=zdlz,dc=com", nil)
	entry.Replace("objectClass", []string{"top", "inetOrgPerson"})
	entry.Replace("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Replace("sn", []string{"01"}) //cert.Subject.SerialNumber
	//entry.Replace("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	entry.Replace("userCertificate;binary", []string{string(cert.Raw)})

	l := GetLdapConn()
	// 发布证书
	err = l.Modify(entry)
	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	fmt.Println("Certificate added successfully")
	return nil
}

func pubCRL() {
	// 加载CRL文件
	crlBytes, err := ioutil.ReadFile("crl.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将CRL编码为base64字符串
	crlBase64 := base64.StdEncoding.EncodeToString(crlBytes)

	// 创建CRL条目
	entry := ldap.NewAddRequest("cn=crl,dc=zdlz,dc=com", []ldap.Control{})
	entry.Attribute("objectClass", []string{"top", "cRLDistributionPoint"})
	entry.Attribute("crlDistributionPoint", []string{"ldap://10.10.10.125:389/cn=crl,dc=zdlz,dc=com?certificateRevocationList?base?(objectClass=cRLDistributionPoint)"})
	entry.Attribute("certificateRevocationList", []string{crlBase64})

	// 发布CRL条目
	err = GetLdapConn().Add(entry)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CRL已成功发布")
}

/*
func search() {
	// 设置连接参数
	l, err := ldap.DialTLS("tcp", "ldap.example.com:636", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// 绑定
	err = l.Bind("cn=admin,dc=example,dc=com", "password")
	if err != nil {
		log.Fatal(err)
	}

	// 查询证书
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=pkiCertificate)",
		[]string{"cn", "certificateRevocationList;binary", "userCertificate;binary"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Printf("Search result: %d entries\n", len(sr.Entries))
	for _, entry := range sr.Entries {
		fmt.Printf("DN: %s\n", entry.DN)
		for _, cert := range entry.GetAttributeValues("userCertificate;binary") {
			fmt.Printf("Certificate (binary):\n%#v\n", cert)
		}
		for _, crl := range entry.GetAttributeValues("certificateRevocationList;binary") {
			fmt.Printf("CRL (binary):\n%#v\n", crl)
		}
	}
}
*/
