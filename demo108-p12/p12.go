package main

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	gopkcs12 "software.sslmate.com/src/go-pkcs12"
	"time"
)

func certToP12(certBuf, keyBuf []byte) (p12Cert, certPwd string, err error) {

	caBlock, _ := pem.Decode(certBuf)
	crt, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("证书解析异常, Error : %v", err)
		fmt.Printf("%v", err)
		return
	}

	keyBlock, _ := pem.Decode(keyBuf)
	priKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("证书密钥解析key异常, Error : %v", err)
		fmt.Printf("%v", err)
		return
	}

	certPwd = "p12pwdword"

	pfx, err := gopkcs12.Encode(rand.Reader, priKey, crt, nil, certPwd)
	if err != nil {
		err = fmt.Errorf("pem to p12 转换证书异常, Error : %v", err)
		fmt.Printf("%v", err)
		return
	}

	return base64.StdEncoding.EncodeToString(pfx), certPwd, err

}

func getMDMCertCont(cert []byte, mdmPwd string) (certbase64 string, certsn string, after time.Time, err error) {
	// extract key and cert
	_, crt, err := gopkcs12.Decode(cert, mdmPwd)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("%v\n", cert)
	fmt.Printf(" crt.SerialNumber:%v\n", crt.SerialNumber)
	serialHex := fmt.Sprintf("%x", crt.SerialNumber)
	if len(serialHex)%2 == 1 {
		serialHex = fmt.Sprintf("0%s", serialHex)
	}

	return base64.StdEncoding.EncodeToString(cert), serialHex, crt.NotAfter, nil
}

//Error : pkcs12: error reading P12 data: asn1: structure error: tags don't match (16 vs {class:1 tag:13 length:73 isCompound:false}) {optional:false explicit:false application:false private:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} pfxPdu @2err:pkcs12: error reading P12 data: asn1: structure error: tags don't match (16 vs {class:1 tag:13 length:73 isCompound:false}) {optional:false explicit:false application:false private:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} pfxPdu @2
//Process finished with the exit code 0
