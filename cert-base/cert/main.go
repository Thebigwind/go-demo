package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

func GetCertificateNodeIp(certFile string) (string, error) {
	if len(certFile) == 0 {
		return "", errors.New("cert is null")
	}
	//获取证书text
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -text", certFile)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	// 获取ip
	compileRegex1 := regexp.MustCompile("IP Address:(.+)")

	//compileRegex := regexp.MustCompile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	matchArr := compileRegex1.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate ip error")
	}
	ipv4 := matchArr[0]
	return ipv4, nil
}

func main() {
	node, err := GetCertificateIp("/home/node.crt")
	fmt.Printf("nodeip1:%s\n", node)
	fmt.Printf("err:%s\n", err)

	node, err = GetNodeCerIp("/home/node.crt")
	fmt.Printf("nodeip2:%s\n", node)
	fmt.Printf("err:%s\n", err)

	//GetCertIp
	//node,_, err = GetCertIp("/home/node.crt")
	//fmt.Printf("nodeip2:%s\n", node)
	//fmt.Printf("err:%s\n", err)

	rep, err := GetNodeCerIp("/home/node-cert-test.crt")
	fmt.Printf("res1:%s\n", rep)
	fmt.Printf("err:%s\n", err)

	res2, err := GetOrgServerAddr("/home/node-cert-test.crt")
	fmt.Printf("res2:%s\n", res2)
	fmt.Printf("err:%s\n", err)
}
func GetCertificateIp(certFile string) (string, error) {
	if len(certFile) == 0 {
		return "", errors.New("cert is null")
	}
	//获取证书subject信息
	// subject= /C=CN/ST=Beijing/L=Beijing/O=zdlz/OU=zdlz/CN=10.10.10.65
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -subject", certFile)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	// 获取ip
	compileRegex := regexp.MustCompile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate ip error")
	}
	ipv4 := matchArr[0]
	return ipv4, nil
}
func GetOrgServerAddr(certFile string) (string, error) {
	//获取证书text
	res, err := GetCertText(certFile)
	if err != nil {
		return "", err
	}
	// 获取ip
	compileRegex1 := regexp.MustCompile("DNS:(.+)")
	matchArr := compileRegex1.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get org server info error")
	}
	data := matchArr[0]

	orgServerAddr := strings.TrimPrefix(data, "DNS:")

	return orgServerAddr, nil
}

func GetCertText(certFile string) (string, error) {
	if len(certFile) == 0 {
		return "", errors.New("cert is null")
	}
	//获取证书text
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -text", certFile)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	return res, nil
}
func GetNodeCerIp(certFile string) (string, error) {

	//获取证书text
	res, err := GetCertText(certFile)
	if err != nil {
		return "", err
	}
	// 获取ip
	compileRegex1 := regexp.MustCompile("IP Address:(.+)")
	matchArr := compileRegex1.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate data error")
	}
	data := matchArr[0]
	fmt.Printf("data:%s", data)
	compileRegex2 := regexp.MustCompile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	matchArr = compileRegex2.FindStringSubmatch(data)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate ip error")
	}
	nodeIp := matchArr[0]

	return nodeIp, nil
}
func GetCertIp(certFile string) (string, string, error) {
	if len(certFile) == 0 {
		return "", "", errors.New("cert is null")
	}
	//获取证书text
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -text", certFile)
	res, err := Command(cmdStr)
	if err != nil {
		return "", "", err
	}
	// 获取ip
	compileRegex1 := regexp.MustCompile("IP Address:(.+)")
	matchArr := compileRegex1.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", "", errors.New("get certificate ip error")
	}
	data := matchArr[0]

	compileRegex2 := regexp.MustCompile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	matchArr = compileRegex2.FindStringSubmatch(data)
	if len(matchArr) == 0 {
		return "", "", errors.New("get certificate ip error")
	}
	nodeIp := matchArr[0]
	orgIp := ""
	if len(matchArr) == 2 {
		orgIp = matchArr[1]
	}
	return nodeIp, orgIp, nil
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
		return "", fmt.Errorf("Error:The command is err:%s, cmd:%+v", err, arg)
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
