package common

import (
	"fmt"
	"github.com/mr-tron/base58"
	"log"
	"math/big"
)

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

func Base58Decode(encoded string) ([]byte, error) {
	decoded := big.NewInt(0)
	base := big.NewInt(58)

	for _, char := range encoded {
		index := big.NewInt(int64(base58AlphabetIndex(byte(char))))
		decoded.Mul(decoded, base)
		decoded.Add(decoded, index)
	}

	return decoded.Bytes(), nil
}

func base58AlphabetIndex(char byte) int {
	for i, c := range base58Alphabet {
		if byte(c) == char {
			return i
		}
	}
	return -1
}

func testbase58() {
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
}
