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
	isAdmin  bool
}

type visitorConnection struct {
	ip       string
	connHour int
}

type notifier interface {
	notify()
}

func (g guestConnection) isAllowed() bool {
	return !isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

// one implementation for guestConnection
func (g guestConnection) notify() {
	fmt.Println("Guest connection from user name:", g.userName)
}

// and a different implementation for visitorConnection
func (v visitorConnection) notify() {
	fmt.Println("Visitor connected at:", v.connHour)
}

func main() {
	notifiers := getAllConnections()
	for _, c := range notifiers {
		c.notify()
	}

}

func getAllConnections() []notifier {
	gConn := &guestConnection{ip: "192.168.0.10", userName: "Darth Vader"}
	vConn := &visitorConnection{ip: "192.168.0.11", connHour: time.Now().Hour()}

	return []notifier{gConn, vConn}
}

/*
func main() {
	ip := util.GetGuestIP()
	userName := "Obi-Wan"

	gConn := &guestConnection{ip: ip, userName: userName}
	fmt.Println("Before auth", gConn)
	authorizeAdmin(gConn)
	fmt.Println("After auth", gConn)
}
*/

func authorizeAdmin(c *guestConnection) {
	if c.isAllowed() && c.ip == "192.168.0.10" {
		c.isAdmin = true
	}
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
