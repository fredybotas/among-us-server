package cont

type Location struct {
	lat float64
	lon float64
}

func NewLocation(lat, lon float64) Location {
	return Location{lat, lon}
}
