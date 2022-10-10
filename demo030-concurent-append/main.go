package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var s []string
	//var mu sync.Mutex = sync.Mutex{}

	for i := 0; i < 9999; i++ {
		go func() {
			//mu.Lock()
			s = append(s, "脑子进煎鱼了")
			//mu.Unlock()
		}()
	}
	time.Sleep(time.Second * 4)

	fmt.Printf("进了 %d 只煎鱼", len(s))
}

//不指定索引，动态扩容并发向切片添加数据
func concurrentAppendSliceNotForceIndex() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl = append(sl, num)
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}

// 指定索引，指定容量,并发向切片添加数据
func concurrentAppendSliceForceIndex() {
	sl := make([]int, 100)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl[num] = num
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}
