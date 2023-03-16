package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	// 方法1：
	// GenerateFromPassword 以给定的代价返回密码的 bcrypt 哈希值。如果给定的成本小于 MinCost
	// 则成本将设置为 DefaultCost。使用此包中定义的 CompareHashAndPassword 将返回的散列密码与其明文版本进行比较。
	pass := "zdlzdev@zdlz.tech1"
	//salt := "$2a$04$qDNasmVgJTgoIv8QVt1/LO"
	salt := ""
	pwd := []byte(pass + salt)
	//pwd = []byte("123456")
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		//log.Fatalln(err.Error())
	}
	fmt.Println(string(hash))
	// $2a$10$2/KZKEaWwxDp.vN2A4wplu3JBZ78Fhd6ECEjapoW.3Jx3Kpw924.C
	//$2a$10$CnHZNeqejS7F/CgeiNkCnOKbjC39Py5A6a1wKurAqTKQ6G4dYD.Xi

	//dbpas := "$2a$10$OVW0MmWrPVIl8f0ACVDQme3fe8EBiMPe1U0cpfdtnDsvBlByz8/yS"
	dbpas := "$2a$04$qDNasmVgJTgoIv8QVt1/LOivCIfs24jpRuSSZJpqG2ih2YAKGEfvK"
	dbpas = "$2a$10$yao3Vw/sNqDkAHrehyGJjOqSUwrxlxx8j7TEtn31QgyRpXKPGR6fC"
	dbpas = "$2a$10$x8z8ioXrdDpvfP32DTMCN.T1O3ImAbFgxYhbm203SAxF49os5HutW"
	// CompareHashAndPassword 将 bcrypt 散列密码与其可能的明文等效密码进行比较。成功时返回 nil，失败时返回错误。
	if err := bcrypt.CompareHashAndPassword([]byte(dbpas), pwd); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("密码正确")
	}

}
