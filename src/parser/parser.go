package parser

import (
	"cont"
	"errors"
)

/***
PROTOCOL
CLIENT:
	AUS:ROOM:123345:42.70898:42.322442: // Protocol, Command, Code, Lat, Lon
	AUS:REFR:20:						// Protocol, Command, Proximity
	AUS:REFR::							// Protocol, Command, Proximity

SERVER:
	AUS:123435:323133:243231:432443:	// Protocol, Room list
*/

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

// Parser for AddRoom command
func ParseRoomPayload(payload []byte) (*cont.Room, error) {
	//TODO: Implement
	return cont.NewRoom("", 0, 0), nil
}

// Parser for GetRooms command
func ParseProximityPayload(payload []byte) (float64, error) {
	//TODO: Implement
	return 0, nil
}
