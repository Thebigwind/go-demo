package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ExecShell(s string) (error, string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("exec_shell:", err.Error())
		return err, ""
	}
	//fmt.Printf("%s", out.String())
	return err, strings.Trim(out.String(), "\n")
}
