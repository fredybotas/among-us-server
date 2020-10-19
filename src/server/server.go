package server

import (
	"fmt"
	"net"
)

const port int = 1221
const ip string = "127.0.0.1"

type UDPServer struct {
	socket *net.UDPConn
}

func (server *UDPServer) Init() {
	fmt.Printf("Starting server\n")
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	sock, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error binding sock %v\n", err)
		return
	}
	server.socket = sock
}
