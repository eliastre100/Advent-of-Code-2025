package _5

import "github.com/charmbracelet/log"

type Ingredient int

type Database []Range

func (d Database) IsFresh(i Ingredient) bool {
	for _, r := range d {
		if r.Start <= int(i) && int(i) <= r.End {
			return true
		}
	}
	return false
}

func (d *Database) Optimise() {
	var ranges []Range
	changed := false

	for _, oldRange := range *d {
		merged := false
		for i, newRange := range ranges {
			if Overlap(oldRange, newRange) {
				log.Debugf("merging %v into %v", oldRange, newRange)
				ranges[i] = Merge(newRange, oldRange)
				changed = true
				merged = true
				break
			}
		}
		if !merged {
			ranges = append(ranges, oldRange)
		}
	}

	*d = ranges
	if changed {
		d.Optimise()
	}
}

func Count(d Database) int {
	log.Debugf("before optimisation: %+v", d)
	d.Optimise()
	log.Debugf("after optimisation: %+v", d)
	c := 0

	for _, r := range d {
		c += r.End - r.Start + 1
	}

	return c
}

func (d *Database) Insert(r Range) {
	*d = append(*d, r)
}
