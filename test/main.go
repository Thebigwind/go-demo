package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	value := "365"
	now := time.Now().Unix()
	interval, _ := strconv.ParseInt(value, 10, 64)
	fmt.Printf("now:%s\n", now)
	upTime := "2022-09-30 12:10:16"
	fmt.Printf("now:%v", now)
	fmt.Printf("upTime:%v", upTime)

	fmt.Printf("diff:%v\n", now-TimeStringToGoTime(upTime).Unix())
	fmt.Printf("interval:%v\n", interval)
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
