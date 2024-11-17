package timep

var (
	stdParser, _ = DurationParser(unitMap)
	commParser   = commonParser()
)

func commonParser() *durationParser {
	um := BaseUnitMap()

	um["D"] = uint64(Day)
	um["d"] = uint64(Day) // for flexibility

	um["W"] = uint64(Week)
	um["w"] = uint64(Week) // for flexibility

	um["M"] = uint64(Month)

	um["Y"] = uint64(Year)
	um["y"] = uint64(Year) // for flexibility

	cp, _ := DurationParser(um)

	return cp
}
