package _7

import (
	"slices"
	"sort"
)

type Position struct {
	X int
	Y int
}

type Diagram struct {
	Width     int
	Height    int
	Start     Position
	Splitters map[int][]int
}

func NewDiagram() *Diagram {
	return &Diagram{Splitters: map[int][]int{}}
}

func (d *Diagram) Simulate() int {
	beams := []int{d.Start.X}
	splitCount := 0

	for y := d.Start.Y; y < d.Height; y++ {
		var nextBeams []int

		for _, beam := range beams {
			if slices.Contains(d.Splitters[y], beam) {
				splitCount += 1
				nextBeams = append(nextBeams, beam+1, beam-1)
			} else {
				nextBeams = append(nextBeams, beam)
			}
		}

		sort.Ints(nextBeams)
		beams = slices.Compact(nextBeams)
	}

	return splitCount
}
