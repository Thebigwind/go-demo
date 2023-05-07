package main

/*
import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func publishCRL() {
	// 加载 CRL 文件
	crlData, err := ioutil.ReadFile("crl.crl")
	if err != nil {
		log.Fatal(err)
	}

	// 解码 CRL
	var crl pkix.CertificateList
	_, err = asn1.Unmarshal(crlData, &crl)
	if err != nil {
		log.Fatal(err)
	}

	// 连接到 LDAP 服务器
	l, err := ldap.Dial("tcp", "ldap.example.com:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// 绑定
	err = l.Bind("cn=admin,dc=example,dc=com", "password")
	if err != nil {
		log.Fatal(err)
	}

	// 准备 CRL 对象
	entry := ldap.NewAddRequest("cn="+crl.TBSCertList.Issuer.String()+",ou=crls,dc=example,dc=com", nil)
	entry.Attribute("objectClass", []string{"top", "cRLDistributionPoint", "authorityRevocationList"})
	entry.Attribute("cn", []string{crl.TBSCertList.Issuer.String()})
	entry.Attribute("certificateRevocationList;binary", [][]byte{crlData})
	entry.Attribute("cRLDistributionPoint;binary", [][]byte{crlData})
	entry.Attribute("authorityRevocationList;binary", [][]byte{crlData})
	entry.Attribute("cACertificate;binary", [][]byte{crl.TBSCertList.IssuerName.FullBytes})

	// 设置 CRL 的有效期
	now := time.Now()
	entry.Attribute("validityPeriodStart", []string{now.Format("20060102150405Z")})
	entry.Attribute("validityPeriodEnd", []string{crl.TBSCertList.NextUpdate.Format("20060102150405Z")})

	// 发布 CRL
	err = l.Add(entry)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CRL added successfully")
}


//这个示例代码查询了 OpenLDAP 中 ou=crls,dc=example,dc=com 下的所有 CRL，并对每个 CRL 进行解析，打印出 CRL 名称、发布时间、下次更新时间和吊销证书数量等信息。
//请注意，这个示例代码只处理单个 CRL，如果您需要批量查询 CRL，需要对代码进行修改。另外，您可能需要根据您的 LDAP 服务器设置调整代码中的服务器地址和绑定参数。

func searchCRL() {

	// 连接到 LDAP 服务器
	l, err := ldap.Dial("tcp", "ldap.example.com:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// 绑定
	err = l.Bind("cn=admin,dc=example,dc=com", "password")
	if err != nil {
		log.Fatal(err)
	}

	// 查询所有 CRL
	searchRequest := ldap.NewSearchRequest(
		"ou=crls,dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=cRLDistributionPoint)",
		[]string{"cn", "certificateRevocationList;binary", "authorityRevocationList;binary", "cRLDistributionPoint;binary"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 解析 CRL
	for _, entry := range sr.Entries {
		cn := entry.GetAttributeValue("cn")
		crlData := entry.GetRawAttributeValue("certificateRevocationList;binary")
		_, crl, err := ParseCRL(crlData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CRL: %s, This update: %s, Next update: %s, Revoked certificates: %d\n", cn, crl.ThisUpdate.Format("2006-01-02 15:04:05"), crl.NextUpdate.Format("2006-01-02 15:04:05"), len(crl.TBSCertList.RevokedCertificates))
	}
}

func ParseCRL(data []byte) (*pkix.CertificateList, *pkix.TBSCertList, error) {
	var crl pkix.CertificateList
	var tbsCertList pkix.TBSCertList
	if _, err := asn1.Unmarshal(data, &crl); err != nil {
		return nil, nil, fmt.Errorf("Failed to parse CRL: %s", err)
	}
	if _, err := asn1.Unmarshal(crl.TBSCertList.Raw, &tbsCertList); err != nil {
		return nil, nil, fmt.Errorf("Failed to parse CRL: %s", err)
	}
	return &crl, &tbsCertList, nil
}
*/
