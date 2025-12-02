package _2

import (
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

type Validator func(i int) bool

func InvalidValuesInRange(r Range, validator Validator) []int {
	var values []int

	for i := r.Start; i <= r.End; i++ {
		if !validator(i) {
			values = append(values, i)
		}
	}
	return values
}

func NotDuplicate(i int) bool {
	str := strconv.Itoa(i)
	build := strings.Repeat(str[:len(str)/2], 2)
	return build != str
}

func NotRepeating(i int) bool {
	str := strconv.Itoa(i)
	starts := patternStarts(strconv.Itoa(i), string(str[0]))

	for patternSize := 1; patternSize <= len(starts)/2; patternSize++ { // The minimum pattern length is at most half the string, such as at least half the starts
		pattern := str[0:starts[patternSize]]
		patternBuild := strings.Repeat(pattern, len(starts)/patternSize)
		if patternBuild == str {
			return false
		}
	}
	return true
}

func patternStarts(str string, pattern string) []int {
	var starts []int

	for i := 0; ; {
		idx := strings.Index(str[i:], pattern)
		if idx == -1 {
			break
		}
		starts = append(starts, i+idx)
		i += idx + 1
	}
	return starts
}

/*func InvalidValuesInRange(r Range) []int {
	var values []int

	i := extractGroupDefinition(r.Start, 0)
	end := extractGroupDefinition(r.End, 1)
	for ; i <= end; i++ {
		expended := expandGroup(i)
		if expended >= r.Start && expended <= r.End {
			values = append(values, expended)
		}
	}
	return values
}

func extractGroupDefinition(v int, oddBonus int) int {
	s := strconv.Itoa(v)
	if len(s)%2 == 0 {
		v, _ = strconv.Atoi(s[0 : len(s)/2]) // There should never be a way to have an invalid number definition here
	} else {
		v, _ = strconv.Atoi(s[0 : len(s)/2+oddBonus]) // There should never be a way to have an invalid number definition here
	}
	return v
}

func expandGroup(v int) int {
	// This method could probably be done more efficiently by using bitshift, will see if part 2 might benefit from it later
	representation := strconv.Itoa(v)
	v, _ = strconv.Atoi(representation + representation)
	return v
}*/
