package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func test() {
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
