package cont

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

const CheckInterval = 5
const maxServerCount = 5

type Container struct {
	entries   map[string]*Room
	writeLock sync.RWMutex
}

func NewContainer() *Container {
	var cont Container
	cont.entries = make(map[string]*Room)
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

func (container *Container) InsertEntry(entry *Room) bool {
	container.writeLock.Lock()
	defer container.writeLock.Unlock()

	element, entryExists := container.entries[entry.code]
	if entryExists {
		element.isActive = true
		element.location = entry.location
		return false
	}
	entry.isActive = true
	container.entries[entry.code] = entry

	return true
}

func (container *Container) Query(location Location, radius float64) []Room {
	container.writeLock.RLock()
	result := make([]Room, 0)
	for _, value := range container.entries {
		result = append(result, *value)
	}
	container.writeLock.RUnlock()
	sort.Slice(result, func(i, j int) bool {
		return location.GetDistance(result[i].GetLocation()) < location.GetDistance(result[j].GetLocation())
	})
	resultCount := maxServerCount
	if len(result) < maxServerCount {
		resultCount = len(result)
	}
	return result[:resultCount]
}

func (container *Container) GetCount() int {
	container.writeLock.RLock()
	defer container.writeLock.RUnlock()
	return len(container.entries)
}
