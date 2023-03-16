package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := []byte("http://golang.org/pkg/encoding/base64/#variables")

	base64Str := base64.StdEncoding.EncodeToString(s)
	fmt.Printf("%s\n", base64Str)

	str, _ := base64.StdEncoding.DecodeString(base64Str)
	fmt.Printf("%s\n", str)

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(uDec))
}
