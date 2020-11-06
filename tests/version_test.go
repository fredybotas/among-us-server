package tests

import (
	"testing"
	"version"
)

func TestGetVersion(t *testing.T) {
	ver := version.GetVersion("AUS")
	if ver != version.One {
		t.Errorf("Invalid version returned")
	}

	ver = version.GetVersion("AUR")
	if ver != version.Two {
		t.Errorf("Invalid version returned")
	}

	ver = version.GetVersion("AAAA")
	if ver != version.Unknown {
		t.Errorf("Invalid version returned")
	}
}
