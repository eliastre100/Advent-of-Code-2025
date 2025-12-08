package _8

import (
	"maps"
	"math"
	"slices"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/utils"
)

type JunctionBox struct {
	X       int
	Y       int
	Z       int
	Circuit *Circuit
}

type Link struct {
	A        *JunctionBox
	B        *JunctionBox
	Distance int
}

func Distance(a JunctionBox, b JunctionBox) int {
	return int(math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2)))
}

func GenerateLinks(boxes []*JunctionBox) []Link {
	links := map[int][]Link{}

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			if i == j {
				continue
			}
			distance := Distance(*boxes[i], *boxes[j])

			links[distance] = append(links[distance], Link{boxes[i], boxes[j], distance})
		}
	}

	return utils.Flatten(slices.Collect(maps.Values(links)))
}

func LinkJunctionBoxes(boxes []*JunctionBox, threshold int) {
	links := GenerateLinks(boxes)
	slices.SortFunc(links, func(a, b Link) int {
		return a.Distance - b.Distance
	})

	processed := 0
	for _, link := range links {
		if link.A.Circuit != link.B.Circuit {
			linkBoxes(link.A.Circuit, link.B.Circuit)
		}
		processed++
		if processed >= threshold {
			break
		}
	}
}
