package _3

import "testing"

func TestMaxBatteryJoltage(t *testing.T) {
	tests := []struct {
		Battery  Battery
		Expected int
	}{
		{Battery: Battery{Joltages: []int{1, 1}}, Expected: 11},
		{Battery: Battery{Joltages: []int{8, 9, 1}}, Expected: 91},
	}

	for _, test := range tests {
		value := MaxBatteryJoltage(test.Battery)
		if value != test.Expected {
			t.Errorf("Expected %d for battery %+v, got %d", test.Expected, test.Battery, value)
		}
	}
}
