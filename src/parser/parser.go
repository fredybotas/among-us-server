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
	AUS:ROOM:123345:42.70898:42.322442: // Protocol, Command, Code, Lat, Lon
			:  6b  :   8b   :    8b   :
			:		PAYLOAD			  :
	AUS:REFR:20000000:42.70898:42.322442:		// Protocol, Command, Proximity
            :   8b   :   8b   :   8b    :
SERVER:
	AUS:123435:323133:243231:432443:	// Protocol, Room list
*/

// Parser for AddRoom command
func ParseRoomPayload(payload []byte) (*cont.Room, error) {
	if len(payload) != 24 {
		return nil, errors.New("wrong payload received")
	}
	if string(payload[6]) != ":" || string(payload[15]) != ":" {
		return nil, errors.New("wrong payload received: delimeters not correct")
	}

	return cont.NewRoom(
		string(payload[0:6]),
		math.Float64frombits(binary.BigEndian.Uint64(payload[7:15])),
		math.Float64frombits(binary.BigEndian.Uint64(payload[16:24])),
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

	return math.Float64frombits(binary.BigEndian.Uint64(payload[0:8])),
		&location,
		nil
}
