package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	test2()
}

func test2() {
	command := "cat /etc/system-release | awk '{print $1}'"
	sysType, err := Command(command)
	if err != nil {
		fmt.Errorf("查询系统类型err:%s", err.Error())
	}

	fmt.Printf("sysType:%s", sysType)
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
