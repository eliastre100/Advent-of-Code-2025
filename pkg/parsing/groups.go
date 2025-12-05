package parsing

import (
	"io"
	"os"
	"strings"
)

func ReadGroups(path string, delimiter string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), delimiter), nil
}

func ReadEmptyLineGroups(path string) ([][]string, error) {
	var groups [][]string
	var current []string

	if err := ReadInputLines(path, func(line string) error {
		if len(line) == 0 {
			groups = append(groups, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	groups = append(groups, current)
	return groups, nil
}
