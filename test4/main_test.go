package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseCert(t *testing.T) {
	certBytes, err := os.ReadFile("sign.crt")
	got, err := ParseCert(string(certBytes))
	if err != nil {
		t.Errorf("ParseCert() error = %v, wantErr %v", err)
		return
	}
	fmt.Printf("got:%v", got)
}
