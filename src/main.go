package main

import (
	"server"
)

func main() {
	var serv server.Server
	serv.Init()
	serv.Serve()
}
