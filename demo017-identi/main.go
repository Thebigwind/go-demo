package idcard

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 身份证信息
type idInfo struct {
	AddressCode   int
	Abandoned     int
	Address       string
	AddressTree   []string
	Birthday      time.Time
	Constellation string
	ChineseZodiac string
	Sex           int
	Length        int
	CheckBit      string
}

// IsValid 验证身份证号合法性
func IsValid(id string, strict bool) bool {
	code, err := generateCode(id)
	if err != nil {
		return false
	}

	// 检查顺序码、生日码、地址码
	if !checkOrderCode(code["order"]) || !checkBirthdayCode(code["birthdayCode"]) {
		return false
	}

	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 校验码
	return code["checkBit"] == generatorCheckBit(code["body"])
}

// GetInfo 获取身份证信息
func GetInfo(id string, strict bool) (idInfo, error) {
	// 验证有效性
	if !IsValid(id, strict) {
		return idInfo{}, errors.New("Not Valid ID card number.")
	}

	code, _ := generateCode(id)
	addressCode, _ := strconv.Atoi(code["addressCode"])

	// 生日
	birthday, _ := time.Parse("20060102", code["birthdayCode"])

	// 性别
	sex := 1
	sexCode, _ := strconv.Atoi(code["order"])
	if (sexCode % 2) == 0 {
		sex = 0
	}

	// 长度
	length, _ := strconv.Atoi(code["type"])

	return idInfo{
		AddressCode: addressCode,
		Abandoned:   0,
		Birthday:    birthday,
		Sex:         sex,
		Length:      length,
		CheckBit:    code["checkBit"],
	}, nil
}

// UpgradeId 15位升级18位号码
func UpgradeId(id string) (string, error) {
	if !IsValid(id, true) {
		return "", errors.New("Not Valid ID card number.")
	}

	code, _ := generateShortCode(id)

	body := code["addressCode"] + code["birthdayCode"] + code["order"]

	return body + generatorCheckBit(body), nil
}

// 生成Bit码
func generatorCheckBit(body string) string {
	// 位置加权
	var posWeight [19]float64
	for i := 2; i < 19; i++ {
		weight := int(math.Pow(2, float64(i-1))) % 11
		posWeight[i] = float64(weight)
	}

	// 累身份证号body部分与位置加权的积
	var bodySum int
	bodyArray := strings.Split(body, "")
	count := len(bodyArray)
	for i := 0; i < count; i++ {
		bodySub, _ := strconv.Atoi(bodyArray[i])
		bodySum += bodySub * int(posWeight[18-i])
	}

	// 生成校验码
	checkBit := (12 - (bodySum % 11)) % 11
	if checkBit == 10 {
		return "x"
	}
	return strconv.Itoa(checkBit)
}

// 生成数据
func generateCode(id string) (map[string]string, error) {
	length := len(id)
	if length == 15 {
		return generateShortCode(id)
	}

	if length == 18 {
		return generateLongCode(id)
	}

	return map[string]string{}, errors.New("Invalid ID card number length.")
}

// 生成短数据
func generateShortCode(id string) (map[string]string, error) {
	if len(id) != 15 {
		return map[string]string{}, errors.New("Invalid ID card number length.")
	}

	mustCompile := regexp.MustCompile("(.{6})(.{6})(.{3})")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))

	return map[string]string{
		"body":         subMatch[0],
		"addressCode":  subMatch[1],
		"birthdayCode": "19" + subMatch[2],
		"order":        subMatch[3],
		"checkBit":     "",
		"type":         "15",
	}, nil
}

// 生成长数据
func generateLongCode(id string) (map[string]string, error) {
	if len(id) != 18 {
		return map[string]string{}, errors.New("Invalid ID card number length.")
	}
	mustCompile := regexp.MustCompile("((.{6})(.{8})(.{3}))(.)")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))

	return map[string]string{
		"body":         subMatch[1],
		"addressCode":  subMatch[2],
		"birthdayCode": subMatch[3],
		"order":        subMatch[4],
		"checkBit":     subMatch[5],
		"type":         "18",
	}, nil
}

// 检查出生日期码
func checkBirthdayCode(birthdayCode string) bool {
	year, _ := strconv.Atoi(substr(birthdayCode, 0, 4))
	if year < 1800 {
		return false
	}

	nowYear := time.Now().Year()
	if year > nowYear {
		return false
	}

	_, err := time.Parse("20060102", birthdayCode)

	return err == nil
}

// 检查顺序码
func checkOrderCode(orderCode string) bool {
	return len(orderCode) == 3
}

// substr 截取字符串
func substr(source string, start int, end int) string {
	r := []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
