package _5

type Ingredient int

type Range struct {
	Start int
	End   int
}

type Database []Range

func (d Database) IsFresh(i Ingredient) bool {
	for _, r := range d {
		if r.Start <= int(i) && int(i) <= r.End {
			return true
		}
	}
	return false
}

func (d *Database) Insert(r Range) {
	*d = append(*d, r)
}
