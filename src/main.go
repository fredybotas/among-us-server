package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)

	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	fmt.Printf("Starting server\n")
	p := make([]byte, 1024)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	for {
		_, remoteaddr, err := server.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s ,,,\n", remoteaddr, p)
		fmt.Printf("Length %v\n", len(p))
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(server, remoteaddr)
	}
}
