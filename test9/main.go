package main

import (
	"fmt"
	"strings"
)

func main() {
	a1 := "1\n2345 \n"
	a2 := "1 2\n3 45"
	fmt.Println(ComparePubkey(a1, a2))
}
func PublicReplace(pubkey string) string {
	res := strings.Replace(pubkey, "\n", "", -1)
	res = strings.Replace(res, " ", "", -1)
	return res
}

func ComparePubkey(pub1 string, pub2 string) bool {
	pub1 = PublicReplace(pub1)
	pub2 = PublicReplace(pub2)
	return pub1 == pub2
}
