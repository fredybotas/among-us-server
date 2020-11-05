package parser

import "cont"

func SerializeRoomsToPacket(rooms []cont.Room) []byte {
	result := make([]byte, 0)
	result = append(result, []byte("AUS:")...)
	for _, room := range rooms {
		result = append(result, []byte(room.GetCode())...)
		result = append(result, []byte(":")...)
		result = append(result, []byte(room.GetServerLocation())...)
		result = append(result, []byte(":")...)
	}
	if len(rooms) == 0 {
		result = append(result, []byte(":")...)
	}
	return result
}
