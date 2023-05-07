package main

import (
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
		fmt.Println(v.GetAttributeValues("gidNumber"))
		fmt.Println(v.GetAttributeValues("homeDirectory"))
		fmt.Println(v.GetAttributeValues("uid"))
		fmt.Println(v.GetAttributeValues("uidNumber"))
		fmt.Println(v.GetAttributeValues("cn"))

	}
}

func TestQuery2(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
	}

	baseDN := "ou=certs,dc=zdlz,dc=com"
	filter := "(objectClass=*)"

	err, data := Query(baseDN, filter)
	if err != nil {
		t.Log("query err:", err)
	}
	for _, v := range data {
		v.PrettyPrint(4)
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

func TestPublish(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	if err := publish(); err != nil {
		t.Log("add err:", err)
	}
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
