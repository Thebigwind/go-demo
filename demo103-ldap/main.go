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
	ldapusername     = kingpin.Flag("username", "ldap connect usernmae").Default("cn=admin,dc=zdlz,dc=com").String()
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

func QueryCert(baseDN string, filter string) (error, []*ldap.Entry) {
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, //查找范围：整个树
		ldap.NeverDerefAliases, //在搜索中别名(cn, ou)是否废弃
		0,                      // 大小设置,一般设置为0
		0,                      //时间设置,一般设置为0
		false,
		filter, //过滤条件
		[]string{"dn", "cn", "uid", "sn", "userCertificate;binary"}, //返回的属性值
		nil,
	)
	searchResult, err := GetLdapConn().Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
		return err, nil
	}

	if len(searchResult.Entries) == 0 {
		fmt.Println("No certificate found")
		return fmt.Errorf("no certificate found"), nil
	}
	//certBytes := searchResult.Entries[0].GetRawAttributeValue("userCertificate;binary") //binary
	//if len(certBytes) == 0 {
	//	fmt.Println("No certificate found")
	//	return fmt.Errorf("no certificate found"), nil
	//}
	//cert := base64.StdEncoding.EncodeToString(certBytes)
	//fmt.Printf("cert:%v\n", string(cert))
	fmt.Println("-----------")
	//// 遍历搜索结果并打印证书信息
	//for _, entry := range searchResult.Entries {
	//	fmt.Printf("DN: %s\n", entry.DN)
	//
	//	for _, certBytes := range entry.GetRawAttributeValues("userCertificate;binary") {
	//		cert := base64.StdEncoding.EncodeToString(certBytes)
	//		fmt.Printf("Certificate: %s\n", cert)
	//	}
	//}

	//block, _ := pem.Decode([]byte(certBytes))
	//if block == nil {
	//	fmt.Println("Failed to decode certificate")
	//	return fmt.Errorf("Failed to decode certificate"), nil
	//}
	//cert, err := x509.ParseCertificate(block.Bytes)
	//if err != nil {
	//	fmt.Printf("Failed to parse certificate: %v\n", err)
	//	return fmt.Errorf("Failed to parse certificate"), nil
	//}
	//fmt.Printf("cert:%v\n", cert)

	//

	//os.Exit(0)
	return nil, searchResult.Entries
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
		[]string{"dn", "cn", "uid", "sn", "userpassword", "homedirectory", "uidNumber", "gidNumber", "description", "objectClass"}, //返回的属性值
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

func ModifyDN(DN, newDN string, ifDeleteOldDN bool) error {
	//"uid=luxuefeng,ou=golang,ou=dev,dc=xyz,dc=com
	modifyDNReq := ldap.NewModifyDNRequest(DN, newDN, ifDeleteOldDN, "")

	err := GetLdapConn().ModifyDN(modifyDNReq)
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

/*
查询了所有 objectClass 为 pkiCertificate 的对象。它还请求了 cn, userCertificate;binary, certificateRevocationList;binary 属性。
最后，它遍历了所有的结果，输出了每个对象的 DN 和证书/撤销列表。请注意，代码中使用了 InsecureSkipVerify 参数来禁用了 TLS 证书验证，您可能需要在实际应用中使用正确的证书验证方式。
*/
func SearchCert() {
	// 查询证书
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=pkiCertificate)",
		[]string{"cn", "certificateRevocationList;binary", "userCertificate;binary"},
		nil,
	)

	sr, err := GetLdapConn().Search(searchRequest)
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
