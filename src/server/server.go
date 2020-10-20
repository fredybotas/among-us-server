package server

import (
	"cont"
	"fmt"
	"net"
)

/***
PROTOCOL
CLIENT:
	AUS:ROOM:123345:42.70898:42.322442: // Protocol, Command, Code, Lat, Lon
	AUS:REFRESH:20:						// Protocol, Command, Proximity

SERVER:
	AUS:123435:323133:243231:432443:	// Protocol, Room list
*/

const port int = 1221
const ip string = "127.0.0.1"

type Server struct {
	socket    *net.UDPConn
	container *cont.Container
}

func (server *Server) Init() {
	server.container = cont.NewContainer()

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	fmt.Printf("Starting server on endpoint: %v\n", addr)
	sock, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error binding sock %v\n", err)
		return
	}

	server.socket = sock
}

func (server *Server) onDataReceive(dataReceived []byte, receivedFrom *net.UDPAddr) {
	_, err := server.socket.WriteToUDP([]byte("Test response"), receivedFrom)
	if err != nil {
		fmt.Printf("Error while responding: %v\n", err)
	}
}

func (server *Server) Serve() {
	fmt.Println("Server waiting for data...")
	var p [1024]byte
	for {
		n, remoteAddress, err := server.socket.ReadFromUDP(p[:])
		if err != nil {
			fmt.Printf("Error receiving data: %v", err)
			continue
		}
		fmt.Printf("Received data from %v with length %d\n", remoteAddress, n)
		currDataCopy := p
		go server.onDataReceive(currDataCopy[:], remoteAddress)
	}
}
