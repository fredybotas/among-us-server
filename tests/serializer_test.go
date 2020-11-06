package tests

import (
	"cont"
	"parser"
	"testing"
	"version"
)

func TestSerializeRoomsToPacketVersionTwo(t *testing.T) {
	var rooms []cont.Room
	rooms = append(rooms, *cont.NewRoom("AAAAAA", "CN", 0, 0))
	rooms = append(rooms, *cont.NewRoom("BBBBBB", "EU", 0, 0))
	packet := parser.SerializeRoomsToPacket(rooms, version.Two)
	if string(packet) != string(version.Two)+":AAAAAA:CN:BBBBBB:EU:" {
		t.Errorf("wrong packet created")
		return
	}
}

func TestSerializeRoomsToPacketVersionOne(t *testing.T) {
	var rooms []cont.Room
	rooms = append(rooms, *cont.NewRoom("AAAAAA", "CN", 0, 0))
	rooms = append(rooms, *cont.NewRoom("BBBBBB", "EU", 0, 0))
	packet := parser.SerializeRoomsToPacket(rooms, version.One)
	if string(packet) != string(version.One)+":AAAAAA:BBBBBB:" {
		t.Errorf("wrong packet created")
		return
	}
}

func TestSerializeEmptyRoomsToPacket(t *testing.T) {
	var rooms []cont.Room
	packet := parser.SerializeRoomsToPacket(rooms, version.One)
	if string(packet) != "AUS::" {
		t.Errorf("wrong packet created")
		return
	}
}
