package main

import "fmt"

type A struct {
	Id     int
	Number []string
}

type B struct {
	Id     int
	Number string
}

func Deal(err error) {
	fmt.Println(err)
	if err != nil {
		fmt.Println("xxx")
	} else {
		fmt.Println("ooo")
	}
}

func test() {
	fmt.Println("")
}

func main() {
	var err error
	defer Deal(err)
	//a1 := A{Id:1,Number:[]string{"1","2"}}
	//a2 := A{Id:1,Number:[]string{"1","2"}}
	//if a1 == a2{
	//	//	fmt.Printf("")
	//	//}

	b1 := B{Id: 1, Number: "1"}
	b2 := B{Id: 1, Number: "1"}
	if b1 == b2 {

	}

	var m = map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
		"d": "dd",
		"e": "ee",
		"f": "ff",
	}
	for k, v := range m {
		fmt.Printf("k:%s,v:%s\n", k, v)
	}

	s := []string{"asong", "Golang梦工厂"}
	modifySlice(s)
	fmt.Println("inner slice: ", s)

	//s1 := []string{"asong", "Golang梦工厂"}
	//appendSlice(s1)
	//fmt.Println("inner slice: ", s1)

	m1 := map[string]int{"a": 1, "b": 2}
	modifyM(m1)
	fmt.Printf("m1:%#v", m1)

	err = fmt.Errorf("xxxxxxxxxxxx")
	Deal(err)
}

func modifyM(m map[string]int) {

	m["a"] = 11111
}

func modifySlice(s []string) {

	s[0] = "song"
	s[1] = "Golang"
	fmt.Println("out slice: ", s)
}

func appendSlice(s []string) {
	s = append(s, "快关注！！")
	fmt.Println("out slice: ", s)
}

var aa = "ockJkmHKBzfyGsHGGCztJscbFjGh2+TF0doZwmsQKvBtzdLlyUF3wTsTSDtSqxS7\nGJNgBLzcUTiCTonVmR5NYyFdGm/ySdSY8h0i5EspIwLbj7FDuqPo4dU1UBsJHu8H\nl61/xlyunE1o4EzZzJ1m8Gope+P9bcC2XdxdUCqbjcuJqhwH3YU+GafuU9G1KYmc\nBG2FJqnbehuUgjwEOC3dDNrFmF2mkxOjlMwv81Rj0yD5SzF0wsoLomeokxJ3ZTRZ\nBeWIz7q6EDgkrwiTD4q0nm/wfxVxKfl5boeQwkS6Z4bPMxwUlZmfuH1Z0ImN9eJQ\nS9i6ABCwG4ZyYosLRNBebQ=="

var bb = "Jtgc9YgyJFZHLZGc5C76kTRyxdzPyF9UKzWHzpgqO7c7ujghIr8Lh9S+c/0vRvTf\n6/eqUR4YeJSPDtblVesO6M6boRW8KvHqkgk7BFJTWyeazO7C2KT3MqodA3RYb7Fb\nDUm3VnHZmeymKzpgz0zQhK1tjH81991Nn7JMSfX59/X14tUws99XxiSe8rR9fyl5\n82LTkm89sNPYUjNDR1NTzfLVyyCXFYEPMWI0HM8ASFtsCdZLynVpC2vo1CsPjm6c\nL91TDKx8TcYflpVeYzF5mn4uqoD7WWQuZ20sHhibQi/iiZB+EorV0+JDY1oAH87O\nFGmNljEXqab57FN02zrjpg=="
