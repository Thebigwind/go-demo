package main

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (string, error) {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// 遍历接口
	for _, iface := range interfaces {
		// 排除回环接口
		if iface.Flags&net.FlagLoopback == 0 {
			// 获取接口的地址
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			// 遍历地址
			for _, addr := range addrs {
				// 检查是否为 IP 地址
				if ipnet, ok := addr.(*net.IPNet); ok {
					// 检查是否为 IPv4 地址
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("No non-loopback IP address found")
}

func main() {
	ip, err := GetLocalIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Local IP:", ip)
}

func PublicReplace(pubkey string) string {
	res := strings.Replace(pubkey, "\n", "", -1)
	res = strings.Replace(res, " ", "", -1)
	return res
}

func ComparePubkey(pub1 string, pub2 string) bool {
	pub1 = PublicReplace(pub1)
	pub2 = PublicReplace(pub2)
	return pub1 == pub2
}
