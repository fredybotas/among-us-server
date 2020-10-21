package tests

import (
	"cont"
	"parser"
	"testing"
)

func TestSerializeRoomsToPacket(t *testing.T) {
	var rooms []cont.Room
	rooms = append(rooms, *cont.NewRoom("AAAAAA", 0, 0))
	rooms = append(rooms, *cont.NewRoom("BBBBBB", 0, 0))
	packet := parser.SerializeRoomsToPacket(rooms)
	if string(packet) != "AUS:AAAAAA:BBBBBB:" {
		t.Errorf("wrong packet created")
		return
	}
}
