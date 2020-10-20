package cont

import (
	"fmt"
	"sync"
	"time"
)

const CheckInterval = 5

type Container struct {
	entries   map[string]*RoomEntry
	writeLock sync.RWMutex
}

func NewContainer() *Container {
	var cont Container
	cont.entries = make(map[string]*RoomEntry)
	go cont.periodicClean()
	return &cont
}

func (container *Container) periodicClean() {
	for {
		time.Sleep(CheckInterval * time.Second)
		container.writeLock.Lock()

		var keysToDelete []string
		for key, value := range container.entries {
			if value.isActive {
				value.isActive = false
			} else {
				keysToDelete = append(keysToDelete, key)
			}
		}

		for _, key := range keysToDelete {
			delete(container.entries, key)
			fmt.Printf("Removing room: %s due to inactivity\n", key)
		}

		container.writeLock.Unlock()
	}
}

func (container *Container) InsertEntry(entry *RoomEntry) bool {
	container.writeLock.Lock()
	defer container.writeLock.Unlock()

	element, entryExists := container.entries[entry.code]
	if entryExists {
		element.isActive = true
		return false
	}
	entry.isActive = true
	container.entries[entry.code] = entry

	return true
}

func (container *Container) Query(location Location, radius float64) []RoomEntry {
	container.writeLock.RLock()
	defer container.writeLock.RUnlock()
	result := make([]RoomEntry, 0)
	for _, value := range container.entries {
		if checkProximity(value.GetLocation(), location, radius) {
			result = append(result, *value)
		}
	}
	return result
}

func (container *Container) GetCount() int {
	container.writeLock.RLock()
	defer container.writeLock.RUnlock()
	return len(container.entries)
}

func checkProximity(location1 Location, location2 Location, radius float64) bool {
	// TODO: Implement correctly
	return true
}
