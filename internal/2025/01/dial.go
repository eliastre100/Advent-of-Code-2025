package _1

import (
	"math"

	"github.com/charmbracelet/log"
)

type Dial struct {
	Position int
}

type Direction int

const (
	Left Direction = iota
	Right
)

// Turn the dial in the specified direction by the specified number of steps and return the number of times it pointed to 0 by doing so.
// Only handle steps lower than 100
func (d *Dial) Turn(direction Direction, steps int) int {
	if direction == Left {
		steps = -steps
	}

	initialPosition := d.Position
	expectedPosition := d.Position + steps
	log.Debug("Turning the dial", "steps", steps, "expected", expectedPosition)

	if expectedPosition > 0 {
		d.Position = expectedPosition % 100
		return expectedPosition / 100
	} else {
		d.Position = (100 + (expectedPosition % 100)) % 100
		return int(math.Abs(float64(expectedPosition/100))) + signFlipBonus(initialPosition, expectedPosition)
	}
}

func signFlipBonus(initialPosition int, expectedPosition int) int {
	if initialPosition != 0 && expectedPosition < initialPosition {
		return 1
	} else {
		return 0
	}
}
