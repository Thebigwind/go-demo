package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os/exec"
)

func BigInt2Hex(data *big.Int) string {
	//bi := big.NewInt(1234)

	// convert bi to a hexadecimal string
	hex := fmt.Sprintf("%x", data)

	fmt.Println(hex) // output: 75bcd15
	return hex
}
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 Hash函数
func MD5(input string) string {
	tokenByte := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", tokenByte)
}

func Command(arg ...string) (string, error) {
	name := "/bin/bash"
	c := "-c"
	args := append([]string{c}, arg...)
	cmd := exec.Command(name, args...)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("Error:can not obtain stdout pipe for command:%s\n", err.Error())
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	//执行命令
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("Error:The command is err:%s, cmd:%+v", err.Error(), arg)
	}

	//读取所有输出
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("ReadAll Stdout:%s", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("wait:%s, cmd:%+v, err:%+v", err.Error(), arg, stderr.String())
	}

	result := string(outBytes)
	return result, nil
}

func RemoveTmpFile(path string) {

	command := "rm -f " + path
	fmt.Printf("command:%s\n", command)
	_, err := Command(command)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
}

func IF(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
