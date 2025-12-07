package _7

import (
	"slices"
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

func (d *Diagram) SplitCount() (int, int) {
	beams := map[int]int{d.Start.X: 1}
	splitCount := 0

	for y := d.Start.Y; y < d.Height; y++ {
		nextBeams := map[int]int{}

		for beam, count := range beams {
			if slices.Contains(d.Splitters[y], beam) {
				splitCount += 1
				nextBeams[beam+1] = nextBeams[beam+1] + count
				nextBeams[beam-1] = nextBeams[beam-1] + count
			} else {
				nextBeams[beam] = nextBeams[beam] + count
			}
		}

		beams = nextBeams
	}

	timelines := 0
	for _, count := range beams {
		timelines += count
	}
	return splitCount, timelines
}
