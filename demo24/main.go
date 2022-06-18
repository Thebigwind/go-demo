package main

import (
	"encoding/json"
	//"errors"
	"fmt"
)

/*
系统自身的error处理一般是 errors.New()或fmt.Errorf()等，对一些需要复杂显示的，不太友好，我们可以扩展下error。
error在标准库中被定义为一个接口类型，该接口只有一个Error()方法
*/
//type error interface {
//	Error() string
//}
type Err struct {
	Code int
	Msg  string
}

//自定义error只要拥有Error()方法，就实现了error接口
func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func (e *Err) New(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

func main() {
	err := Err{}
	fmt.Println(err.New(401, "无此权限"))
}
