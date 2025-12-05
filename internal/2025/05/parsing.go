package _5

import (
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) (*Database, []Ingredient, error) {
	groups, err := parsing.ReadEmptyLineGroups(path)
	if err != nil {
		return nil, nil, err
	}

	db, err := buildDb(groups[0])
	if err != nil {
		return nil, nil, err
	}

	ingredients, err := buildIngredients(groups[1])
	if err != nil {
		return nil, nil, err
	}

	return db, ingredients, nil
}

func buildDb(definitions []string) (*Database, error) {
	var db Database
	for _, ingredients := range definitions {
		bounds := strings.Split(ingredients, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, err
		}
		db.Insert(Range{start, end})
	}
	return &db, nil
}

func buildIngredients(definitions []string) ([]Ingredient, error) {
	var ingredients []Ingredient

	for _, ingredient := range definitions {
		ingredient, err := strconv.Atoi(ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, Ingredient(ingredient))
	}
	return ingredients, nil
}
