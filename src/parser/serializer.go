package parser

import (
	"cont"
	"version"
)

func SerializeRoomsToPacket(rooms []cont.Room, ver version.Version) []byte {
	result := make([]byte, 0)
	result = append(result, []byte(ver)...)
	result = append(result, []byte(":")...)
	for _, room := range rooms {
		result = append(result, []byte(room.GetCode())...)
		result = append(result, []byte(":")...)
		if ver == version.Two {
			result = append(result, []byte(room.GetServerLocation())...)
			result = append(result, []byte(":")...)
		}
	}
	if len(rooms) == 0 {
		result = append(result, []byte(":")...)
	}
	return result
}
