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
		{InitialState: 50, Direction: Left, Steps: 68, Expected: 82},
	}

	for _, test := range testRotations {
		dial.Position = test.InitialState
		dial.Turn(test.Direction, test.Steps)
		if dial.Position != test.Expected {
			t.Errorf("Expected position %d, got %d", test.Expected, dial.Position)
		}
	}
}

func TestDialRotationZeroCount(t *testing.T) {
	dial := Dial{Position: 0}
	testRotations := []struct {
		InitialState int
		Direction    Direction
		Steps        int
		Expected     int
	}{
		{InitialState: 50, Direction: Right, Steps: 2, Expected: 0},   // Right movement without crossing
		{InitialState: 50, Direction: Left, Steps: 2, Expected: 0},    // Left movement without crossing
		{InitialState: 50, Direction: Right, Steps: 50, Expected: 1},  // Right movement to 0
		{InitialState: 50, Direction: Left, Steps: 50, Expected: 1},   // Left movement to 0
		{InitialState: 50, Direction: Right, Steps: 51, Expected: 1},  // Right underflow once
		{InitialState: 50, Direction: Left, Steps: 51, Expected: 1},   // Left overflow once
		{InitialState: 50, Direction: Right, Steps: 200, Expected: 2}, // Right underflow multiple
		{InitialState: 50, Direction: Left, Steps: 200, Expected: 2},  // Left overflow multiple
		{InitialState: 50, Direction: Right, Steps: 1000, Expected: 10},
		{InitialState: 50, Direction: Left, Steps: 1000, Expected: 10},
		{InitialState: 0, Direction: Left, Steps: 1, Expected: 0},   // Was already at 0
		{InitialState: 0, Direction: Right, Steps: 1, Expected: 0},  // Was already at 0
		{InitialState: 99, Direction: Right, Steps: 1, Expected: 1}, // Land on 0
		{InitialState: 1, Direction: Left, Steps: 1, Expected: 1},   // Land on 0
	}

	for i, test := range testRotations {
		dial.Position = test.InitialState
		count := dial.Turn(test.Direction, test.Steps)
		if count != test.Expected {
			t.Errorf("[%d] Expected to have crossed %d times, got %d (from %d, steps %d, direction %d)", i+1, test.Expected, count, test.InitialState, test.Steps, test.Direction)
		}
	}
}
