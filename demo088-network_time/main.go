package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().In(loc)
	fmt.Println("Shanghai time:", t)
}
