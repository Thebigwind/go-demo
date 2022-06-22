package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//读取输入
	//var str string
	//fmt.Scanln(&str)
	//str = strings.Replace(str,"\n"," ",-1)
	//fmt.Printf("str:",str)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", " ", -1)

	arr := strings.Split(str, " ")
	result := ""
	for i := len(arr) - 1; i >= 0; i-- {
		result = result + arr[i] + " "
	}
	result = strings.TrimSuffix(result, " ")
	fmt.Println(result)
}
