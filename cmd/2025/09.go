package _025

import (
	"maps"
	"slices"

	"github.com/charmbracelet/log"
	_9 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/09"
	"github.com/spf13/cobra"
)

var D09Cmd = &cobra.Command{
	Use:   "09",
	Short: "Implementations of 2025's day 9 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 09/2025 with input file %s", cmd.Flag("input").Value)

		tiles, err := _9.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		squares := _9.BuildSquares(tiles)
		maxSize := slices.Max(slices.Sorted(maps.Keys(squares)))

		log.Infof("The largest square possible have a size of %d", maxSize)
	},
}

func init() {
	D09Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D09Cmd.MarkFlagRequired("input")
	D09Cmd.MarkFlagFilename("input")
}
