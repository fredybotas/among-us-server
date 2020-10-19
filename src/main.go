package main

import (
	"server"
)

func main() {
	var serv server.UDPServer
	serv.Init()
	serv.Serve()
}
