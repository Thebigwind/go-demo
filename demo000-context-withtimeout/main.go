package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go dosometing(ctx, ch)
	select {
	case <-ctx.Done():
		fmt.Println("超时：", ctx.Err())
	case result := <-ch:
		fmt.Printf("result:%v", result)
	}
	time.Sleep(100 * time.Second)
}

func dosometing(ctx context.Context, ch chan int) {
	fmt.Println("do sleep..")
	time.Sleep(3 * time.Second)
	fmt.Println("wake up")
	ch <- 1
}

func valueTest() {
	ctx := context.WithValue(context.Background(), "id", "123045")
	Get(ctx, "id")
	go func() {
		ctx2 := context.WithValue(ctx, "name", "lai")
		Get(ctx2, "id")
		go func() {
			Get(ctx2, "name")
		}()
	}()

	time.Sleep(1 * time.Second)

}

func Get(ctx context.Context, k string) {
	if v, ok := ctx.Value(k).(string); ok {
		fmt.Println(v)
	}
}
