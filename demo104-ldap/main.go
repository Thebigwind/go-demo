package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"sync"
)

var (
	ldapaddr         = kingpin.Flag("addr", "ldap addr").Default("10.10.10.125").String()
	ldapport         = kingpin.Flag("port", "ldap connect port").Default("389").Int()
	ldapusername     = kingpin.Flag("username", "ldap connect usernmae").Default("cn=admin,dc=xyz,dc=com").String()
	ldapuserpassword = kingpin.Flag("password", "ldap connect password").Default("123456").String()
	debug            = kingpin.Flag("debug", "run with debug").Default("false").Bool()
)
var once sync.Once

func main() {
	//初始化连接
	once.Do(InitLdap)

	//查询
	err, data := Query("ou=dev,dc=xyz,dc=com", "(objectClass=*)")
	if err != nil {
		return
	}
	for _, item := range data {
		item.Print()
	}

	//增加
}

//查询
func Query(baseDN string, filter string) (error, []*ldap.Entry) {
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, //查找范围：整个树
		ldap.NeverDerefAliases, //在搜索中别名(cn, ou)是否废弃
		0,                      // 大小设置,一般设置为0
		0,                      //时间设置,一般设置为0
		false,
		filter, //过滤条件
		[]string{"dn", "cn", "uidNumber", "gidNumber", "description", "objectClass"}, //返回的属性值
		nil,
	)
	searchResult, err := GetLdapConn().Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
		return err, nil
	}

	return nil, searchResult.Entries
}

//添加
func Add(DN string, adds map[string][]string) error {
	addReq := ldap.NewAddRequest(DN, nil)
	if adds != nil {
		for k, v := range adds {
			addReq.Attribute(k, v)
		}
	}
	err := GetLdapConn().Add(addReq)
	if err != nil {
		log.Printf("add err:%v", err)
		return err
	}
	return nil
}

//修改
func Modify(DN string, adds map[string][]string, replaces map[string][]string) error {
	//"uid=luxuefeng,ou=golang,ou=dev,dc=xyz,dc=com
	modifyReq := ldap.NewModifyRequest(DN, nil)
	//添加
	if adds != nil {
		for k, v := range adds {
			modifyReq.Add(k, v)
		}
	}
	//替换
	if replaces != nil {
		for k, v := range replaces {
			modifyReq.Replace(k, v)
		}
	}

	err := GetLdapConn().Modify(modifyReq)
	if err != nil {
		log.Printf("modify err:%v", err)
		return err
	}
	return nil
}

//删除
func Delete(DN string) error {
	delReq := ldap.NewDelRequest(DN, nil)
	err := GetLdapConn().Del(delReq)
	if err != nil {
		log.Printf("del err:%v", err)
		return err
	}
	return nil
}

func test2() {
	kingpin.Parse()
	fmt.Printf("%v, %d\n", *ldapaddr, *ldapport)
	//Dail
	con, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", *ldapaddr, *ldapport))
	if err != nil {
		log.Fatal("connect err:", err)
		return
	}
	defer con.Close()
	con.Debug.Enable(*debug)
	//认证
	err = con.Bind(*ldapusername, *ldapuserpassword)
	if err != nil {
		log.Fatal("bind err:", err)
		return
	}

	//查询
	searchRequest := ldap.NewSearchRequest(
		"ou=dev,dc=xyz,dc=com", //ou=golang,ou=dev,dc=xyz,dc=com"
		ldap.ScopeWholeSubtree, //查找范围：整个树
		ldap.NeverDerefAliases, //在搜索中别名(cn, ou)是否废弃
		0,                      // 大小设置,一般设置为0
		0,                      //时间设置,一般设置为0
		false,
		"(objectClass=*)", //过滤条件
		[]string{"dn", "cn", "uidNumber", "gidNumber", "description", "objectClass"}, //返回的属性值
		nil,
	)
	searchResult, err := con.Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
		return
	}
	//log.Printf("%d", len(searchResult.Entries))
	for _, item := range searchResult.Entries {
		item.PrettyPrint(4)
	}

	//添加
	sql := ldap.NewAddRequest("uid=xuchengbao,ou=golang,ou=dev,dc=xyz,dc=com", nil)
	sql.Attribute("uidNumber", []string{"1011"})
	sql.Attribute("gidNumber", []string{"1003"})
	sql.Attribute("userPassword", []string{"123456"})
	sql.Attribute("homeDirectory", []string{"/home/xuchengbao"})
	sql.Attribute("cn", []string{"许成宝"})
	sql.Attribute("uid", []string{"xuchengbao"})
	sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})

	err = con.Add(sql)
	if err != nil {
		log.Printf("add err:%v", err)
		//return
	}

	//修改
	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest("uid=luxuefeng,ou=golang,ou=dev,dc=xyz,dc=com", nil)
	//modify.Add("description", []string{"An example user luxuefeng"})
	modify.Replace("description", []string{"luxf@zdlz.cc"})

	err = con.Modify(modify)
	if err != nil {
		log.Printf("modify err:%v", err)
		return
	}

	//删除
	del := ldap.NewDelRequest("ou=shell,ou=dev,dc=xyz,dc=com", nil)

	err = con.Del(del)
	if err != nil {
		log.Printf("del err:%v", err)
		return
	}

}

/*
10.10.10.125, 389
    DN: dc=xyz,dc=com
      objectClass: [top dcObject organization]
    DN: cn=admin,dc=xyz,dc=com
      objectClass: [simpleSecurityObject organizationalRole]
      cn: [admin]
    DN: ou=People,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=golang,ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=clang,ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=java,ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=python,ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
    DN: ou=shell,ou=dev,dc=xyz,dc=com
      objectClass: [organizationalUnit]
*/

func test() {
	// The username and password we want to check
	username := "admin"
	password := "123456"

	bindusername := "cn=admin,dc=xyz,dc=com"
	// bindusername := "admin"
	bindpassword := "123456"

	// l, err := DialURL("ldap://ldap.example.com:389")
	l, err := ldap.DialURL("ldap://10.10.10.125:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("0000")
	// Reconnect with TLS
	// err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("1111")
	// First bind with a read only user
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2222")
	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		"dc=xyz,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		log.Fatal(err)
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
}
