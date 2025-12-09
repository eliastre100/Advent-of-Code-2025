package _9

import (
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]Tile, error) {
	var tiles []Tile

	if err := parsing.ReadInputLines(path, func(line string) error {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0]) // In a real-life scenario, the conversion should be checked, but AoC inputs are always valid
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, Tile{x, y})
		return nil
	}); err != nil {
		return nil, err
	}

	return tiles, nil
}
