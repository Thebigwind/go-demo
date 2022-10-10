package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {

}

func VerifyCertificate(rootCert string, certFile string) (bool, *time.Time, error) {
	if len(rootCert) == 0 || len(certFile) == 0 {
		return false, nil, fmt.Errorf("cert is null")
	}
	verifyCmdStr := fmt.Sprintf("gmssl verify -verbose -no-CAfile -no-CApath -partial_chain -trusted %s %s", rootCert, certFile)
	res, err := Command(verifyCmdStr)
	if err != nil {
		return false, nil, err
	}
	if !strings.Contains(res, "OK") {
		return false, nil, fmt.Errorf("Certificate validation failed, res:%s", res)
	}

	//获取证书有效期
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -dates", certFile)
	res, err = Command(cmdStr)
	if err != nil {
		return false, nil, err
	}
	compileRegex := regexp.MustCompile("notAfter=(.+)") // 提取过期时间
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return false, nil, fmt.Errorf("get certificate expiration time error")
	}
	afterStr := matchArr[len(matchArr)-1]
	afterTime, err := GetGmtTime(afterStr)
	if err != nil {
		return false, nil, err
	}
	if time.Now().Unix() > afterTime.Unix() {
		return false, nil, fmt.Errorf("Certificate expired")
	}

	return true, afterTime, nil
}

// 根据 GMT 字符串获取对应的时间
func GetGmtTime(secStr string) (*time.Time, error) {
	// May 29 08:00:17 2023 GMT
	formatTimeStr := "Jan 2 15:04:05 2006 GMT"
	start, err := time.Parse(formatTimeStr, secStr)
	if err == nil {
		realTime := start.Add(8 * time.Hour)
		return &realTime, nil
	} else {
		return nil, err
	}
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

/**
*  CheckGetCertificateSn 验证节点证书，并从证书中获取签名值
* @param tmpPath 临时证书目录 传入config.DefaultConfig.CertManage.TmpPath
* @param orgCerContent 组织证书内容
* @param cerContent 节点证书内容
* @return string, error  返回值：返回签名值和错误信息
 */
func CheckGetCertificateSn(tmpPath string, orgCerContent string, cerContent string) (string, error) {
	err := CreateDir(tmpPath)
	if err != nil {
		return "", err
	}
	tmpCert := fmt.Sprintf("%s/tmp.crt", tmpPath)
	cmd := fmt.Sprintf(`echo "%s" > %s`, cerContent, tmpCert)
	_, err = Command(cmd)
	if err != nil {
		return "", fmt.Errorf("create tmp cert file err: %v", err.Error())
	}
	orgTmpCert := fmt.Sprintf("%s/org_tmp.crt", tmpPath)
	cmd = fmt.Sprintf(`echo "%s" > %s`, orgCerContent, orgTmpCert)
	_, err = Command(cmd)
	//删除临时文件
	defer func() {
		rmFileCmd := fmt.Sprintf("rm -rf %s %s", orgTmpCert, tmpCert)
		_, err = Command(rmFileCmd)
	}()
	if err != nil {
		return "", fmt.Errorf("create tmp cert file err: %v", err.Error())
	}

	verifyRes, expiredAt, err := VerifyCertificate(orgTmpCert, tmpCert)
	if !verifyRes || err != nil {
		return "", fmt.Errorf("certificate validation err: %v", err)
	}

	if expiredAt.Unix() < time.Now().Unix() {
		return "", fmt.Errorf("Certificate has expired")
	}

	//获取证书里的签名值
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -subject", tmpCert)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	compileRegex := regexp.MustCompile("OU = (.+),")
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate sn error")
	}
	sn := matchArr[len(matchArr)-1]
	return sn, err
}

// GetCertificateIp 从证书中获取ip
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

// GetCertificatePubkey 从证书中获取Pubkey
func GetCertificatePubkey(certFile string) (string, error) {
	if len(certFile) == 0 {
		return "", errors.New("cert is null")
	}
	//获取证书pubkey信息
	cmdStr := fmt.Sprintf("openssl x509 -in %s -noout -pubkey", certFile)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		return "", errors.New("get certificate pubkey error")
	}
	return res, nil
}

/**
*  GetCertificateSn 从证书中获取签名值
* @param tmpPath 临时证书目录 传入config.DefaultConfig.CertManage.TmpPath
* @param cerContent 节点证书内容
* @return string, error  返回值：返回签名值和错误信息
 */
func GetCertificateSn(tmpPath string, cerContent string) (string, error) {
	err := CreateDir(tmpPath)
	if err != nil {
		return "", err
	}
	tmpCert := fmt.Sprintf("%s/tmp.crt", tmpPath)
	cmd := fmt.Sprintf(`echo "%s" > %s`, cerContent, tmpCert)
	_, err = Command(cmd)
	if err != nil {
		return "", fmt.Errorf("create tmp cert file err: %v", err.Error())
	}

	//获取证书里的签名值
	cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -subject", tmpCert)
	res, err := Command(cmdStr)
	if err != nil {
		return "", err
	}
	compileRegex := regexp.MustCompile("OU = (.+),")
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", errors.New("get certificate sn error")
	}
	sn := matchArr[len(matchArr)-1]
	return sn, err
}

func CreateDir(path string) error {
	if IsDirExists(path) {
		return nil
	} else {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = os.Chmod(path, 0777) //通过chmod重新赋权限
		return err
	}
}

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err) //err!=nil,使用os.IsExist()判断为false,说明文件或文件夹不存在
	} else {
		return fi.IsDir() //err==nil,说明文件或文件夹存在
	}
}
