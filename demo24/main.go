package main

import (
	"encoding/json"
	//"errors"
	"fmt"
)

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
