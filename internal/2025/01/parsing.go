package _1

import (
	"strconv"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

type Instruction struct {
	Direction Direction
	Steps     int
}

func Parse(path string) ([]Instruction, error) {
	instructions := make([]Instruction, 0)
	if err := parsing.ReadInputLines(path, func(line string) error {
		if len(line) != 0 {
			instruction := Instruction{Direction: Left, Steps: 0}
			if line[0] == 'R' {
				instruction.Direction = Right
			}
			instruction.Steps, _ = strconv.Atoi(line[1:len(line)])
			instructions = append(instructions, instruction)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return instructions, nil
}
