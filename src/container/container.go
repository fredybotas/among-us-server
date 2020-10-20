package container

import "sync"

type Container struct {
	entries       []RoomEntry
	entriesExists map[string]bool
	writeLock     sync.RWMutex
}

func (container *Container) Init() {
	//TODO: Test if works also without make
	//container.entries = make([]RoomEntry, 0)
	container.entriesExists = make(map[string]bool)
}

func (container *Container) InsertEntry(entry RoomEntry) bool {
	container.writeLock.RLock()
	_, entryExists := container.entriesExists[entry.code]
	container.writeLock.RUnlock()

	if entryExists {
		return false
	} else {
		container.writeLock.Lock()
		defer container.writeLock.Unlock()
		container.entriesExists[entry.code] = true
	}

	container.entries = append(container.entries, entry)
	return true
}

func (container *Container) Query(location location, radius float64) []RoomEntry {
	container.writeLock.RLock()
	defer container.writeLock.RUnlock()
	result := make([]RoomEntry, 0)
	for _, element := range container.entries {
		if checkProximity(element.location, location, radius) {
			result = append(result, element)
		}
	}
	return result
}

func checkProximity(location1 location, location2 location, radius float64) bool {
	// TODO: Implement correctly
	return true
}
