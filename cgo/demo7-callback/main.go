package main

// C语言中的 void *和 Go中的unsafe.Pointer对应
// C语言中的函数和 Go中的函数可以通过 //export NAME的方式来建立对应关系

/*
#include "api.h"
extern void cgoCall(void *, int); // 这里与api.h中的IntCallback保持类型一致
*/
import "C"
import "unsafe"

type Caller interface {
	Call(int)
}

//此处省略上述重复代码

//export cgoCall
func cgoCall(p unsafe.Pointer, number C.int) {
	caller := *(*Caller)(p)
	caller.Call(int(number))
}

//此处省略上述重复代码

type OneCaller struct{}
type AnotherCaller struct{}

func (o OneCaller) Call(value int) {
	println("one:", value)
}

func (o AnotherCaller) Call(value int) {
	println("another:", value)
}

func SetCallback(caller Caller) {
	C.SetIntCallback(C.IntCallback(C.cgoCall), unsafe.Pointer(&caller))
}

func DoCallback(value int) {
	C.DoIntCallback(C.int(value))
}

func main() {
	one := OneCaller{}

	SetCallback(one)
	DoCallback(1234)
	another := AnotherCaller{}

	SetCallback(another)
	DoCallback(5678)
}
