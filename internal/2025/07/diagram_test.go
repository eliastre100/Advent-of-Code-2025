package _7

import "testing"

func TestSplitCount(t *testing.T) {
	tests := []struct {
		Diagram           Diagram
		ExpectedSplit     int
		ExpectedTimelines int
	}{
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{}}, 0, 1},
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{
			1: {1},
		}}, 1, 2},
		{Diagram{Width: 3, Height: 3, Start: Position{X: 1, Y: 0}, Splitters: map[int][]int{
			1: {1},
			2: {0, 1, 2},
		}}, 3, 4},
	}

	for i, test := range tests {
		splits, timelines := test.Diagram.SplitCount()
		if splits != test.ExpectedSplit {
			t.Errorf("[%d] Expected %d splits, got %d", i+1, test.ExpectedSplit, splits)
		}
		if timelines != test.ExpectedTimelines {
			t.Errorf("[%d] Expected %d timelines, got %d", i+1, test.ExpectedTimelines, timelines)
		}
	}
}
