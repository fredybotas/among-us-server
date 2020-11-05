package tests

import (
	"cont"
	"math/rand"
	"testing"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestContainerAdd(t *testing.T) {
	container := cont.NewContainer()
	inserted := container.InsertEntry(cont.NewRoom("aaa", "CN", 1, 2))
	if inserted == false {
		t.Errorf("Element was not inserted")
	}
}

func TestContainerAddAndFetch(t *testing.T) {
	container := cont.NewContainer()
	container.InsertEntry(cont.NewRoom("aaa", "CN", 1, 2))
	res := container.Query(cont.NewLocation(1, 2), 1)
	if len(res) != 1 {
		t.Errorf("Element was not inserted")
	}

	if res[0].GetCode() != "aaa" {
		t.Errorf("Received wrong element")
	}
}

func TestContainerAddDuplicate(t *testing.T) {
	container := cont.NewContainer()
	container.InsertEntry(cont.NewRoom("aaa", "CN", 1, 2))
	inserted := container.InsertEntry(cont.NewRoom("aaa", "CN", 1, 2))
	if inserted == true {
		t.Errorf("Element inserted twice")
	}
	res := container.Query(cont.NewLocation(1, 2), 1)
	if len(res) != 1 {
		t.Errorf("Element was not inserted")
	}

	if res[0].GetCode() != "aaa" {
		t.Errorf("Received wrong element")
	}
}

func TestContainerAddConcurrently(t *testing.T) {
	container := cont.NewContainer()

	for i := 0; i < 1000; i++ {
		go container.InsertEntry(cont.NewRoom(StringWithCharset(3, "abcdefghijklmnop"), "CN", 1, 2))
	}

	if container.GetCount() == 0 {
		t.Errorf("Elements were not inserted")
	}
}

func TestContainerAddAndReadConcurrently(t *testing.T) {
	container := cont.NewContainer()

	for i := 0; i < 100000; i++ {
		go container.InsertEntry(cont.NewRoom(StringWithCharset(5, "abcdefghijklmnop"), "CN", 1, 2))
		go container.Query(cont.NewLocation(1, 2), 1)
	}

	if container.GetCount() == 0 {
		t.Errorf("Elements were not inserted")
	}
}

func TestRemoveInactiveEntries(t *testing.T) {
	container := cont.NewContainer()
	for i := 0; i < 1000; i++ {
		go container.InsertEntry(cont.NewRoom(StringWithCharset(5, "abcdefghijklmnop"), "CN", 1, 2))
	}
	time.Sleep((cont.CheckInterval*2 + 1) * time.Second)
	if container.GetCount() != 0 {
		t.Errorf("Elements were not deleted")
	}
}
