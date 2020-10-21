package cont

type Location struct {
	lat float64
	lon float64
}

func NewLocation(lat, lon float64) Location {
	return Location{lat, lon}
}

func (location Location) CheckProximity(location1 Location, radius float64) bool {
	// TODO: Implement correctly
	return true
}

func (location Location) GetLon() float64 {
	return location.lon
}

func (location Location) GetLat() float64 {
	return location.lat
}
