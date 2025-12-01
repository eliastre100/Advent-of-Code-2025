package _1

type Dial struct {
	Position int
}

type Direction int

const (
	Left Direction = iota
	Right
)

// Turn the dial in the specified direction by the specified number of steps.
// Only handle steps lower than 100
func (d *Dial) Turn(direction Direction, steps int) {
	if direction == Left {
		steps = -steps
	}
	d.Position = (d.Position + steps + 100) % 100
}
