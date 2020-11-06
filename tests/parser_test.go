package tests

import (
	"encoding/binary"
	"math"
	"parser"
	"testing"
	"version"
)

func prepareAddRoomPacketVersionTwo(code string, serverLocation string, lat, lon float64) []byte {
	result := make([]byte, 0)
	result = append(result, []byte(string(version.Two)+":ROOM:")...)
	result = append(result, []byte(code)...)
	result = append(result, []byte(":")...)
	result = append(result, []byte(serverLocation)...)
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

func prepareAddRoomPacketVersionOne(code string, lat, lon float64) []byte {
	result := make([]byte, 0)
	result = append(result, []byte(string(version.One)+":ROOM:")...)
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

func prepareGetRoomsPacket(proximity, lat, lon float64, ver version.Version) []byte {
	result := make([]byte, 0)
	result = append(result, []byte(string(ver)+":REFR:")...)
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

func TestPayloadParseGetRoomsVersionOne(t *testing.T) {
	lat := 1.2
	lon := 2.5
	proximity := 10.0
	packet := prepareGetRoomsPacket(proximity, lat, lon, version.One)
	payload, _, ver, _ := parser.ValidatePacket(packet)
	prox, location, err := parser.ParseRequestPayload(payload, ver)
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

func TestPayloadParseGetRoomsVersionTwo(t *testing.T) {
	lat := 1.2
	lon := 2.5
	proximity := 10.0
	packet := prepareGetRoomsPacket(proximity, lat, lon, version.Two)
	payload, _, ver, _ := parser.ValidatePacket(packet)
	prox, location, err := parser.ParseRequestPayload(payload, ver)
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

func TestPayloadParseAddRoomVersionTwo(t *testing.T) {
	lat := 1.2
	lon := 2.5
	code := "AAAAAA"
	serverLocation := "CN"
	packet := prepareAddRoomPacketVersionTwo(code, serverLocation, lat, lon)
	payload, _, ver, _ := parser.ValidatePacket(packet)
	room, err := parser.ParseRoomPayload(payload, ver)
	if err != nil {
		t.Errorf("failed while parsing")
		return
	}
	if room.GetCode() != code {
		t.Errorf("failed reading code")
	}
	if room.GetServerLocation() != serverLocation {
		t.Errorf("failed reading server location")
	}
	if room.GetLocation().GetLat() != lat || room.GetLocation().GetLon() != lon {
		t.Errorf("failed reading location")
	}
}

func TestPayloadParseAddRoomVersionOne(t *testing.T) {
	lat := 1.2
	lon := 2.5
	code := "AAAAAA"
	packet := prepareAddRoomPacketVersionOne(code, lat, lon)
	payload, _, ver, _ := parser.ValidatePacket(packet)
	room, err := parser.ParseRoomPayload(payload, ver)
	if err != nil {
		t.Errorf("failed while parsing")
		return
	}
	if room.GetCode() != code {
		t.Errorf("failed reading code")
	}
	if room.GetServerLocation() != parser.ServerLocationDefault {
		t.Errorf("failed reading server location")
	}
	if room.GetLocation().GetLat() != lat || room.GetLocation().GetLon() != lon {
		t.Errorf("failed reading location")
	}
}
