package main

/*
#include <stdlib.h>

typedef struct users_info_ {
    char name[32];
} users_info;

int GetUsers(users_info** users, unsigned int* users_len);
*/
import "C"
import (
	"unsafe"
)

func main() {
	var users **C.users_info
	var users_len C.uint
	res := C.GetUsers(&users, &users_len)
	if res != 0 {
		// handle error
	}
	defer C.free(unsafe.Pointer(users))

	// convert to slice of strings
	var userNames []string
	for i := 0; i < int(users_len); i++ {
		user := (*C.users_info)(unsafe.Pointer(uintptr(unsafe.Pointer(users)) + uintptr(i)*unsafe.Sizeof(*users)))
		userNames = append(userNames, C.GoString(&user.name[0]))
	}

	// use userNames as needed
}
