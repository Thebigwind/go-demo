package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
)

//提取证书信息
func ParseCert(cert string) (CertInfo, error) {

	certBytes := []byte(cert)

	block, _ := pem.Decode(certBytes)
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Printf("Failed to parse certificate: %s", err)
		return CertInfo{}, err
	}
	//解析subject
	subject := c.Subject.String()

	certInfo := CertInfo{
		NotAfter:      c.NotAfter,
		NotBefore:     c.NotBefore,
		Sn:            Bigint2Hex(c.SerialNumber),
		Subject:       subject,
		OrgName:       c.Subject.Organization[0],
		Name:          GetCertName(subject), //Str2Md5(cert),
		Cert:          cert,
		IssuedBy:      c.Issuer.CommonName,
		KeyAlgorithm:  c.PublicKeyAlgorithm.String(),
		SignAlgorithm: c.SignatureAlgorithm.String(),
		//KeyLen:        c.PublicKey.(*rsa.PublicKey).Size() * 8,
	}
	//fmt.Printf("subject:%+v", c.Subject.String())
	return certInfo, nil
}

func GetCertName(subject string) string {
	arr := strings.Split(subject, ",")
	data := ""
	//提取CN=name
	if len(arr) > 1 {
		data = arr[0]
	}
	target := strings.Split(data, "=")
	name := ""
	//解析CN
	if len(target) >= 2 {
		//提取name
		name = target[1]
	}
	return name
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
	KeyLen        int
}
