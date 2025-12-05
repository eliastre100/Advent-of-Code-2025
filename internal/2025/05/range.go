package _5

import "slices"

type Range struct {
	Start int
	End   int
}

func Overlap(r1, r2 Range) bool {
	r2StartIncluded := r1.Start <= r2.Start && r2.Start <= r1.End
	r1StartIncluded := r2.Start <= r1.Start && r1.Start <= r2.End
	r1EndIncluded := r1.Start <= r2.End && r2.End <= r1.End
	r2EndIncluded := r2.Start <= r1.End && r1.End <= r2.End

	return r2StartIncluded || r1StartIncluded || r1EndIncluded || r2EndIncluded
}

func Merge(r Range, other Range) Range {
	lower := slices.Min([]int{r.Start, other.Start})
	upper := slices.Max([]int{r.End, other.End})
	return Range{lower, upper}
}
