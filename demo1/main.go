package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	resultChan := make(chan struct{}, 1)
	//通过 context.WithCancel，返回一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())
	//起多个 context
	for i := 0; i < 10; i++ {
		go t(ctx, resultChan, i)
	}
	//主goroutine定义超时时间
	tim := time.NewTimer(time.Second * 3)
	defer tim.Stop()
	//当主goroutine超时了，或者某个子goroutine找到结果，向 resultChan 发送了结果通知， 调用cancel(),取消子goroutine的执行
	select {
	case <-tim.C:
		fmt.Println("all goroutine timeout, start cancel...")
		cancel()
	case <-resultChan:
		fmt.Println("allready find the result, start cancel...")
		cancel()
	}

	fmt.Println("main end...")
}

func t(ctx context.Context, resultChan chan struct{}, index int) {
	//在default里做逻辑处理，如果找到结果就向 resultChan 中发送结果通知
	//如果接受到主goroutine的退出信号，则退出执行
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel()....")
			return
		default:
			fmt.Printf("index %d is looking for...\n", index)
			if index == 2 {
				time.Sleep(time.Second * 4)
				fmt.Println("index 2 find the result, done")
				resultChan <- struct{}{}
				return
			} else {
				time.Sleep(time.Second * 6)
			}

			//if FindTarget(){
			//	resultChan <- struct{}{}
			//}
			//return
		}
	}
}
