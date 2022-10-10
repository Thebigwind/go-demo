package main

import (
	"bytes"
	"fmt"
)

func BytesCombine1(pBytes ...[]byte) []byte {
	length := len(pBytes)
	s := make([][]byte, length)
	for index := 0; index < length; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

func BytesCombine2(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func main() {
	fmt.Println(BytesCombine1([]byte{1, 2, 3}, []byte{4, 5, 6}))
	fmt.Println(BytesCombine1([]byte("one"), []byte("two")))
	fmt.Printf("%d\n", BytesCombine1([]byte{1, 2, 3}, []byte{4, 5, 6}))
	fmt.Printf("%s\n", BytesCombine1([]byte("one"), []byte("two")))
	fmt.Printf("xx%s\n", BytesCombine2([]byte{1, 2, 3}, []byte{4, 5, 6}))
}
