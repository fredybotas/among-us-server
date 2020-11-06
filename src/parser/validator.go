package parser

import "errors"
import "version"

type Command byte

const (
	Unknown  Command = iota // c0 == 0
	AddRoom  Command = iota // c1 == 1
	GetRooms Command = iota // c2 == 2
)

func ValidatePacket(data []byte) ([]byte, Command, version.Version, error) {
	if len(data) < 4 {
		return nil, Unknown, version.Unknown, errors.New("unknown packet received")
	}
	ver := version.GetVersion(string(data[:3]))
	if ver == version.Unknown {
		return nil, Unknown, ver, errors.New("received packet has not our protocol")
	}
	if len(data) < 10 || string(data[len(data)-1]) != ":" || string(data[8]) != ":" || string(data[3]) != ":" {
		return nil, Unknown, version.Unknown, errors.New("corrupted packet received")
	}
	command := string(data[4:8])
	if command == "ROOM" {
		return data[9 : len(data)-1], AddRoom, ver, nil
	} else if command == "REFR" {
		return data[9 : len(data)-1], GetRooms, ver, nil
	}
	return nil, Unknown, ver, errors.New("unknown command received")
}
