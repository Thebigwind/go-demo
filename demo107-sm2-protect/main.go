package main

import "fmt"

func main() {
	demo1()
}

func demo1() {
	certStrToCertStruct()
	fmt.Println("--------")
	keyStrToKeyStruct()
	// panic: failed to parse private key: x509: failed to parse private key (use ParsePKCS1PrivateKey instead for this key format)
}

func demo2() {
	//SM2EnvelopedKeyPairMarshal2()
}
