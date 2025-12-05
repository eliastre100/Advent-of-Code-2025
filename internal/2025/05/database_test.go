package _5

import "testing"

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
