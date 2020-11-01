package cont

import "math"

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

func (location Location) GetDistance(location1 Location) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * location.lat / 180.0)
	radlat2 := float64(PI * location1.lat / 180.0)

	theta := float64(location.lon - location1.lon)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	// KM
	dist = dist * 1.609344

	return dist
}

func (location Location) GetLon() float64 {
	return location.lon
}

func (location Location) GetLat() float64 {
	return location.lat
}
