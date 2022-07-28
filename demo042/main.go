package main

import (
	"bytes"
	"fmt"
)

func main() {
	GerenateSM2Key()
	src := []byte("这是使用SM2椭圆曲线算法进行数据加解密测试")
	cipherText := EncryptSM2(src, "sm2Public.pem")
	plainText := DecryptSM2(cipherText, "sm2Private.pem")
	flag := bytes.Equal(plainText, src)
	fmt.Println("解密是否成功：", flag)

	srcC := []byte("这是使用SM2椭圆曲线算法进行的签名验签测试")
	signSM2 := SignSM2(srcC, "sm2Private.pem")
	flagC := VerifySM2(srcC, signSM2, "sm2Public.pem")
	fmt.Println("验签结果：", flagC)

}
