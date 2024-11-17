package timep

import (
	"maps"
	"time"
)

type Duration = time.Duration

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour

	Day   = 24 * Hour
	Week  = 7 * Day
	Month = 30 * Day
	Year  = 365 * Day
)

var unitMap = map[string]uint64{
	"ns": uint64(Nanosecond),
	"us": uint64(Microsecond),
	"µs": uint64(Microsecond), // U+00B5 = micro symbol
	"μs": uint64(Microsecond), // U+03BC = Greek letter mu
	"ms": uint64(Millisecond),
	"s":  uint64(Second),
	"m":  uint64(Minute),
	"h":  uint64(Hour),
}

func BaseUnitMap() map[string]uint64 {
	return maps.Clone(unitMap)
}
