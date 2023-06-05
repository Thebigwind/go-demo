package main

import "fmt"

func main() {
	var num int32 = 42
	ptr := &num

	fmt.Println("Value:", num)
	fmt.Println("Pointer:", *ptr)

}
