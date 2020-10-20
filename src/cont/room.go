package cont

type Room struct {
	code     string
	location Location
	isActive bool
}

func NewRoom(code string, lat float64, lon float64) *Room {
	return &Room{
		code:     code,
		location: NewLocation(lat, lon),
		isActive: false,
	}
}

func (entry Room) GetCode() string {
	return entry.code
}

func (entry Room) GetLocation() Location {
	return entry.location
}
