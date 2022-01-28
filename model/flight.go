package model

import (
	"sort"
	"time"
)

type Flight struct {
	Number      string
	Origin      string
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

type Flights []Flight

func (fs Flights) OrderBy(field string, order string) {
	rev := true
	if order == "DESC" || order == "desc" {
		rev = false
	}

	switch field {
	case "number":
		sort.Slice(fs, func(i, j int) bool { return (fs[i].Number < fs[j].Number) && rev })
	case "origin":
		sort.Slice(fs, func(i, j int) bool { return (fs[i].Number < fs[j].Number) && rev })
	case "destination":
		sort.Slice(fs, func(i, j int) bool { return (fs[i].Destination < fs[j].Destination) && rev })
	case "departure":
		sort.Slice(fs, func(i, j int) bool { return !rev || (fs[i].Departure.Unix() < fs[j].Departure.Unix()) })
	case "arrival":
		sort.Slice(fs, func(i, j int) bool { return !rev || (fs[i].Arrival.Unix() < fs[j].Arrival.Unix()) })
	}
}
