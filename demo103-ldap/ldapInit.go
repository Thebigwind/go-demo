package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var globalCon *ldap.Conn

func GetLdapConn() *ldap.Conn {
	return globalCon
}

func InitLdap() {

	kingpin.Parse()
	fmt.Printf("%v, %d\n", *ldapaddr, *ldapport)
	//Dail
	con, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", *ldapaddr, *ldapport))
	if err != nil {
		log.Printf("connect err:%v", err)
		return
	}
	//defer con.Close()
	con.Debug.Enable(*debug)
	//认证
	err = con.Bind(*ldapusername, *ldapuserpassword)
	if err != nil {
		log.Printf("bind err:%v", err)
		return
	}
	globalCon = con
	return
}

func TestInit() error {
	addr := "10.10.10.125"
	port := 389
	con, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", addr, port))
	if err != nil {
		log.Printf("connect err:%v", err)
		return err
	}
	//defer con.Close()
	con.Debug.Enable(*debug)
	//认证
	username := "cn=admin,dc=zdlz,dc=com"
	password := "123456"
	err = con.Bind(username, password)
	if err != nil {
		log.Printf("bind err:%v", err)
		return err
	}
	globalCon = con
	return nil
}
