package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path"
)

func main() {
	dirpath := "/Users/me/zdlz/qkms-backend/third_party/libs/amd64"

	infos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}

	for _, v := range infos {
		if v.IsDir() {
			continue
		}
		filepath := path.Join(dirpath, v.Name())

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf("Error reading:%s", err.Error())
			return
		}
		fingerPrint := Sha256(string(contents))
		fmt.Printf("filepath:%s,print:%s\n", filepath, fingerPrint)
	}
}

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
