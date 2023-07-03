// hello.go
package main

/*
#include <stdlib.h>
#include <stdio.h>

typedef struct users_info_ {
    char name[32];
} users_info;

int GetUsers(users_info** users, unsigned int* users_len);

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct users_info_ {
    char name[32];
} users_info;

static users_info user_list[] = {
    {"Alice"},
    {"Bob"},
    {"Charlie"}
};
static const unsigned int user_list_len = sizeof(user_list) / sizeof(user_list[0]);

int GetUsers(users_info** users, unsigned int* users_len) {
    if (users == NULL || users_len == NULL) {
        return -1;
    }
    *users = (users_info*)malloc(user_list_len * sizeof(users_info));
    if (*users == NULL) {
        return -2;
    }
    memcpy(*users, user_list, user_list_len * sizeof(users_info));
    *users_len = user_list_len;
    return 0;
}

int main() {
    users_info* users;
    unsigned int users_len;
    int res = GetUsers(&users, &users_len);
    if (res != 0) {
        printf("Failed to get user list: %d\n", res);
        return 1;
    }
    for (unsigned int i = 0; i < users_len; i++) {
        printf("User %d: %s\n", i+1, users[i].name);
    }
    free(users);
    return 0;
}

*/
import "C"
import (
	"unsafe"
)

//在import "C"语句前的注释中可以通过#cgo语句设置编译阶段和链接阶段的相关参数。编译阶段的参数主要用于定义相关宏和指定头文件检索路径。链接阶段的参数主要是指定库文件检索路径和要链接的库文件。

type users_info struct {
	name [32]C.char
}

func main() {
	// initialize users and users_len
	var users **C.users_info
	var len C.uint
	res := C.GetUsers(&users, &len)
	if res != 0 {
		// handle error
	}
	defer C.free(unsafe.Pointer(users))

	// mock data
	var data []*users_info
	for i := 0; i < int(len); i++ {
		info := (*users_info)(unsafe.Pointer(uintptr(unsafe.Pointer(users)) + uintptr(i)*unsafe.Sizeof(*users)))
		data = append(data, info)
	}

	// use data as needed
}
