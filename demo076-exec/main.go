package main

import (
	"fmt"
	"os/exec"
)

func main() {
	QskmBakup()
}

var BakupCommand = "cp /root/qkms/qkms-backend /root/qkms/qkms-backend-old"

//备份
func QskmBakup() error {
	cmd := exec.Command("/bin/bash", "-c", BakupCommand)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Printf("err:%v", err.Error())
		return err
	}
	fmt.Printf("success:%v", string(bytes))

	return nil
}
