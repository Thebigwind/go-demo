package main

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str := "Hello, world!"
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	voidPtr := unsafe.Pointer(cstr)
	fmt.Printf("%v\n", voidPtr)
}
