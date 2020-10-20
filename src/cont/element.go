package cont

type RoomEntry struct {
	code     string
	location Location
	isActive bool
}

func NewEntry(code string, lat float64, lon float64) RoomEntry {
	var entry RoomEntry
	entry.code = code
	entry.location = NewLocation(lat, lon)
	return entry
}

func (entry RoomEntry) GetCode() string {
	return entry.code
}

func (entry RoomEntry) GetLocation() Location {
	return entry.location
}
