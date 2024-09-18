package timep

import (
	"errors"
	"strings"
)

type durationParser struct {
	unitMap map[string]uint64
}

func DurationParser(unitMap map[string]uint64) (*durationParser, error) {
	//  checking keys [not containing dig]
	for key, _ := range unitMap {
		if strings.ContainsFunc(key, func(r rune) bool { return '0' <= r && r <= '9' }) {
			return nil, errors.New("invalid key [" + key + "] contains number")
		}
	}

	return &durationParser{unitMap: unitMap}, nil
}

func (p *durationParser) ParseDuration(s string) (Duration, error) {
	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	var d uint64
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("time: invalid duration " + quote(orig))
	}
	for s != "" {
		var (
			v, f  uint64      // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, errors.New("time: invalid duration " + quote(orig))
		}
		// Consume [0-9]*
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("time: invalid duration " + quote(orig))
		}
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("time: invalid duration " + quote(orig))
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return 0, errors.New("time: missing unit in duration " + quote(orig))
		}
		u := s[:i]
		s = s[i:]
		unit, ok := p.unitMap[u]
		if !ok {
			return 0, errors.New("time: unknown unit " + quote(u) + " in duration " + quote(orig))
		}
		if v > 1<<63/unit {
			// overflow
			return 0, errors.New("time: invalid duration " + quote(orig))
		}
		v *= unit
		if f > 0 {
			// float64 is needed to be nanosecond accurate for fractions of hours.
			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
			v += uint64(float64(f) * (float64(unit) / scale))
			if v > 1<<63 {
				// overflow
				return 0, errors.New("time: invalid duration " + quote(orig))
			}
		}
		d += v
		if d > 1<<63 {
			return 0, errors.New("time: invalid duration " + quote(orig))
		}
	}
	if neg {
		return -Duration(d), nil
	}
	if d > 1<<63-1 {
		return 0, errors.New("time: invalid duration " + quote(orig))
	}
	return Duration(d), nil
}
