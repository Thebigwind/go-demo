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

	baseDN := "ou=dev,dc=zdlz,dc=com"
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

	DN := "ou=python,ou=dev,dc=zdlz,dc=com"
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

func TestModify(t *testing.T) {
	err := TestInit()
	if err != nil {
		fmt.Printf("init err:%s", err)
		return
	}

	dn := "uid=luxuefeng,ou=golang,ou=dev,dc=zdlz,dc=com"

	replaces := map[string][]string{
		"cn": {"路雪峰"},
	}

	if err = Modify(dn, nil, replaces); err != nil {
		t.Log("modify err:", err)
	}
}
