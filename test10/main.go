package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	//select {}
	fmt.Println("test done")
	for {
		time.Sleep(time.Second)
	}
}

func test() {
	fmt.Println("test start")
	go doSome()
	time.Sleep(time.Second * 3)
	fmt.Println("test end")
}
func doSome() {
	fmt.Println("xxxxx")
	time.Sleep(10 * time.Second)
	fmt.Println("oooooo")
}
