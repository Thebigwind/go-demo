package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
	}

	//baseDN := "ou=dev,dc=zdlz,dc=com"
	baseDN := "ou=People,dc=zdlz,dc=com"
	filter := "(objectClass=*)"

	err, data := Query(baseDN, filter)
	if err != nil {
		t.Log("query err:", err)
	}
	for _, v := range data {
		v.PrettyPrint(4)
	}
}

func TestQuery1(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
	}

	baseDN := "uid=luffy,ou=golang,ou=dev,dc=zdlz,dc=com"
	filter := "(objectClass=*)"

	err, data := Query(baseDN, filter)
	if err != nil {
		t.Log("query err:", err)
	}
	for _, v := range data {
		//v.PrettyPrint(4)
		fmt.Println("------")
		fmt.Println(v.GetAttributeValues("homeDirectory"))
		fmt.Println(v.GetAttributeValues("gidNumber"))
		fmt.Println(v.GetAttributeValues("userPassword"))
		fmt.Println(v.GetAttributeValues("uid"))
		fmt.Println(v.GetAttributeValues("uidNumber"))
		fmt.Println(v.GetAttributeValues("cn"))
	}
}

func TestQueryCert(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
	}

	baseDN := "uid=john,ou=certs,dc=zdlz,dc=com"
	filter := "(objectClass=*)"
	//filter = "(userCertificateSerialNumber=123456789)"

	err, data := QueryCert(baseDN, filter)
	if err != nil {
		t.Log("query err:", err)
	}

	for _, entry := range data {
		fmt.Printf("DN: %s\n", entry.DN)
		//for _, certBytes := range entry.GetRawAttributeValues("userCertificate;binary") {
		//	cert := base64.StdEncoding.EncodeToString(certBytes)
		//	//如果需要使用证书数据，可以使用 base64.StdEncoding.DecodeString 方法将其转换回二进制数据
		//	fmt.Printf("Certificate: %s\n", cert)
		//}
		certBytes := entry.GetRawAttributeValue("userCertificate;binary")
		cert := base64.StdEncoding.EncodeToString(certBytes)
		sn := entry.GetAttributeValue("sn")
		cn := entry.GetAttributeValue("cn")
		uid := entry.GetAttributeValue("uid")
		fmt.Printf("Certificate: %s\n", cert)
		fmt.Printf("sn: %s\n", sn)
		fmt.Printf("cn: %s\n", cn)
		fmt.Printf("uid: %s\n", uid)

	}
}

func TestDelete(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	DN := "ou=People,dc=zdlz,dc=com"
	if err = Delete(DN); err != nil {
		t.Log("delete err:", err.Error())
	}
}

func TestDelete2(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	DN := "uid=luffy,ou=golang,ou=dev,dc=zdlz,dc=com"
	if err = Delete(DN); err != nil {
		t.Log("delete err:", err.Error())
	}
}

func TestAdd(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "uid=luffy,ou=golang,ou=dev,dc=zdlz,dc=com"
	attrs := map[string][]string{
		"uidNumber":     {"1012"},
		"gidNumber":     {"1013"},
		"userPassword":  {"123456"},
		"homeDirectory": {"/home/luffy"},
		"cn":            {"路飞"},
		"uid":           {"luffy"},
		"objectClass":   {"shadowAccount", "posixAccount", "account"},
	}
	if err := Add(dn, attrs); err != nil {
		t.Log("add err:", err)
	}
}

func TestAdd2(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "uid=lxf,ou=golang,ou=dev,dc=zdlz,dc=com"
	attrs := map[string][]string{
		"uidNumber":     {"1013"},
		"gidNumber":     {"1013"},
		"userPassword":  {"123456"},
		"homeDirectory": {"/home/lxf"},
		"cn":            {"路雪峰"},
		"uid":           {"lxf"},
		"objectClass":   {"shadowAccount", "posixAccount", "account"},
	}
	if err := Add(dn, attrs); err != nil {
		t.Log("add err:", err)
	}
}

func TestAdd3(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	// 准备要发布的用户信息
	dn := "uid=johndoe1,ou=cert,dc=zdlz,dc=com"
	attrs := map[string][]string{
		"givenName":    {"John1"},
		"userPassword": {"{SSHA}i9y43hf8ygrf49y2h38fh298g"},
		"cn":           {"John Doe1"},
		"sn":           {"Doe1"},
		"uid":          {"lxf"},
		"objectClass":  {"top", "person", "organizationalPerson", "inetOrgPerson"},
		"mail":         {"johndoe1@example.com"},
	}

	if err := Add(dn, attrs); err != nil {
		t.Log("add err:", err)
	}
}

func TestAddDN(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "ou=People,dc=zdlz,dc=com"
	attrs := map[string][]string{
		"objectClass": {"organizationalUnit"},
		"ou":          {"People"},
	}
	if err := Add(dn, attrs); err != nil {
		t.Log("add err:", err)
	}
}

func TestAddDN2(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "ou=users,dc=zdlz,dc=com"
	attrs := map[string][]string{
		"objectClass": {"organizationalUnit"},
	}
	if err := Add(dn, attrs); err != nil {
		t.Log("add err:", err)
	}
}

func TestPubCert(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}
	certPath := "root.crt"
	if err := pubCert(certPath); err != nil {
		t.Log("add err:", err)
	}
	t.Log("test success")
}
func TestPublish2(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	if err := publish2(); err != nil {
		t.Log("add err:", err)
	}
}
func TestPublishmodify(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	if err := publishmodify(); err != nil {
		t.Log("add err:", err)
	}
}

func TestPubCRL(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	pubCRL()
}
func TestModify(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "uid=luffy,ou=golang,ou=dev,dc=zdlz,dc=com"

	replaces := map[string][]string{
		"cn": {"路雪峰"},
	}

	adds := map[string][]string{
		"userid": {"luffy"},
	}

	if err = Modify(dn, adds, replaces); err != nil {
		t.Log("modify err:", err)
	}
}

func TestModifyDN(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "uid=lxf,ou=golang,ou=dev,dc=zdlz,dc=com"
	newDn := "uid=luxuefeng"

	if err = ModifyDN(dn, newDn, false); err != nil {
		t.Log("modify err:", err)
	}
}
