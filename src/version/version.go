package version

type Version string

const (
	One     Version = "AUS"
	Two     Version = "AUR"
	Unknown Version = "UNK"
)

func GetVersion(protocol string) Version {
	switch protocol {
	case string(One):
		return One
	case string(Two):
		return Two
	default:
		return Unknown
	}
}
