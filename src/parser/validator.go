package parser

import "errors"

type Command byte

const (
	Unknown  Command = iota // c0 == 0
	AddRoom  Command = iota // c1 == 1
	GetRooms Command = iota // c2 == 2
)

func ValidatePacket(data []byte) ([]byte, Command, error) {
	if len(data) < 4 {
		return nil, Unknown, errors.New("unknown packet received")
	}
	if string(data[:4]) != "AUS:" {
		return nil, Unknown, errors.New("received packet has not our protocol")
	}
	if len(data) < 10 || string(data[len(data)-1]) != ":" || string(data[8]) != ":" {
		return nil, Unknown, errors.New("corrupted packet received")
	}
	command := string(data[4:8])
	if command == "ROOM" {
		return data[9 : len(data)-1], AddRoom, nil
	} else if command == "REFR" {
		return data[9 : len(data)-1], GetRooms, nil
	}
	return nil, Unknown, errors.New("unknown command received")
}
