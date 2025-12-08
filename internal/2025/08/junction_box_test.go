package _8

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		A        JunctionBox
		B        JunctionBox
		Expected int
	}{
		{JunctionBox{0, 0, 0, nil}, JunctionBox{1, 1, 1, nil}, 3},
		{JunctionBox{0, 0, 0, nil}, JunctionBox{-1, 1, 1, nil}, 3},
	}

	for i, test := range tests {
		result := Distance(test.A, test.B)
		if result != test.Expected {
			t.Errorf("[%d] Expected %d, got %d", i+1, test.Expected, result)
		}
	}
}
