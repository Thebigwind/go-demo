package main

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"os"
	"strings"
)

var key = []byte("Y!1EoEF$en@1pq&^")

func main() {
	cipher := kingpin.Command("cipher", "cipher tools")
	aes_tool := cipher.Command("aes", "aes tool")

	aes_encoding := aes_tool.Flag("encode", "要加密的原始数据：").String()
	aes_decoding := aes_tool.Flag("decode", "要解密的原始数据：").String()

	kingpin.CommandLine.HelpFlag.Short('h')
	osParse := kingpin.Parse()
	cmds := strings.Split(osParse, " ")

	switch cmds[0] {
	case "cipher":

		switch cmds[1] {

		case "aes":
			if *aes_encoding != "" {
				out, err := AesEncoding(*aes_encoding)
				if err != nil {
					fmt.Printf("aesEncoding err:%s\n", err.Error())
					os.Exit(1)
				} else {
					//fmt.Printf("aesEncoding:%s\n", out)
					fmt.Printf("%s", out)
					os.Exit(0)
				}
			}

			if *aes_decoding != "" {
				out, err := AesDecoding(*aes_decoding)
				if err != nil {
					fmt.Printf("aesDecoding err:%s\n", err.Error())
					os.Exit(1)
				} else {
					//fmt.Printf("aesDecoding out:%s\n", out)
					fmt.Printf("%s", out)
					os.Exit(0)
				}
			}

		}
	}

}

func PadPwd(srcByte []byte, blockSize int) []byte {
	// 16 13       13-3 = 10
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

func AesEncoding(src string) (string, error) {
	srcByte := []byte(src)
	// safer
	block, err := aes.NewCipher(key)
	if err != nil {
		return src, err
	}
	// 密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) //由于字节长度不够，所以要进行字节的填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd, nil
}

// 去掉填充的部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	// 去掉的长度
	unpadNum := int(dst[len(dst)-1])
	if unpadNum == 0 {
		for i := len(dst) - 1; i >= 0; i-- {
			if dst[i] == 0 {
				unpadNum++
			} else {
				break
			}
		}

	}
	return dst[:(len(dst) - unpadNum)], nil
}

// 解密
func AesDecoding(pwd string) (string, error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)

	if err != nil {
		return pwd, err
	}
	block, errBlock := aes.NewCipher(key)
	if errBlock != nil {
		return pwd, errBlock
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, _ = UnPadPwd(dst) // 填充的要去掉
	return string(dst), nil
}
