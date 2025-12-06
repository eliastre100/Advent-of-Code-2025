package _6

import (
	"slices"
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]Problem, error) {
	var problems []Problem

	if err := parsing.ReadInputLines(path, func(line string) error {
		chunks := slices.Collect(func(yield func(string) bool) {
			for _, chunk := range strings.Split(line, " ") {
				if len(chunk) != 0 {
					yield(chunk)
				}
			}
		})
		// might need to remove empty chunks
		if len(problems) == 0 {
			problems = make([]Problem, len(chunks))
		}

		if slices.Contains([]string{"*", "+"}, chunks[0]) {
			setOperator(problems, chunks)
		} else {
			if err := addOperands(problems, chunks); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return problems, nil
}

func setOperator(problems []Problem, chunks []string) {
	operators := map[string]Operator{
		"+": add,
		"*": mult,
	}
	for i, chunk := range chunks {
		problems[i].Operator = operators[chunk]
	}
}

func addOperands(problems []Problem, chunks []string) error {
	for i, chunk := range chunks {
		v, err := strconv.Atoi(chunk)
		if err != nil {
			return err
		}
		problems[i].Values = append(problems[i].Values, v)
	}
	return nil
}
