package timep

import "time"

var (
	stdParser, _ = DurationParser(unitMap)
	commParser   = commonParser()
)

func commonParser() *durationParser {
	const day = uint64(24 * time.Hour)

	um := BaseUnitMap()
	um["d"] = day
	um["w"] = 7 * day
	um["M"] = 30 * day
	um["y"] = 365 * day

	cp, _ := DurationParser(um)
	return cp
}
