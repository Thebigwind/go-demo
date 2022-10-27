package main

import (
	"bytes"
	"fmt"
)

type A struct {
	Id   int
	Name string
}

func main() {
	a1 := &A{
		Id:   1,
		Name: "aa1",
	}

	a2 := &A{
		Id:   2,
		Name: "aa2",
	}

	a3 := &A{
		Id:   3,
		Name: "aa3",
	}
	arr := make([]*A, 0)
	arr = append(arr, a1, a2, a3)
	for i, v := range arr {
		v.Id = i + 5
		fmt.Printf("arr:%v", v)
	}

	fmt.Printf(toBinaryRunes("abcdef"))
}
func toBinaryRunes(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		fmt.Fprintf(&buffer, "%b", runeValue)
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}
