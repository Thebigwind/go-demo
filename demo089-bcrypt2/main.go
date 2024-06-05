package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const SALT = "$2a$04$qDNasmVgJTgoIv8QVt1/LO"

func main() {

	// 方法1：
	// GenerateFromPassword 以给定的代价返回密码的 bcrypt 哈希值。如果给定的成本小于 MinCost
	// 则成本将设置为 DefaultCost。使用此包中定义的 CompareHashAndPassword 将返回的散列密码与其明文版本进行比较。
	pass := "zdlzdev@zt.tech"

	salt := ""
	pwd := []byte(pass + salt)
	//pwd = []byte("123456")
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		//log.Fatalln(err.Error())
	}
	fmt.Println(string(hash))

	dbpas := "$2a$10$4Ux0/4iAeN54BSMOrRfdyeiR4cOT5kPtKVlhNXwamhrb6IOGADt4G"
	// CompareHashAndPassword 将 bcrypt 散列密码与其可能的明文等效密码进行比较。成功时返回 nil，失败时返回错误。
	if err := bcrypt.CompareHashAndPassword([]byte(dbpas), pwd); err != nil {
		fmt.Printf("密码错误：%s", err.Error())
		return
	} else {
		fmt.Println("密码正确")
	}

}

//func hashPwd(pwd string) (string, error) {
//	pwdByte := []byte(pwd)
//	salt64 := []byte(SALT)[7:]
//	passhash, err := bcrypt.Bcrypt(pwdByte, 4, salt64)
//	if err != nil {
//		return "", err
//	}
//	ans := SALT + string(passhash)
//	return ans, nil
//}
