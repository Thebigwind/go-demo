package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	DoIndex()
}

func DoIndex() {

	var a int
	var b int
	var c int
	var err error

	var wg = sync.WaitGroup{}
	wg.Add(3)

	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		defer wg.Done()
		a, err = A(ctx)
		if err != nil {
			fmt.Printf("err B")
		}
	}()
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		defer wg.Done()
		b, err = B(ctx)
		if err != nil {
			fmt.Printf("err B")
		}
	}()

	go func() {
		defer wg.Done()
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		c, err = C(ctx)
		if err != nil {
			fmt.Printf("err B")
		}
	}()

	wg.Wait()

	fmt.Printf("A:%v,B:%v,C:%v\n", a, b, c)
}

func A(ctx context.Context) (int, error) {

	sigChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 4)
		//call url
		result := 1
		sigChan <- result
	}()

	select {
	case <-ctx.Done():
		fmt.Println("A timeout")
		break
	case result := <-sigChan:
		fmt.Println("A done")
		return result, nil

	}
	return 0, nil
}

func B(ctx context.Context) (int, error) {
	sigChan := make(chan int, 1)
	go func() {
		//call url
		result := 2
		sigChan <- result
	}()

	select {
	case <-ctx.Done():
		fmt.Println("B timeout")
		break
	case result := <-sigChan:
		fmt.Println("B done")
		return result, nil

	}
	return 0, nil
}

func C(ctx context.Context) (int, error) {
	sigChan := make(chan int, 1)
	go func() {
		//call url
		result := 3
		sigChan <- result
	}()

	select {
	case <-ctx.Done():
		fmt.Println("C timeout")
		break
	case result := <-sigChan:
		fmt.Println("C done")
		return result, nil

	}
	return 0, nil
}
