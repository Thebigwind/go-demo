package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func HandleText(textfile string) error {
	file, err := os.Open(textfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // or
		//line := scanner.Bytes()

		//do_your_function(line)
		//fmt.Printf("%s\n", line)
		Deal(line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
		return err
	}

	return nil
}
func main() {
	//fmt.Println(54/10*10)
	if err := HandleText("/Users/me/Thebigwind/lal/rtppacket-outout4.log"); err != nil {
		return
	}

	PrintResult(AnalyzeInnerMap, AnalyzeBetweenMap, Count, "96 type 统计结果")
	PrintResult(Analyze97InnerMap, Analyze97BetweenMap, Count97, "97 type 统计结果")
}

var AnalyzeInnerMap = map[int64]int{}
var AnalyzeBetweenMap = map[int64]int{}

type Data struct {
	Marker    string
	TimeStamp string
	Type      string
}

var preData = Data{}
var pre97Data = Data{}

var Count = 0
var Count97 = 0

var Analyze97InnerMap = map[int64]int{}
var Analyze97BetweenMap = map[int64]int{}

func Deal(str string) {

	validData := str[49 : len(str)-1]
	//fmt.Printf("%s\n", strData)

	TimeStamp := validData[4:20]
	marker := validData[len(validData)-1:]
	packetType := validData[len(validData)-11 : len(validData)-9]
	//fmt.Printf("packetType:%v",packetType)
	//os.Exit(1)
	if packetType == "96" {
		Count = Count + 1
		if preData.Marker == "" {
			preData.Marker = marker
			preData.TimeStamp = TimeStamp
		} else {
			//当前包时间
			currentTimeStamp, _ := strconv.ParseInt(TimeStamp, 10, 64)
			//前一个包时间
			preTimeStamp, _ := strconv.ParseInt(preData.TimeStamp, 10, 64)

			if currentTimeStamp-preTimeStamp < 0 {
				//fmt.Printf(str +"\n")
				Count = Count - 1
				return
			}
			//帧内
			if preData.Marker == "0" {
				timeC := int64(0)
				if currentTimeStamp-preTimeStamp > 100 {
					timeC = (currentTimeStamp - preTimeStamp) / 100 * 100
				} else {
					timeC = (currentTimeStamp - preTimeStamp) / 10 * 10
				}

				_, ok := AnalyzeInnerMap[timeC]
				if !ok {
					AnalyzeInnerMap[timeC] = 1
				} else {
					AnalyzeInnerMap[timeC] = AnalyzeInnerMap[timeC] + 1
				}
			}

			if preData.Marker == "1" { //帧之间
				timeC := (currentTimeStamp - preTimeStamp) / 1000 //转化为毫秒
				if timeC > 10 {
					timeC = timeC / 10 * 10
				}

				_, ok := AnalyzeBetweenMap[timeC]
				if !ok {
					AnalyzeBetweenMap[timeC] = 1
				} else {
					AnalyzeBetweenMap[timeC] = AnalyzeBetweenMap[timeC] + 1
				}
			}

			//保存数据
			preData.Marker = marker
			preData.TimeStamp = TimeStamp
		}
	}

	if packetType == "97" {
		Count97 = Count97 + 1
		if pre97Data.Marker == "" {
			pre97Data.Marker = marker
			pre97Data.TimeStamp = TimeStamp
		} else {
			//当前包时间
			currentTimeStamp, _ := strconv.ParseInt(TimeStamp, 10, 64)
			//前一个包时间
			preTimeStamp, _ := strconv.ParseInt(pre97Data.TimeStamp, 10, 64)

			if currentTimeStamp-preTimeStamp < 0 {
				//fmt.Printf(str +"\n")
				Count97 = Count97 - 1
				return
			}
			//帧内
			if pre97Data.Marker == "0" {
				timeC := int64(0)
				if currentTimeStamp-preTimeStamp > 100 {
					timeC = (currentTimeStamp - preTimeStamp) / 100 * 100
				} else {
					timeC = (currentTimeStamp - preTimeStamp) / 10 * 10
				}

				_, ok := Analyze97InnerMap[timeC]
				if !ok {
					Analyze97InnerMap[timeC] = 1
				} else {
					Analyze97InnerMap[timeC] = Analyze97InnerMap[timeC] + 1
				}
			}

			if pre97Data.Marker == "1" { //帧之间
				timeC := (currentTimeStamp - preTimeStamp) / 1000 //转化为毫秒
				if timeC > 10 {
					timeC = timeC / 10 * 10
				}

				_, ok := Analyze97BetweenMap[timeC]
				if !ok {
					Analyze97BetweenMap[timeC] = 1
				} else {
					Analyze97BetweenMap[timeC] = Analyze97BetweenMap[timeC] + 1
				}
			}

			//保存数据
			pre97Data.Marker = marker
			pre97Data.TimeStamp = TimeStamp
		}
	}
}

func qsort(a []int64) []int64 {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	//已最右侧的元素为分区点
	for i := range a {
		if a[i] < a[right] {
			//将a[i]放到左侧分区
			a[left], a[i] = a[i], a[left]
			left++ //左侧分区+1
		}
	}
	//将分区点放到合适的位置
	a[left], a[right] = a[right], a[left]
	//递归
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func PrintResult(AnalyzeInnerMap map[int64]int, AnalyzeBetweenMap map[int64]int, Count int, desc string) {
	//打印结果
	//总包数
	fmt.Println(desc)
	fmt.Printf("packet Number:%v\n", Count)

	innerArr := []int64{}
	//帧内
	for k := range AnalyzeInnerMap {
		innerArr = append(innerArr, k)
	}

	qsort(innerArr)

	for _, v := range innerArr {
		fmt.Printf("帧内packet延迟时间（微秒）:%v,:帧内packet数量：%v\n", v, AnalyzeInnerMap[v])
	}

	fmt.Printf("----------\n")
	btweenArr := []int64{}
	//帧间
	for k := range AnalyzeBetweenMap {
		btweenArr = append(btweenArr, k)
	}
	qsort(btweenArr)
	for _, v := range btweenArr {
		fmt.Printf("帧间packet延迟时间（毫秒）:%v,帧间packet数量:%v\n", v, AnalyzeBetweenMap[v])
	}
	fmt.Println("==================")
}
