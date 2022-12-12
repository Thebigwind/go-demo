package main

import (
	"math/rand"
	"time"
)

var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func RandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}
