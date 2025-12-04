package _4

import "github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"

func Parse(path string) (Map, error) {
	var m Map
	m.RollPositions = make(map[int]map[int]bool)

	if err := parsing.ReadInputLines(path, func(line string) error {
		m.RollPositions[m.Height] = make(map[int]bool)

		for i, c := range line {
			m.RollPositions[m.Height][i] = c == '@'
		}

		m.Width = len(line)
		m.Height += 1

		return nil
	}); err != nil {
		return m, err
	}

	return m, nil
}
