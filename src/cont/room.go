package cont

type Room struct {
	code           string
	serverLocation string
	location       Location
	isActive       bool
}

func NewRoom(code string, serverLocation string, lat float64, lon float64) *Room {
	return &Room{
		code:           code,
		serverLocation: serverLocation,
		location:       NewLocation(lat, lon),
		isActive:       false,
	}
}

func (entry Room) GetCode() string {
	return entry.code
}

func (entry Room) GetServerLocation() string {
	return entry.serverLocation
}

func (entry Room) GetLocation() Location {
	return entry.location
}
