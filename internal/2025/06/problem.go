package _6

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type Operator func(int, int) int

type Problem struct {
	Values   []int
	Operator Operator
}

func mult(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func Solve(p Problem) (int, error) {
	if p.Operator == nil {
		return 0, fmt.Errorf("no operator defined")
	}

	result := p.Values[0]
	for _, v := range p.Values[1:] {
		result = p.Operator(result, v)
	}
	log.Debugf("solving %+v gives %d", p.Values, result)
	return result, nil
}
