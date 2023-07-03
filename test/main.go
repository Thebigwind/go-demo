package main

import (
	"fmt"
	"math/rand"
	"time"
)
import "C"

var p1 = make(chan *C.char, 3)

var str = `{"code":0,"msg":"success","result":{"checkCode":"ASD123456"}}`

func main() {
	bytes := []byte{123, 34, 99, 111, 100, 101, 34, 58, 48, 44, 34, 109, 115, 103, 34, 58, 34, 115, 117, 99, 99, 101, 115, 115, 34, 44, 34, 114, 101, 115, 117, 108, 116, 34, 58, 123, 34, 99, 104, 101, 99, 107, 67, 111, 100, 101, 34, 58, 34, 65, 83, 68, 49, 50, 51, 52, 53, 54, 34, 125, 125}

	fmt.Println(string(bytes))
}
func test3() {
	//GetRandom(2)
	p1 <- C.CString("a")
	a := <-p1
	fmt.Printf("a:%v", C.GoString(a))
}

func Test() {
	pool <- 1
	pool <- 2
	pool <- 3
	fmt.Println(Test2())
}

var pool = make(chan int, 3)

func Test2() int {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("time.Sleep...")
		case data := <-pool:
			return data
		}
	}
}
func main1() {

}

func GetRandom(max int) int {
	// 设置种子值
	rand.Seed(int64(time.Now().UnixNano()))
	// 生成随机整数
	randomInt := rand.Intn(max)
	fmt.Println(randomInt)
	return randomInt
}

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

/* 时间格式字符串转换 */
func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
