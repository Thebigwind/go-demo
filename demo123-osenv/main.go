package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/mr-tron/base58"
)

func main() {
	// Data to encode
	data := "Hello, Base58!!"

	// Encode the data
	encoded := base58.Encode([]byte(data))
	fmt.Println("Encoded:", encoded)

	// Decode the data
	decoded, err := base58.Decode(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded:", string(decoded))
	fmt.Println(Base58Encode([]byte(data)))
}

func test() {
	err := os.Setenv("DRIVER_TYPE", "soft")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	value := os.Getenv("DRIVER_TYPE")
	fmt.Printf("DRIVER_TYPE=%s", value)
}

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func Base58Encode(data []byte) string {
	decoded := new(big.Int).SetBytes(data)
	encoded := ""

	base := big.NewInt(58)
	zero := big.NewInt(0)

	for decoded.Cmp(zero) > 0 {
		mod := new(big.Int)
		decoded.DivMod(decoded, base, mod)
		encoded = string(base58Alphabet[mod.Int64()]) + encoded
	}

	return encoded
}
