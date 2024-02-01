package main

import (
	"fmt"
	"strconv"
)

type Callback func(msg string)

//将字符串转换为int64，如果转换失败调用Callback
func stringToInt(s string, callback Callback) int64 {
	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		callback(err.Error())
		return 0
	} else {
		return value
	}
}

// 记录日志消息的具体实现
func errLog(msg string) {
	fmt.Println("Convert error: ", msg)
}

func main() {
	fmt.Println(stringToInt("18", errLog))
	fmt.Println(stringToInt("hh", errLog))
}
