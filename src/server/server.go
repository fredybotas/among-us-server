package server

import (
	"cont"
	"fmt"
	"net"
	"parser"
)

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
	payload, command, ver, err := parser.ValidatePacket(dataReceived)
	if err != nil {
		fmt.Printf("Packet received error: %v\n", err)
		return
	}

	if command == parser.AddRoom {
		room, err := parser.ParseRoomPayload(payload, ver)
		if err != nil {
			fmt.Printf("Error parsing packet: %v\n", err)
			return
		}
		server.container.InsertEntry(room)
	} else if command == parser.GetRooms {
		proximity, loc, err := parser.ParseRequestPayload(payload, ver)
		if err != nil {
			fmt.Printf("Error parsing packet: %v\n", err)
			return
		}
		result := server.container.Query(*loc, proximity)
		_, err1 := server.socket.WriteToUDP(parser.SerializeRoomsToPacket(result, ver), receivedFrom)
		if err1 != nil {
			fmt.Printf("Error while responding: %v\n", err)
		}
	} else {
		fmt.Printf("Unknown command received")
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
		go server.onDataReceive(currDataCopy[:n], remoteAddress)
	}
}
