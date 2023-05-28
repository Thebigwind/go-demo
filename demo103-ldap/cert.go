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

func parseCert(certPath string) (*x509.Certificate, error) {
	// 加载证书
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Printf("err:%s", err)
		return nil, err
	}

	data := string(certData)[27:1149]
	fmt.Println(data)
	// 解码证书
	block, _ := pem.Decode(certData)
	if block == nil {
		log.Printf("Failed to decode certificate")
		return nil, err
	}

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		log.Printf("err:%s", err)
		return nil, err
	}
	return cert, nil
}

func pubCert(certPath string) error {
	cert, err := parseCert(certPath)
	if err != nil {
		return err
	}
	if err = publish(cert); err != nil {
		return err
	}
	return nil
}
func publish(cert *x509.Certificate) error {

	fmt.Println("------------")
	//fmt.Printf("subject:%+v\n", cert.Subject)
	//fmt.Printf("Issuer:%+v\n", cert.Issuer)
	////fmt.Printf("Issuer raw:%+v\n", string(cert.RawIssuer))
	SerialNumberStr := cert.SerialNumber.Text(16)
	fmt.Printf("SerialNumber:%+v,SerialStr:%s\n", cert.SerialNumber, SerialNumberStr)
	//fmt.Printf("Raw:%+v\n", cert.Raw)
	//fmt.Printf("NotBefore:%+v\n", cert.NotBefore)
	//fmt.Printf("NotAfter:%+v\n", cert.NotAfter)
	//fmt.Printf("Extensions:%+v\n", cert.Extensions)
	fmt.Printf("xxx:%v\n", base64.StdEncoding.EncodeToString(cert.Raw))
	fmt.Printf("+cert.Subject.CommonName:%v\n", cert.Subject.CommonName)
	fmt.Println("------------")
	// 准备证书对象
	entry := ldap.NewAddRequest("cn="+cert.Subject.CommonName+",ou=certs,dc=zdlz,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "person", "inetOrgPerson", "pkiUser"}) // "inetOrgPerson",
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Attribute("sn", []string{SerialNumberStr})
	aa := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"
	//aa = []byte("certificate data")
	entry.Attribute("userCertificate;binary", []string{aa})
	//entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})

	//entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//entry.Attribute("userCertificate", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//entry.Attribute("data", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	l := GetLdapConn()
	// 发布证书
	err := l.Add(entry)
	if err != nil {
		log.Printf("err:%s", err)
		return err
	}

	fmt.Println("Certificate added successfully")
	return nil
}

func publish1() error {
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
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson", "organizationalPerson", "pkiUser"}) // "inetOrgPerson",
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("cn", []string{cert.Subject.CommonName})
	//entry.Attribute("sn", []string{"01"}) //cert.Subject.SerialNumber
	entry.Attribute("sn", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//entry.Attribute("userCertificate", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//entry.Attribute("data", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
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
	certData, err := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo103-ldap/root.crt")
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
	entry.Attribute("objectClass", []string{"top", "inetOrgPerson", "person", "pkiUser"}) //ipsecCert
	entry.Attribute("cn", []string{cert.Subject.CommonName})
	entry.Attribute("sn", []string{"01"}) //cert.Subject.SerialNumber
	////cert.Raw = []byte("MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu")
	cert.Raw = []byte("test")
	entry.Attribute("userCertificate;binary", []string{base64.StdEncoding.EncodeToString(cert.Raw)})
	//data := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"
	//entry.Attribute("userCertificate;binary", []string{data})

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
