package main

import (
	"crypto/rand"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"os"
)

func GenerateSM2Key2() {
	//1.生成sm2密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	//2.通过x509将私钥反序列化并进行pem编码
	privateKeyToPem, err := x509.WritePrivateKeyToPem(privateKey, nil)
	if err != nil {
		panic(err)
	}
	//3.将私钥写入磁盘文件
	file, err := os.Create("sm2Private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(privateKeyToPem)
	if err != nil {
		panic(err)
	}
	//4.进行SM2公钥断言
	publicKey := privateKey.Public().(*sm2.PublicKey)
	//5.将公钥通过x509序列化并进行pem编码
	publicKeyToPem, err := x509.WritePublicKeyToPem(publicKey)
	if err != nil {
		panic(err)
	}
	//6.将公钥写入磁盘文件
	file, err = os.Create("sm2Public.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(publicKeyToPem)
	if err != nil {
		panic(err)
	}
}

//加密
func EncryptSM2(plainText []byte, pubFileName string) []byte {
	//1.打开公钥文件读取公钥
	file, err := os.Open(pubFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}
	//2.将pem格式公钥解码并反序列化
	publicKeyFromPem, err := x509.ReadPublicKeyFromPem(buf)
	if err != nil {
		panic(err)
	}
	//3.加密
	cipherText, err := publicKeyFromPem.EncryptAsn1(plainText, rand.Reader)
	if err != nil {
		panic(err)
	}
	return cipherText
}

//解密
func DecryptSM2(cipherText []byte, priFileName string) []byte {
	//1.打开私钥问价读取私钥
	file, err := os.Open(priFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}
	//2.将pem格式私钥文件解码并反序列话
	privateKeyFromPem, err := x509.ReadPrivateKeyFromPem(buf, nil)
	if err != nil {
		panic(err)
	}
	//3.解密
	planiText, err := privateKeyFromPem.DecryptAsn1(cipherText)
	if err != nil {
		panic(err)
	}
	return planiText
}
