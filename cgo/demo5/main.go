// hello.go
package main

import "C"

import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}

func main() {
	SayHello(C.GoString("abcdef"))
}
