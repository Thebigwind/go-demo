package main

import (
	"fmt"
)

//sm4
func main() {
	//加密
	sm4Key := []byte("12")
	data := []byte("asdfawev34t¥%")
	result, err := Sm4CFB(sm4Key, data, true)
	if err != nil {
		fmt.Printf("err:%v\n", err.Error())
	} else {
		fmt.Printf("result:%v\n", result)
	}

	test1()
}
