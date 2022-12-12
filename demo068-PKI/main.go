package main

import (
	"fmt"
	mrand "math/rand"
	"os/exec"
	"path"
	"time"
)

const (
	Private = "private"
	Conf    = "conf"
	Csr     = "csr"
	Cert    = "cert"
)

var CreateTmpMap = map[string]string{
	Private: "openssl genrsa -out %s.key 2048",
	Conf: `cat > %s.conf <<EOF
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
distinguished_name = dn
req_extensions = v3_req

[ dn ]
C = CN
ST = Beijing
L = Beijing
O = zdlz
OU = zdlz
CN = device

[v3_req]
basicConstraints=critical,CA:TRUE

EOF`, //1个%s
	Csr:  "openssl req -new -key %s.key -out %s.csr -config %s.conf",            // 3 个 %s
	Cert: "openssl x509 -req -in %s.csr -out %s.crt -signkey %s.key -days 3650", // 3个 %s
}

func ExecCommand(commond string) error {

	cmd := exec.Command("/bin/bash", "-c", commond)
	bytes, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

func CreatePrivateKey(path string) error {
	commond := fmt.Sprintf(CreateTmpMap[Private], path)
	if err := ExecCommand(commond); err != nil {
		return err
	}
	return nil
}

func CreateConf(path string) error {
	commond := fmt.Sprintf(CreateTmpMap[Conf], path)
	if err := ExecCommand(commond); err != nil {
		return err
	}
	return nil
}

func CreateCsr(path string) error {
	commond := fmt.Sprintf(CreateTmpMap[Csr], path, path, path)
	if err := ExecCommand(commond); err != nil {
		return err
	}
	return nil
}

func CreateCert(path string) error {
	commond := fmt.Sprintf(CreateTmpMap[Cert], path, path, path)
	if err := ExecCommand(commond); err != nil {
		return err
	}
	return nil
}

func RemoveTmpCertFile(path string) error {
	// .key  .conf .csr .crt
	commond := fmt.Sprintf("rm %s.key %s.conf %s.csr %s.crt", path, path, path, path)
	if err := ExecCommand(commond); err != nil {
		return err
	}
	return nil
}

var RandBytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func RandStr(n int) string {
	mrand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = RandBytes[mrand.Int31()%62]
	}
	return string(result)
}

func GetPath() string {
	suffix := RandStr(10)
	return path.Join("./", "device-"+suffix)
}

func main() {

	//path := GetPath()
	//fmt.Printf("path:%s\n", path)
	//
	//if err := CreatePrivateKey(path);err != nil{
	//	fmt.Printf("err :%s\n",err.Error())
	//}
	//if err := CreateConf(path);err != nil{
	//	fmt.Printf("err :%s\n",err.Error())
	//}
	//
	//if err := CreateCsr(path);err != nil{
	//	fmt.Printf("err :%s\n",err.Error())
	//}
	//
	//
	//if err := CreateCert(path);err != nil{
	//	fmt.Printf("err :%s\n",err.Error())
	//}

	path := "device-UlE3YZocbL"
	//time.Sleep(time.Second *5 )
	if err := RemoveTmpCertFile(path); err != nil {
		fmt.Printf("err :%s\n", err.Error())
	}
}
