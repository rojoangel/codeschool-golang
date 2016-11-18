package main

import "chat/lib"

func main() {
	notifiers := lib.GetAllConnections()
	for _, c := range notifiers {
		c.Notify()
	}
}
