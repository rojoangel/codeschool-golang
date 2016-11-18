package util

import "fmt"

func RunGuest(ipPort string) {
	fmt.Println("Running guest", ipPort)
}

func RunHost(ipPort string) {
	fmt.Println("Running host", ipPort)
}

func GetLocalNetworkIP() string {
	return "192.168.1.0"
}

func SaveBlockedIPs(ips [4]string) {
	//... saves IPs
}
