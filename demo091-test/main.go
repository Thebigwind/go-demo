package main

import (
	"fmt"
)

//
func main() {

	//panic("aaa")
	//aa := "ChF6ZGx6ZGV2QHpkbHoudGVjaBo8JDJhJDA0JHFETmFzbVZnSlRnb0l2OFFWdDEvTE9pdkNJZnMyNGpwUnVTU1pKcHFHMmloMllBS0dFZnZLIjwkMmEkMDQkcUROYXNtVmdKVGdvSXY4UVZ0MS9MTzJxY2ZhNnJjbGtyekN3VUY5MG0vRnFINjV6YmlMVzIqADIA"
	////decodeString, err := base64.StdEncoding.DecodeString(req.GetReset_())
	//decodeString, err := base64.StdEncoding.DecodeString(aa)
	//if err != nil {
	//	fmt.Errorf("err:",err.Error())
	//	return
	//}
	//fmt.Printf("PasswordReset:%s\n",string(decodeString))

	salt := "abc"
	pass := salt + "12"
	fmt.Println(salt[3:])
	fmt.Println(salt[len(salt):])
	fmt.Printf("%s", pass[len(salt):])
}

//func testDonghui() {
//	key := "zdlzdev@zdlz.tech1"
//	str := "w1Y05kIo7aL3p6p+59caew=="
//	out, err := base64.StdEncoding.DecodeString(str)
//	if err != nil {
//		fmt.Printf("err:%v", err)
//	} else {
//		decData, err := cgo.Sm4EncryptDecrypt([]byte(key), out, false)
//		if err != nil {
//			fmt.Printf("err:%v", err)
//		} else {
//			fmt.Printf("encData:%v", decData)
//			if decData == "password!" {
//				fmt.Println("验证sm4加解密成功")
//			}
//		}
//	}
//}

//func testSecret() {
//	mdata := map[string]string{
//		"cur":  "5KN+R6yU4sKEwBVh8+/Rl0eOM36ZEKFohqZoSlaY2WY=",
//		"next": "hBoqoH77t/4cgA9wufbv+EeOM36ZEKFohqZoSlaY2WY=",
//		"fixd": "PLP9jed8x/a9aYG3BcLYbUeOM36ZEKFohqZoSlaY2WY=",
//	}
//
//	key := "zdlzdev@zdlz.tech1"
//	newkey := "zdlzdev@zdlz.tech"
//	for k, v := range mdata {
//		out, err := base64.StdEncoding.DecodeString(v)
//		if err != nil {
//			fmt.Printf("err:%v", err)
//		} else {
//			decData, err := cgo.Sm4EncryptDecrypt([]byte(key), out, false)
//			if err != nil {
//				fmt.Printf("err k:%s,:%v", k, err)
//			}
//			fmt.Printf("k:%s,v:%s\n", k, decData)
//
//			encData, err := cgo.Sm4EncryptDecrypt([]byte(newkey), []byte(decData), true)
//			if err != nil {
//				fmt.Printf("err:%v", err)
//			}
//			newValue := base64.StdEncoding.EncodeToString([]byte(encData))
//			fmt.Printf("k:%s,newV:%s\n", k, newValue)
//		}
//
//		//
//	}
//
//}
