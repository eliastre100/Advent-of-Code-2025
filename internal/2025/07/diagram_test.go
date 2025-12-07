package _7

import "testing"

func TestSimulate(t *testing.T) {
	tests := []struct {
		Diagram  Diagram
		Expected int
	}{
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{}}, 0},
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{
			1: {1},
		}}, 1},
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{
			1: {1},
			2: {0, 1, 2},
		}}, 3},
	}

	for i, test := range tests {
		result := test.Diagram.Simulate()
		if result != test.Expected {
			t.Errorf("[%d] Expected %d, got %d", i+1, test.Expected, result)
		}
	}
}
