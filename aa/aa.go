package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func isAlphaNumeric(s string) bool {
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", s)
	if err != nil {
		// 处理错误
		fmt.Println("Error:", err)
		return false
	}
	return matched
}

func main() {
	tests := []string{"123abc", "abc123", "abc", "123", "abc@", "123$", "abc#123"}
	for _, test := range tests {
		fmt.Printf("%s: %t\n", test, isAlphaNumeric(test))
	}

	time, err := UnixToTime("20241220011840")
	if err != nil {
		fmt.Printf("err:%v", err)
	} else {
		fmt.Printf("time:%v", time)
	}

}

func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}
