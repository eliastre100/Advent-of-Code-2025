package _9

import "math"

type Tile struct {
	X int
	Y int
}

type Square struct {
	Tiles []Tile
	Size  int
}

func BuildSquares(tiles []Tile) map[int][]Square {
	squares := map[int][]Square{}

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			square := newSquare(tiles[i], tiles[j])
			squares[square.Size] = append(squares[square.Size], square)
		}
	}

	return squares
}

func newSquare(a Tile, b Tile) Square {
	size := int((math.Abs(float64(a.X-b.X)) + 1) * (math.Abs(float64(a.Y-b.Y)) + 1))
	return Square{
		Tiles: []Tile{a, b},
		Size:  size,
	}
}
