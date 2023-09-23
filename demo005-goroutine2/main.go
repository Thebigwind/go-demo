package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func processSlice(ctx context.Context, id int, slice []int, timeout time.Duration, done chan<- bool) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for _, num := range slice {
		select {
		case <-ctxWithTimeout.Done():
			fmt.Printf("Goroutine %d: Timeout\n", id)
			done <- true
			return
		default:
			// Process the element here (for example, print it)
			fmt.Printf("Goroutine %d: Element: %d\n", id, num)
		}
	}

	done <- true
}

func main() {
	slices := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	done := make(chan bool)
	ctx := context.Background()

	for i, slice := range slices {
		timeout := 2 * time.Second
		go processSlice(ctx, i+1, slice, timeout, done)
	}

	// Wait for all goroutines to finish or timeout
	for i := 0; i < len(slices); i++ {
		select {
		case <-done:
			fmt.Printf("Goroutine %d completed successfully\n", i+1)
		case <-time.After(2 * time.Second):
			fmt.Printf("Goroutine %d timed out\n", i+1)
		}
	}

	fmt.Println("All Goroutines finished.")
}

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) { test(t, doBadthing) }
