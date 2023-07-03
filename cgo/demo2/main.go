package main

import "C"
import (
	"fmt"
	"time"
)

func main() {
	// 创建一个缓冲区大小为 3 的 chan
	ch := make(chan *C.char, 3)

	// 启动一个 goroutine 来往 chan 中写入数据
	go func() {
		for i := 0; i < 3; i++ {
			s := fmt.Sprintf("Message %d", i)
			ch <- C.CString(s)
			time.Sleep(time.Second)
		}
	}()

	// 从 chan 中读取数据并打印
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println(C.GoString(msg))
		//C.free(unsafe.Pointer(msg))
	}
}
