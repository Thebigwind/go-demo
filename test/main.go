package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type A struct {
	Id   int
	Name string
}

func main() {

	//wg := sync.WaitGroup{}
	//wg.Add(10)
	//var lock sync.Mutex
	//for i := 0; i < 10; i++ {
	//	lock.Lock()
	//	go func(index int) {
	//		defer wg.Done()
	//		test2(index)
	//	}(i)
	//	lock.Unlock()
	//}
	//wg.Wait()
	//fmt.Println((1672693771 - 1672643371)/86400+1)
	//fmt.Printf("aa:%s\n",path.Dir("/aa/bb/cc"))
	//
	//if err := os.MkdirAll("/home/aa", os.ModePerm); err != nil {
	//	fmt.Printf("os.MkdirAll err:%s", err.Error())
	//}else{
	//	fmt.Println("success")
	//}
	//
	//return
	//a1 := &A{
	//	Id:   1,
	//	Name: "aa1",
	//}
	//
	//a2 := &A{
	//	Id:   2,
	//	Name: "aa2",
	//}
	//
	//a3 := &A{
	//	Id:   3,
	//	Name: "aa3",
	//}
	//arr := make([]*A, 0)
	//arr = append(arr, a1, a2, a3)
	//for i, v := range arr {
	//	v.Id = i + 5
	//	fmt.Printf("arr:%v", v)
	//}
	//
	//fmt.Printf(toBinaryRunes("abcdef"))
}
func toBinaryRunes(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		fmt.Fprintf(&buffer, "%b", runeValue)
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}

var Count int

func test1(index int) {

	fmt.Printf("lock: index:%d\n", index)
	time.Sleep(time.Second * 3)
	Count = Count + 1
	fmt.Printf("Count:%v", Count)

	fmt.Printf("unlock:index:%d\n", index)
}

//
func test2(index int) {
	var lock sync.Mutex
	lock.Lock()
	fmt.Printf("lock: index:%d\n", index)
	time.Sleep(time.Second * 3)
	Count = Count + 1
	fmt.Printf("Count:%v", Count)
	lock.Unlock()
	fmt.Printf("unlock:index:%d\n", index)
}
