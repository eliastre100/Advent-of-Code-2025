package _8

type Circuit struct {
	JunctionBoxes []*JunctionBox
}

func linkBoxes(a *Circuit, b *Circuit) {
	a.JunctionBoxes = append(a.JunctionBoxes, b.JunctionBoxes...)
	for _, box := range b.JunctionBoxes {
		box.Circuit = a
	}
}
