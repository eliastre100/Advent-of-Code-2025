package _2

import (
	"reflect"
	"testing"
)

func TestNotDuplicate(t *testing.T) {
	tests := []struct {
		Value    int
		Expected bool
	}{
		{Value: 0, Expected: true},
		{Value: 1, Expected: true},
		{Value: 11, Expected: false},
		{Value: 123123, Expected: false},
		{Value: 1231234, Expected: true},
		{Value: 12351234, Expected: true},
	}

	for _, test := range tests {
		result := NotDuplicate(test.Value)
		if result != test.Expected {
			t.Errorf("Expected %t for %d, got %t", test.Expected, test.Value, result)
		}
	}
}

func TestNotRepeating(t *testing.T) {
	tests := []struct {
		Value    int
		Expected bool
	}{
		{Value: 0, Expected: true},
		{Value: 1, Expected: true},
		{Value: 11, Expected: false},
		{Value: 99, Expected: false},
	}

	for i, test := range tests {
		if NotRepeating(test.Value) != test.Expected {
			t.Errorf("[%d] Expected %t, got %t", i+1, test.Expected, NotRepeating(test.Value))
		}
	}
}

func TestPatternStart(t *testing.T) {
	tests := []struct {
		Value    string
		Pattern  string
		Expected []int
	}{
		{Value: "23", Pattern: "1", Expected: []int{}},
		{Value: "123", Pattern: "1", Expected: []int{0}},
		{Value: "123123", Pattern: "1", Expected: []int{0, 3}},
	}

	for i, test := range tests {
		starts := patternStarts(test.Value, test.Pattern)
		if len(test.Expected) != len(starts) {
			t.Errorf("[%d] Expected %+v, got %+v", i+1, test.Expected, starts)
		} else if len(test.Expected) != 0 && !reflect.DeepEqual(starts, test.Expected) {
			t.Errorf("[%d] Expected %+v, got %+v", i+1, test.Expected, starts)
		}
	}
}
