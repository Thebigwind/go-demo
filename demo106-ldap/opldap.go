package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"io/ioutil"
	"log"
	"math/big"
	"os/exec"
	"time"
)

/*
   Dn = "uid=%s,ou=certs,dc=zdlz,dc=com"
   filter = "(sn=abc)"

   crlDn = "uid=%s,ou=crlcerts,dc=zdlz,dc=com"
   filter := "(objectClass=*)"


*/
func main() {
	//testpub()
	//ShowCert()
	ParseCert2("")
}

func ParseCert2(cert string) (CertInfo, error) {

	certBytes, err := ioutil.ReadFile("./demo106-ldap/cfg.crt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return CertInfo{}, err
	}

	//certBytes := []byte(cert)
	fmt.Printf("xx:%v\n", string(certBytes))
	//os.Exit(0)
	block, _ := pem.Decode(certBytes)
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Printf("Failed to parse certificate: %s", err)
		return CertInfo{}, err
	}
	certInfo := CertInfo{
		NotAfter:      c.NotAfter,
		NotBefore:     c.NotBefore,
		Sn:            Bigint2Hex(c.SerialNumber),
		ParentSn:      c.Issuer.String(),  //c.Issuer.SerialNumber,
		Subject:       c.Subject.String(), //c.Subject.CommonName,
		OrgName:       c.Subject.Organization[0],
		Name:          "",
		Cert:          cert,
		IssuedBy:      c.Issuer.CommonName,
		KeyAlgorithm:  c.PublicKeyAlgorithm.String(),
		SignAlgorithm: c.SignatureAlgorithm.String(),
		keyLen:        c.PublicKey.(*rsa.PublicKey).Size() * 8,
	}
	fmt.Printf("certInfo:%+v\n", certInfo)
	return certInfo, nil
}

func testpub() {
	uid := "namei"
	sn := "xx1xx2xx3"
	cn := "subject=aa"
	base64CertContent := "MIIDODCCAiACCQCH3OLpJ4IsozANBgkqhkiG9w0BAQUFADBeMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsGA1UECwwEemRsejENMAsGA1UEAwwEcm9vdDAeFw0yMzA0MjUwMzM0MTdaFw0zMzA0MjIwMzM0MTdaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYDVQQKDAR6ZGx6MQ0wCwYDVQQLDAR6ZGx6MQ0wCwYDVQQDDARyb290MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+xMgtkVyjBwUHcSqh4NIQ+GskSCVdYtztQD6coic8g9pH/WNWYf0kQmVLl8XySB50wBwUpXGKYCWhZcP0/bMULEzjfbZr2KKjKSd4Xt0m/5gLX2bhCbIQpPyQisnWS46i+yt2LgzLsk9zJyghQgdNJVj054mIKCLHjmG4/9YoiLfnbNiwMiMUMdu7X5Grp0CJe9sB1cr6Nnby234Kopvd0snWfNVepyO7ETjMpqGUqsaWVU/X2EWGCTaYtIN4dQYnKOBIKn4HrGSkHh5cHtYvoi29BDq7x11BN9Wuw+ls1Bp8MsYf4JmRj/YhBPsOnGLEsoiCyLPTEkS+SBMViuehwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCcxvMe6J9iGFztb7ODps4IMMQ0UvS4PclV9zMvQE3VSsu6U898j7HIqmIVDaf3e1/uP7JVtEGPE4C9n5KXV1Q2j2kygmZXewBlWVMknFYzK/MXFxHhvs3JfJrJDVWBQb5iFu5KgHI6dmFFq91OWdjvECBnomrRaTGOlIBsaeFSEekC5yFRwzpHdhiPieMen4D8f2nl0FlAu5jGMwcH5Ae4TSSIGnKzX7toCSkHD99azVTZ4/CpBW57cXPLckZOg8s2vuXluElSlS+6DLgUWVSwCQvzYbLWs+T+cgs2VA/XN2m+FIiHhwr8N303vL5SjUn9DLSDgBiGX4z1ifJIsKCu"

	dn := fmt.Sprintf(Dn, uid)
	ldifPath := GetMd5String(dn + sn + cn + base64CertContent)
	defer RemoveTmpFile(ldifPath)
	//构建发布结构
	if err := CreateLdif(ldifPath, dn, sn, cn, base64CertContent); err != nil {
		return
	}

	//
	if err := publish(ldifPath); err != nil {
		return
	}
}

const Dn = "uid=%s,ou=certs,dc=zdlz,dc=com"

func CreateLdif(ldifPath, dn, sn, cn, base64CertContent string) error {
	certStr := `cat > %s <<EOF
dn: %s
objectClass: top
objectClass: person
objectClass: inetOrgPerson
objectClass: pkiUser
sn: %s
cn: %s
userCertificate;binary:: %s
EOF
`

	//certPath := GetMd5String(dn + sn + cn + base64CertContent)
	certCmd := fmt.Sprintf(certStr, ldifPath, dn, sn, cn, base64CertContent)

	_, err := Command(certCmd)
	if err != nil {
		return err
	}
	return nil
}

