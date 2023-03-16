package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func Test3() {
	ntpTime, err := ntp.Time("pool.ntp.org") //time.apple.com
	if err != nil {
		fmt.Println(err)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Println("----------------------")
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("+++++++++++++++++++++++++++++++")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

	// so from here, use ntpTime instead of time.Now().Local()
	// or set your machine time to sync with a network time server
	// for example, on Mac OSX, you can tick to enable auto sync
	// with time.apple.com
}
