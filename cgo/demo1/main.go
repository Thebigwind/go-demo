package main

/*  启用CGO特性 */

import "C"
import "fmt"

var p1 = make(chan *C.char, 3)

func main() {
	p1 <- C.CString("a")
	a := <-p1
	fmt.Printf("a:%v", C.GoString(a))
}
