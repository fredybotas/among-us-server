package tests

import (
	"cont"
	"parser"
	"testing"
)

func TestSerializeRoomsToPacket(t *testing.T) {
	var rooms []cont.Room
	rooms = append(rooms, *cont.NewRoom("AAAAAA", "CN", 0, 0))
	rooms = append(rooms, *cont.NewRoom("BBBBBB", "EU", 0, 0))
	packet := parser.SerializeRoomsToPacket(rooms)
	if string(packet) != "AUS:AAAAAA:CN:BBBBBB:EU:" {
		t.Errorf("wrong packet created")
		return
	}
}

func TestSerializeEmptyRoomsToPacket(t *testing.T) {
	var rooms []cont.Room
	packet := parser.SerializeRoomsToPacket(rooms)
	if string(packet) != "AUS::" {
		t.Errorf("wrong packet created")
		return
	}
}
