package main

import (
	"bytes"
	"github.com/secsy/goftp"
	"os"
	"time"
)

func main() {
	config := goftp.Config{
		User:               "jlpicard",
		Password:           "beverly123",
		ConnectionsPerHost: 10,
		Timeout:            10 * time.Second,
		Logger:             os.Stderr,
	}

	client, err := goftp.DialConfig(config, "ftp.example.com")
	if err != nil {
		panic(err)
	}

	// download to a buffer instead of file
	buf := new(bytes.Buffer)
	err = client.Retrieve("pub/interesting_file.txt", buf)
	if err != nil {
		panic(err)
	}
}
