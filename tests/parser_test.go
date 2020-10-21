package tests

import (
	"encoding/binary"
	"math"
	"parser"
	"testing"
)

func prepareAddRoomPacket(code string, lat, lon float64) []byte {
	result := make([]byte, 0)
	result = append(result, []byte("AUS:ROOM:")...)
	result = append(result, []byte(code)...)
	result = append(result, []byte(":")...)
	var buf1 [8]byte
	binary.BigEndian.PutUint64(buf1[:], math.Float64bits(lat))
	result = append(result, buf1[:]...)
	result = append(result, []byte(":")...)
	var buf2 [8]byte
	binary.BigEndian.PutUint64(buf2[:], math.Float64bits(lon))
	result = append(result, buf2[:]...)
	result = append(result, []byte(":")...)
	return result
}

func prepareGetRoomsPacket(proximity, lat, lon float64) []byte {
	result := make([]byte, 0)
	result = append(result, []byte("AUS:REFR:")...)
	var buf1 [8]byte
	binary.BigEndian.PutUint64(buf1[:], math.Float64bits(proximity))
	result = append(result, buf1[:]...)
	result = append(result, []byte(":")...)

	var buf2 [8]byte
	binary.BigEndian.PutUint64(buf2[:], math.Float64bits(lat))
	result = append(result, buf2[:]...)
	result = append(result, []byte(":")...)

	var buf3 [8]byte
	binary.BigEndian.PutUint64(buf3[:], math.Float64bits(lon))
	result = append(result, buf3[:]...)
	result = append(result, []byte(":")...)
	return result
}

func TestPayloadParseGetRooms(t *testing.T) {
	lat := 1.2
	lon := 2.5
	proximity := 10.0
	packet := prepareGetRoomsPacket(proximity, lat, lon)
	payload, _, _ := parser.ValidatePacket(packet)
	prox, location, err := parser.ParseRequestPayload(payload)
	if err != nil {
		t.Errorf("failed while parsing")
		return
	}
	if prox != proximity {
		t.Errorf("failed reading proximity")
	}
	if location.GetLat() != lat || location.GetLon() != lon {
		t.Errorf("failed reading location")
	}
}

func TestPayloadParseAddRoom(t *testing.T) {
	lat := 1.2
	lon := 2.5
	code := "AAAAAA"
	packet := prepareAddRoomPacket(code, lat, lon)
	payload, _, _ := parser.ValidatePacket(packet)
	room, err := parser.ParseRoomPayload(payload)
	if err != nil {
		t.Errorf("failed while parsing")
		return
	}
	if room.GetCode() != code {
		t.Errorf("failed reading code")
	}
	if room.GetLocation().GetLat() != lat || room.GetLocation().GetLon() != lon {
		t.Errorf("failed reading location")
	}
}
