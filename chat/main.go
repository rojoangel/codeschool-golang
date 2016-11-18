package main

import (
	"chat/util"
	"errors"
	"fmt"
	"os"
	"time"
)

type guestConnection struct {
	ip       string
	userName string
}

func (g guestConnection) isAllowed() bool {
	return !isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

func main() {
	ip := util.GetGuestIP()
	userName := "Kerry"

	gConn := guestConnection{ip: ip, userName: userName}
	isAllowedStatus := gConn.isAllowed()
	fmt.Println(isAllowedStatus)
}

/*
func main() {
	hostIP := util.GetHostIP()
	fmt.Println(isIPBlocked(hostIP))
}
*/

func isIPBlocked(ip string) bool {
	blockedIPs := getBlockedIPs()
	for _, blockedIP := range blockedIPs {
		if ip == blockedIP {
			return true
		}
	}
	return false
}

func getBlockedIPs() []string {
	return []string{"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}

/*
func main() {

	args, err := readArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) > 1 {
		hostIP := args[1]
		util.RunGuest(hostIP)
	} else {
		listenIP := util.GetLocalNetworkIP()
		listenPort := getListenPort()
		util.RunHost(listenIP + ":" + listenPort)
	}
}
*/

func addToBlockedList(ips []string) {
	util.SaveBlockedIPs(ips)
}

func readArgs() ([]string, error) {
	args := os.Args
	if len(args) > 1 && len(args[1]) > 15 {
		err := errors.New("Too many arguments")
		return args, err
	}
	return args, nil
}

func getListenPort() string {
	hourOfDay := time.Now().Hour()
	if hourOfDay < 12 {
		return "8080"
	} else if hourOfDay < 20 {
		return "8081"
	} else {
		return "8082"
	}
}
