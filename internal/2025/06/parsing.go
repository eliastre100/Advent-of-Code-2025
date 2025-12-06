package _6

import (
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]Problem, error) {
	data, width, err := loadInputData(path)
	if err != nil {
		return nil, err
	}
	operators := data[len(data)-1]
	data = rotate(data[:len(data)-1], width)

	var problems []Problem
	currentProblem := Problem{}
	for i := 0; i < len(data); i++ {
		line := strings.ReplaceAll(data[i], " ", "")
		if line == "" {
			problems = append(problems, currentProblem)
			currentProblem = Problem{}
		} else {
			operand, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			currentProblem.Values = append(currentProblem.Values, operand)
			if len(operators) > i && operators[i] != ' ' {
				setOperator(&currentProblem, string(operators[i]))
			}
		}
	}
	problems = append(problems, currentProblem)

	return problems, nil
}

func setOperator(problem *Problem, operator string) {
	operators := map[string]Operator{
		"+": add,
		"*": mult,
	}
	problem.Operator = operators[operator]
}

func loadInputData(path string) ([]string, int, error) {
	var data []string
	maxWidth := 0

	if err := parsing.ReadInputLines(path, func(line string) error {
		data = append(data, line)
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
		return nil
	}); err != nil {
		return nil, 0, err
	}

	return data, maxWidth, nil
}

func rotate(data []string, width int) []string {
	result := make([]string, width)

	for i := 0; i < width; i++ {
		for j := 0; j < len(data); j++ {
			if len(data[j]) > i {
				result[i] += string(data[j][i])
			}
		}
	}

	return result
}
