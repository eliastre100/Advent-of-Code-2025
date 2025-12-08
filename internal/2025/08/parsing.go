package _8

import (
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]*JunctionBox, error) {
	var boxes []*JunctionBox

	if err := parsing.ReadInputLines(path, func(line string) error {
		parts := strings.Split(line, ",")

		box, err := initializeJunctionBox(parts[0], parts[1], parts[2])
		if err != nil {
			return err
		}

		boxes = append(boxes, box)

		return nil
	}); err != nil {
		return nil, err
	}

	return boxes, nil
}

func initializeJunctionBox(x string, y string, z string) (*JunctionBox, error) {
	var box JunctionBox

	if x, err := strconv.Atoi(x); err != nil {
		return nil, err
	} else {
		box.X = x
	}

	if y, err := strconv.Atoi(y); err != nil {
		return nil, err
	} else {
		box.Y = y
	}

	if z, err := strconv.Atoi(z); err != nil {
		return nil, err
	} else {
		box.Z = z
	}
	box.Circuit = &Circuit{JunctionBoxes: []*JunctionBox{&box}}

	return &box, nil
}
