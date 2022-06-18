package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/shirou/gopsutil/mem"
	//"github.com/shirou/gopsutil/v3/mem"
)

func printinfo() {
	for {
		// fetch virtual mem info
		mem, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("some wrong", err)
		} else {
			fmt.Println((mem.String()))
		}
	}
}
func main() {
	go printinfo()
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}
