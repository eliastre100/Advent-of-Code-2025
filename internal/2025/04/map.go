package _4

import "github.com/charmbracelet/log"

type Map struct {
	Width         int
	Height        int
	RollPositions map[int]map[int]bool // y,x -> presence
}

func (m *Map) Cleanup() int {
	removedRollCount := 0
	rollPositions := make(map[int]map[int]bool)

	for y := range m.Height {
		rollPositions[y] = make(map[int]bool)
		for x := range m.Width {
			rollPositions[y][x] = m.RollPositions[y][x]

			if m.RollPositions[y][x] && RollAccessible(*m, x, y) {
				log.Debugf("Removing roll at %d,%d", x, y)
				rollPositions[y][x] = false
				removedRollCount++
			}
		}
	}
	m.RollPositions = rollPositions

	return removedRollCount
}

func (m *Map) String() string {
	str := ""
	for y := range m.Height {
		for x := range m.Width {
			if m.RollPositions[y][x] {
				str += "@"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func RollAccessible(m Map, x int, y int) bool {
	adjacent := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			testedX := x + dx
			testedY := y + dy

			if dx == 0 && dy == 0 {
				continue
			}
			if testedX >= 0 && testedX < m.Width &&
				testedY >= 0 && testedY < m.Height &&
				m.RollPositions[y+dy][x+dx] {
				adjacent++
			}
		}
	}
	return adjacent < 4
}
