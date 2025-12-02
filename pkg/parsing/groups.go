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
