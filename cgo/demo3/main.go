// hello.go
package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("你好 Hello, World\n"))
}
