package main

import (
	"flag"
	"testproj/socket/client"
	"testproj/socket/server"
)

func main() {
	c := flag.Int("c", 1, "client0 or server 1")
	flag.Parse()
	if *c == 0 {
		client.Dial()
	} else {
		server.StartServer()
	}
}
