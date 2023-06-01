package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello 嗨客网(www.haicoder.net)")
	//Golang 实现 string 转 float64
	var srtScore = "12.00"
	score, err := strconv.ParseFloat(srtScore, 64)
	fmt.Printf("score = %.2f Err = %v", score, err)
	fmt.Println("score = ", score, "Err = ", err)
}
