package main

import (
	"fmt"
	"reflect"
)

//反射 使用字符串函数名，调用函数
type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

func main() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
