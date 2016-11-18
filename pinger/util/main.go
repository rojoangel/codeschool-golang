package main

import (
	"fmt"
	"pinger/util"
)

func main() {
	isAlive := false
	for !isAlive {
		isAlive = util.PingChatServer()
	}

	fmt.Println("Ping is done")
}
