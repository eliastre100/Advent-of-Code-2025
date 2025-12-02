package _2

import (
	"strconv"
)

type Range struct {
	Start int
	End   int
}

func InvalidValuesInRange(r Range) []int {
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
}
