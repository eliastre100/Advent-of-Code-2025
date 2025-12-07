package _7

import "github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"

func Parse(path string) (*Diagram, error) {
	diagram := NewDiagram()

	if err := parsing.ReadInputLines(path, func(line string) error {
		y := diagram.Height
		for x, c := range line {
			if c == 'S' {
				diagram.Start.Y = y
				diagram.Start.X = x
			} else if c == '^' {
				diagram.Splitters[y] = append(diagram.Splitters[y], x)
			}
		}

		diagram.Height += 1
		diagram.Width = len(line)
		return nil
	}); err != nil {
		return diagram, err
	}

	return diagram, nil
}
