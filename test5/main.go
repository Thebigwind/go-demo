package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func test() {
	numGoroutines := 10
	timeout := 2 * time.Second

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			select {
			case <-time.After(timeout):
				fmt.Printf("Goroutine %d timed out\n", id)
			case <-time.After(1 * time.Second): // Simulating some work
				fmt.Printf("Goroutine %d completed\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished")
}

func t() {
	parse, err := time.Parse("2006-01-02 15:04:05", "2023-09-03 15:40:03")
	if err != nil {
		fmt.Errorf("send device info ts type error,err:%s,ts:%v", err.Error(), "xx")

	}
	fmt.Println(parse)
}

func main() {
	t()
	//YYYYMMDDHHMMSS24 := "2006-01-02 15:04:05"
	//days, err := GetDaysBetween2Date(YYYYMMDDHHMMSS24, "2023-03-06 00:00:00", "2023-03-05 01:00:00")
	//if err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}
	//fmt.Println(days)
	//test2()
}
func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
	// 将字符串转化为Time格式
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0, err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0, err
	}
	fmt.Printf("xxxx:%v\n", date1.Sub(date2).Hours())
	//计算相差天数
	days := int(math.Ceil(date1.Sub(date2).Hours() / 24))
	return days, nil
}

func test2() {
	sevenDaysAgo := time.Now().AddDate(0, 0, -6)
	sevenDaysAgoMidnight := time.Date(sevenDaysAgo.Year(), sevenDaysAgo.Month(), sevenDaysAgo.Day(), 0, 0, 0, 0, sevenDaysAgo.Location())
	fmt.Printf("sss:%v", sevenDaysAgoMidnight)
	//a := time.Now().AddDate(0, 0, -6).Format("")
	ts := time.Now().AddDate(0, 0, -6)
	if ts.Before(sevenDaysAgoMidnight) {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}
