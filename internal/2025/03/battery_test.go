package _3

import "testing"

func TestMaxBatteryJoltage(t *testing.T) {
	tests := []struct {
		Battery  Battery
		Joltages int
		Expected int
	}{
		{Battery: Battery{Joltages: []int{1, 1}}, Joltages: 2, Expected: 11},
		{Battery: Battery{Joltages: []int{8, 9, 1}}, Joltages: 2, Expected: 91},
		{Battery: Battery{Joltages: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}}, Joltages: 12, Expected: 811111111119},
	}

	for _, test := range tests {
		value := MaxBatteryJoltage(test.Battery, test.Joltages)
		if value != test.Expected {
			t.Errorf("Expected %d for battery %+v, got %d", test.Expected, test.Battery, value)
		}
	}
}
