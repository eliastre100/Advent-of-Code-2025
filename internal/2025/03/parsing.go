package _3

import (
	"strconv"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]Battery, error) {
	var batteries []Battery

	if err := parsing.ReadInputLines(path, func(line string) error {
		joltages := make([]int, len(line))
		for i, c := range line {
			joltage, err := strconv.Atoi(string(c))
			if err != nil {
				return err
			}
			joltages[i] = joltage
		}
		batteries = append(batteries, Battery{Joltages: joltages})
		return nil
	}); err != nil {
		return nil, err
	}

	return batteries, nil
}
