package main

import (
	"fmt"
	"strings"
)

func main() {
	//root_cert_conent := "xxxxxxxx"
	//root_cert_file := "root-xx.crt"
	//if err := ioutil.WriteFile(root_cert_file, []byte(root_cert_conent), 0777); err != nil {
	//	fmt.Printf("写根证书文件失败：%v\n", err.Error())
	//	return
	//}else{
	//	fmt.Printf("success：%v\n")
	//}
	str := "Serial Number:\n            8d:78:fc:63:6b:80:52:88"

	arr := strings.Split(str, "\n")

	fmt.Printf("arr[1]:%v\n", arr[1])

	fmt.Println(strings.Trim(arr[1], " "))
	infoStr := "MaxDeviceNum::1234"

	aa := strings.TrimPrefix(infoStr, "MaxDeviceNum:")
	fmt.Println(aa)
}
