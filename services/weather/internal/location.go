package internal

type Location struct {
	Lat, Long string
}

func NewLocation(lat, long string) *Location {
	return &Location{
		Lat:  lat,
		Long: long,
	}
}
