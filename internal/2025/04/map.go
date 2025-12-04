package _4

type Map struct {
	Width         int
	Height        int
	RollPositions map[int]map[int]bool // y,x -> presence
}

func RollAccessible(m Map, x int, y int) bool {
	adjacent := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if m.RollPositions[y+dy][x+dx] {
				adjacent++
			}
		}
	}
	return adjacent < 4
}
