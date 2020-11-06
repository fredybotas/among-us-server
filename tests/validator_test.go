package tests

import (
	"parser"
	"testing"
	"version"
)

func TestPacketEmptyPayload(t *testing.T) {
	payload, _, ver, _ := parser.ValidatePacket([]byte("AUS:ROOM::"))
	if len(payload) != 0 {
		t.Errorf("Returned payload should be empty: %v", payload)
	}
	if ver != version.One {
		t.Errorf("version should be One")
	}
}

func TestSmallPacket(t *testing.T) {
	_, _, _, err := parser.ValidatePacket([]byte("AUS:ROOM:"))
	if err == nil {
		t.Errorf("Small packets should end with error")
	}
}

func TestRandomBytes(t *testing.T) {
	_, _, _, err := parser.ValidatePacket([]byte{0, 150, 255, 132, 200})
	if err == nil {
		t.Errorf("Random bytes not recognized")
	}
}

func TestUnknownPacket(t *testing.T) {
	_, _, _, err := parser.ValidatePacket([]byte("ASDASDASDSADASDAS"))
	if err == nil {
		t.Errorf("Unknown packet should return error")
	}
}

func TestGetPayload(t *testing.T) {
	payload, _, _, _ := parser.ValidatePacket([]byte("AUS:ROOM:AAAA:"))
	if string(payload) != "AAAA" {
		t.Errorf("Wrong payload parsed")
	}
}

func TestCommandGetRoom(t *testing.T) {
	_, command, _, _ := parser.ValidatePacket([]byte("AUS:REFR::"))
	if command != parser.GetRooms {
		t.Errorf("Wrong command parsed")
	}
}

func TestCommandAddRoom(t *testing.T) {
	_, command, _, _ := parser.ValidatePacket([]byte("AUS:ROOM::"))
	if command != parser.AddRoom {
		t.Errorf("Wrong command parsed")
	}
}

func TestCommandUnknown(t *testing.T) {
	_, command, _, _ := parser.ValidatePacket([]byte("AUS:SSSA::"))
	if command != parser.Unknown {
		t.Errorf("Wrong command parsed")
	}
}

func TestVersionOneValidation(t *testing.T) {
	_, _, ver, _ := parser.ValidatePacket([]byte(string(version.One) + ":ROOM::"))
	if ver != version.One {
		t.Errorf("wrong version parsed")
	}
}

func TestVersionTwoValidation(t *testing.T) {
	_, _, ver, _ := parser.ValidatePacket([]byte(string(version.Two) + ":ROOM::"))
	if ver != version.Two {
		t.Errorf("wrong version parsed")
	}
}
