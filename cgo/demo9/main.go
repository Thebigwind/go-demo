package main

/*
#include <stdlib.h>

typedef struct CK_INFO_ {
    char sn[32];
} CK_INFO;

int GetCKInfo(CK_INFO** info_list, unsigned int* info_len);
*/
import "C"
import (
	"unsafe"
)

type ckInfo struct {
	sn [32]C.char
}

func main() {
	var infoList **C.CK_INFO
	var infoLen C.uint
	res := C.GetCKInfo(&infoList, &infoLen)
	if res != 0 {
		// handle error
	}
	defer C.free(unsafe.Pointer(infoList))

	// convert to slice of strings
	var snList []string
	for i := 0; i < int(infoLen); i++ {
		info := (*ckInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(infoList)) + uintptr(i)*unsafe.Sizeof(*infoList)))
		snList = append(snList, C.GoString(&info.sn[0]))
	}

	// use snList as needed
}
