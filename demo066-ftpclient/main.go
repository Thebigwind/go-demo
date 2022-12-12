package main

import (
	"fmt"
	"github.com/secsy/goftp"
)

const (
	ftpServerURL  = "ftp.us.debian.org"
	ftpServerPath = "/debian/"
)

func main() {
	client, err := goftp.Dial(ftpServerURL)
	if err != nil {
		panic(err)
	}
	files, err := client.ReadDir(ftpServerPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
