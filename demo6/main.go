//启动 2个groutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完
// 思路：采用ctx, _ := context.WithTimeout(context.Background(), time.Second*2)实现2s取消。协程执行完后通过channel通知，是否超时。

package main

import (
	"context"
	"fmt"
	"time"
)

//1秒执行完，参数为一个channel
func f1(in chan struct{}) {

	time.Sleep(1 * time.Second)
	in <- struct{}{}

}

//2秒执行完，参数为一个channel
func f2(in chan struct{}) {
	time.Sleep(3 * time.Second)
	in <- struct{}{}
}

func test() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	//定义context的超时时间
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		go f1(ch1)
		select {
		case <-ctx.Done():
			fmt.Println("f1 timeout")
			break
		case <-ch1:
			fmt.Println("f1 done")
		}
	}()

	go func() {
		go f2(ch2)
		select {
		case <-ctx.Done():
			fmt.Println("f2 timeout")
			break
		case <-ch2:
			fmt.Println("f2 done")
		}
	}()
	time.Sleep(time.Second * 5)
}