const (
	User = "cn=admin,dc=zdlz,dc=com"
	Pass = "123456"
	Host = "ldap://10.10.10.125:389"
)

func publish(ldifPath string) error {
	ldapaddCmd := `ldapadd -c -x -D %s -w %s  -H %s -f %s`
	ldapaddCmd = fmt.Sprintf(ldapaddCmd, User, Pass, Host, ldifPath)
	_, err := Command(ldapaddCmd)
	if err != nil {
		return err
	}
	return nil
}

type Cert struct {
	cert string
	sn   string
	cn   string
	uid  string
}

func ShowCert() (err error, certs []Cert) {
	baseDn := "ou=certs,dc=zdlz,dc=com" //baseDn := "uid=lxf,ou=certs,dc=zdlz,dc=com"
	filter := "(sn=abc)"                //filter := "(objectClass=*)" filter := "(uid=lxf)"
	err, data := QueryCert(baseDn, filter)
	if err != nil {
		return
	}

	for i, entry := range data {

		fmt.Printf("DN: %s\n", entry.DN)
		certBytes := entry.GetRawAttributeValue("userCertificate;binary")
		cert := base64.StdEncoding.EncodeToString(certBytes)
		sn := entry.GetAttributeValue("sn")
		cn := entry.GetAttributeValue("cn")
		uid := entry.GetAttributeValue("uid")
		fmt.Printf("Certificate: %s\n", cert)
		fmt.Printf("sn: %s\n", sn)
		fmt.Printf("cn: %s\n", cn)
		fmt.Printf("uid: %s\n", uid)
		//test
		if i == 0 {
			ParseCert(string(certBytes))
		}

		cer := Cert{
			cert: cert,
			sn:   sn,
			cn:   cn,
			uid:  uid,
		}
		certs = append(certs, cer)
	}

	return
}

func Bigint2Hex(data *big.Int) string {
	//bi := big.NewInt(1234)

	// convert bi to a hexadecimal string
	hex := fmt.Sprintf("%x", data)

	fmt.Println(hex) // output: 75bcd15
	return hex
}

//证书信息
type CertInfo struct {
	NotAfter      time.Time
	NotBefore     time.Time
	Sn            string
	ParentSn      string
	Signature     string
	Subject       string
	OrgName       string
	Name          string
	Cert          string
	IssuedBy      string
	SignAlgorithm string
	KeyAlgorithm  string
	keyLen        int
}

func ParseCert(cert string) {
	certBytes, err := ioutil.ReadFile("./demo105-ldap/root.crt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("cert:::%v\n", cert)
	fmt.Printf("certBytes:::%v\n", string(certBytes))
	block, _ := pem.Decode(certBytes)
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Printf("Failed to parse certificate: %s", err)
	}

	notAfter := c.NotAfter
	notBefore := c.NotBefore
	rawIssuer := c.RawIssuer
	serialNumber := c.SerialNumber
	signature := c.Signature
	subject := c.Subject
	pubKey := c.PublicKey

	fmt.Printf("xx:%+v\n", c.Subject.SerialNumber)
	fmt.Printf("notAfter:%+v,notBefore:%+v,rawIssuer:%+v,serialNumber:%+v,signature:%+v,subject:%+v,pubKey:%v",
		notAfter, notBefore, string(rawIssuer), serialNumber, string(signature), subject, pubKey)
}

func QueryCert(baseDN string, filter string) (error, []*ldap.Entry) {
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
	searchResult, err := conn.Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
		return err, nil
	}

	if len(searchResult.Entries) == 0 {
		fmt.Println("No certificate found")
		return fmt.Errorf("no certificate found"), nil
	}

	return nil, searchResult.Entries
}

//查询
func Query(baseDN string, filter string) (error, []*ldap.Entry) {
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
	searchResult, err := conn.Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
		return err, nil
	}

	return nil, searchResult.Entries
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Command(arg ...string) (string, error) {
	name := "/bin/bash"
	c := "-c"
	args := append([]string{c}, arg...)
	cmd := exec.Command(name, args...)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("Error:can not obtain stdout pipe for command:%s\n", err.Error())
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	//执行命令
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("Error:The command is err:%s, cmd:%+v", err.Error(), arg)
	}

	//读取所有输出
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("ReadAll Stdout:%s", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("wait:%s, cmd:%+v, err:%+v", err.Error(), arg, stderr.String())
	}

	result := string(outBytes)
	return result, nil
}

func RemoveTmpFile(path string) {

	command := "rm -f " + path
	fmt.Printf("command:%s\n", command)
	_, err := Command(command)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
}
