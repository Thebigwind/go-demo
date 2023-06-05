package common

import (
	"github.com/dlclark/regexp2"
	"regexp"
)

//CheckMailFormat 验证邮箱是合法
func CheckMailFormat(email string) bool {
	r, err := regexp2.Compile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$`, regexp2.RegexOptions(regexp2.IgnoreCase))
	if err != nil {
		return false
	}
	ok, err := r.MatchString(email)
	if !ok || err != nil {
		return false
	}
	return true
}

//CheckMobileFormat 检验手机号的合法性
func CheckMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
