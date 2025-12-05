package _5

import (
	"reflect"
	"slices"
	"testing"
)

func TestDatabase_IsFresh(t *testing.T) {
	tests := []struct {
		Range      Range
		Ingredient Ingredient
		Expected   bool
	}{
		{Range: Range{10, 15}, Ingredient: 0, Expected: false},
		{Range: Range{10, 15}, Ingredient: 10, Expected: true},
		{Range: Range{10, 15}, Ingredient: 12, Expected: true},
		{Range: Range{10, 15}, Ingredient: 15, Expected: true},
		{Range: Range{10, 15}, Ingredient: 20, Expected: false},
	}

	for _, test := range tests {
		db := Database{test.Range}
		fresh := db.IsFresh(test.Ingredient)
		if fresh != test.Expected {
			t.Errorf("Expected %t for %d in %+v, got %t", test.Expected, test.Ingredient, test.Range, fresh)
		}
	}
}

func TestDatabase_ExportFreshIds(t *testing.T) {
	tests := []struct {
		Ranges   []Range
		Expected []Ingredient
	}{
		{Ranges: []Range{Range{10, 15}, Range{20, 25}}, Expected: []Ingredient{10, 11, 12, 13, 14, 15, 20, 21, 22, 23, 24, 25}},
	}

	for _, test := range tests {
		var db Database
		for _, r := range test.Ranges {
			db.Insert(r)
		}
		result := db.ExportFreshIds()
		slices.Sort(result)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Expected %+v, got %+v", test.Expected, result)
		}
	}
}
