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

func SaveBlockedIPs(ips []string) {
	//... saves IPs
}

func GetHostIP() string {
	// returns host IP
	return "192.168.1.1"
}

func GetGuestIP() string {
	// returns guest IP
	return "192.168.1.2"
}
