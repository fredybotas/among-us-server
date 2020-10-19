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

func (server *UDPServer) onDataReceive(dataReceived []byte, receivedFrom *net.UDPAddr) {
	_, err := server.socket.WriteToUDP([]byte("Test response"), receivedFrom)
	if err != nil {
		fmt.Print("Error while responding: %v", err)
	}
}

func (server *UDPServer) Serve() {
	var p [1024]byte
	for {
		n, remoteAddress, err := server.socket.ReadFromUDP(p[:])
		if err != nil {
			fmt.Printf("Error receiving data  %v", err)
			continue
		}
		fmt.Printf("Received data from %v with length %d\n", remoteAddress, n)

		go server.onDataReceive(p[:n], remoteAddress)
	}
}
