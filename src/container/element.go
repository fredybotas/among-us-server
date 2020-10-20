package container

type location struct {
	lat float64
	lon float64
}

type RoomEntry struct {
	code     string
	location location
}

func CreateEntry(code string, lat float64, lon float64) RoomEntry {
	var entry RoomEntry
	entry.code = code
	entry.location = location{lat, lon}
	return entry
}
