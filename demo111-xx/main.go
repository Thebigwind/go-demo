package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	//gmx509 "github.com/tjfoc/gmsm/x509"
)

func main() {
	fmt.Println(len("中国"))
	str := "中国的城市"
	for i, v := range str {
		fmt.Printf("i:%d,v:%v\n", i, v)
	}
}

// 判断字符串是否为字母数字
func IsAlphaNumeric3(str string) bool {
	// 使用正则表达式匹配字母数字
	reg := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return reg.MatchString(str)
}

// 判断字符串是否为字母数字
func IsAlphaNumeric2(str string) bool {
	// 遍历字符串，判断每个字符是否为字母数字
	for _, ch := range str {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", ch) {
			return false
		}
	}
	return true
}

// 判断字符串是否为字母数字
func IsAlphaNumeric1(str string) bool {
	// 遍历字符串，判断每个字符是否为字母数字
	for _, ch := range str {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
