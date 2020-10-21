package tests

import (
	"parser"
	"testing"
)

func TestPacketEmptyPayload(t *testing.T) {
	payload, _, _ := parser.ValidatePacket([]byte("AUS:ROOM::"))
	if len(payload) != 0 {
		t.Errorf("Returned payload should be empty: %v", payload)
	}
}

func TestSmallPacket(t *testing.T) {
	_, _, err := parser.ValidatePacket([]byte("AUS:ROOM:"))
	if err == nil {
		t.Errorf("Small packets should end with error")
	}
}

func TestUnknownPacket(t *testing.T) {
	_, _, err := parser.ValidatePacket([]byte("ASDASDASDSADASDAS"))
	if err == nil {
		t.Errorf("Unknown packet should return error")
	}
}

func TestGetPayload(t *testing.T) {
	payload, _, _ := parser.ValidatePacket([]byte("AUS:ROOM:AAAA:"))
	if string(payload) != "AAAA" {
		t.Errorf("Wrong payload parsed")
	}
}

func TestCommandGetRoom(t *testing.T) {
	_, command, _ := parser.ValidatePacket([]byte("AUS:REFR::"))
	if command != parser.GetRooms {
		t.Errorf("Wrong command parsed")
	}
}

func TestCommandAddRoom(t *testing.T) {
	_, command, _ := parser.ValidatePacket([]byte("AUS:ROOM::"))
	if command != parser.AddRoom {
		t.Errorf("Wrong command parsed")
	}
}

func TestCommandUnknown(t *testing.T) {
	_, command, _ := parser.ValidatePacket([]byte("AUS:SSSA::"))
	if command != parser.Unknown {
		t.Errorf("Wrong command parsed")
	}
}
