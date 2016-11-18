package lib

import (
	"fmt"
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
	Notify()
}

func (g guestConnection) Notify() {
	fmt.Println("Guest connection from user name:", g.userName)
}

func (v visitorConnection) Notify() {
	fmt.Println("Visitor connected at:", v.connHour)
}

func GetAllConnections() []notifier {
	gConn := &guestConnection{ip: "192.168.0.10", userName: "Darth Vader"}
	vConn := &visitorConnection{ip: "192.168.0.11", connHour: time.Now().Hour()}

	return []notifier{gConn, vConn}
}
