package main

import (
	"chat/util"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	blockedIPs := []string{"192.168.0.19", "192.168.0.20"}
	addToBlockedList(blockedIPs)
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
