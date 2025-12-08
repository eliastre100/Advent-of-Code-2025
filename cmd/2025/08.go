package _025

import (
	"slices"

	"github.com/charmbracelet/log"
	_8 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/08"
	"github.com/spf13/cobra"
)

var D08Cmd = &cobra.Command{
	Use:   "08",
	Short: "Implementations of 2025's day 8 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 08/2025 with input file %s", cmd.Flag("input").Value)

		boxes, err := _8.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		linkBudget, _ := cmd.Flags().GetInt("links")
		part, _ := cmd.Flags().GetInt("part")
		if part == 1 {
			d08Part1(boxes, linkBudget)
		} else {
			d08Part2(boxes)
		}

	},
}

func init() {
	D08Cmd.Flags().StringP("input", "i", "", "Path to the input file")
	D08Cmd.Flags().IntP("links", "l", 1000, "Number of link to be made")
	D08Cmd.Flags().IntP("part", "p", 1, "What part of the puzzle to solve (1 or 2)")

	D08Cmd.MarkFlagRequired("input")
	D08Cmd.MarkFlagFilename("input")
	D08Cmd.MarkFlagRequired("part")
}

func d08Part1(boxes []*_8.JunctionBox, linkBudget int) {
	_8.LinkJunctionBoxes(boxes, linkBudget)

	circuits := collectCircuits(boxes)

	result := 1
	for i, circuit := range circuits {
		if i < 3 {
			result *= len(circuit.JunctionBoxes)
		}
		log.Debugf("[%d] %d", i, len(circuit.JunctionBoxes))
	}
	log.Infof("The solution is %d", result)
}

func collectCircuits(boxes []*_8.JunctionBox) []*_8.Circuit {
	var circuits []*_8.Circuit

	for _, box := range boxes {
		if !slices.Contains(circuits, box.Circuit) {
			circuits = append(circuits, box.Circuit)
		}
	}
	slices.SortFunc(circuits, func(a, b *_8.Circuit) int {
		return len(b.JunctionBoxes) - len(a.JunctionBoxes)
	})
	return circuits
}

func d08Part2(boxes []*_8.JunctionBox) {
	link := _8.LinkAllJunctionBoxes(boxes)

	if link.A == nil {
		log.Fatal("Unable to make a single circuit")
	}

	log.Infof("The last two junction boxes are (%d,%d,%d) and (%d,%d,%d). The result is %d", link.A.X, link.A.Y, link.A.Z, link.B.X, link.B.Y, link.B.Z, link.A.X*link.B.X)
}
