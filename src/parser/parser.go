package parser

import (
	"cont"
	"encoding/binary"
	"errors"
	"math"
)

/***
PROTOCOL
CLIENT:
	AUS:ROOM:123345:CN:42.70898:42.322442: // Protocol, Command, Code, server location, Lat, Lon
			:  6b  :2b:   8b   :    8b   :
			:		  PAYLOAD			 :
	AUS:REFR:20000000:42.70898:42.322442:		// Protocol, Command, Proximity
            :   8b   :   8b   :   8b    :
SERVER:
	AUS:123435:CN:323133:EU:243231:NA:432443:CN:	// Protocol, Room list
*/

// Parser for AddRoom command
func ParseRoomPayload(payload []byte) (*cont.Room, error) {
	if len(payload) != 27 {
		return nil, errors.New("wrong payload received")
	}
	if string(payload[6]) != ":" || string(payload[9]) != ":" || string(payload[18]) != ":" {
		return nil, errors.New("wrong payload received: delimeters not correct")
	}

	return cont.NewRoom(
		string(payload[0:6]),
		string(payload[7:9]),
		math.Float64frombits(binary.BigEndian.Uint64(payload[10:18])),
		math.Float64frombits(binary.BigEndian.Uint64(payload[19:27])),
	), nil
}

// Parser for GetRooms command
func ParseRequestPayload(payload []byte) (float64, *cont.Location, error) {
	if len(payload) != 26 {
		return 0, nil, errors.New("wrong payload received")
	}
	if string(payload[8]) != ":" || string(payload[17]) != ":" {
		return 0, nil, errors.New("wrong payload received: delimeters not correct")
	}

	location := cont.NewLocation(
		math.Float64frombits(binary.BigEndian.Uint64(payload[9:17])),
		math.Float64frombits(binary.BigEndian.Uint64(payload[18:26])),
	)
	proximity := math.Float64frombits(binary.BigEndian.Uint64(payload[0:8]))
	if proximity < 0 {
		proximity = 0
	}
	return proximity, &location, nil
}
