package _1

import "testing"

func TestDialRotation(t *testing.T) {
	dial := Dial{Position: 0}
	testRotations := []struct {
		InitialState int
		Direction    Direction
		Steps        int
		Expected     int
	}{
		{InitialState: 50, Direction: Right, Steps: 2, Expected: 52}, // Right movement
		{InitialState: 50, Direction: Left, Steps: 2, Expected: 48},  // Left movement
		{InitialState: 50, Direction: Right, Steps: 51, Expected: 1}, // Right underflow
		{InitialState: 50, Direction: Left, Steps: 51, Expected: 99}, // Left overflow
	}

	for _, test := range testRotations {
		dial.Position = test.InitialState
		dial.Turn(test.Direction, test.Steps)
		if dial.Position != test.Expected {
			t.Errorf("Expected position %d, got %d", test.Expected, dial.Position)
		}
	}
}
