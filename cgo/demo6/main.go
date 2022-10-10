package main

//#include <stdio.h>
import "C"
import (
	"fmt"
)

func copy_string(str string) {

	//var cstr *C.char
	//cstr = C.CString(str)
	//defer C.free(unsafe.Pointer(cstr))

}

func copy_string2(str string) {
	c_id := [128]C.char{}
	for i := 0; i < len(str) && i < 127; i++ {
		c_id[i] = C.char(str[i])
	}

	fmt.Println(c_id)
}

func main() {
	//copy_string("abcde")
	copy_string2("abcde")
}
