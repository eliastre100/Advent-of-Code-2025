package _2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
)

func Parse(path string) ([]Range, error) {
	groups, err := parsing.ReadGroups(path, ",")
	if err != nil {
		return nil, err
	}
	ranges := make([]Range, len(groups))

	for i, group := range groups {
		ends := strings.Split(group, "-")
		if len(ends) != 2 {
			return nil, fmt.Errorf("invalid range: expected two boundaries, got %d", len(group))
		}
		start, err := strconv.Atoi(ends[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(ends[1])
		if err != nil {
			return nil, err
		}
		ranges[i] = Range{
			Start: start,
			End:   end,
		}
	}

	return ranges, nil
}
