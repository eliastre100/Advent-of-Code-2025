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
	},
}

func init() {
	D08Cmd.Flags().StringP("input", "i", "", "Path to the input file")
	D08Cmd.Flags().IntP("links", "l", 10, "Number of link to be made")

	D08Cmd.MarkFlagRequired("input")
	D08Cmd.MarkFlagFilename("input")
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
