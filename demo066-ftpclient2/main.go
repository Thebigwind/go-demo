package main

import (
	"os"

	"github.com/secsy/goftp"
)

func main() {
	// Create client object with default config
	client, err := goftp.Dial("ftp.example.com")
	if err != nil {
		panic(err)
	}

	// Download a file to disk
	readme, err := os.Create("readme")
	if err != nil {
		panic(err)
	}

	err = client.Retrieve("README", readme)
	if err != nil {
		panic(err)
	}

	// Upload a file from disk
	bigFile, err := os.Open("big_file")
	if err != nil {
		panic(err)
	}

	err = client.Store("big_file", bigFile)
	if err != nil {
		panic(err)
	}
}
