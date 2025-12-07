package _025

import (
	"github.com/charmbracelet/log"
	_7 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/07"
	"github.com/spf13/cobra"
)

var D07Cmd = &cobra.Command{
	Use:   "07",
	Short: "Implementations of 2025's day 7 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 07/2025 with input file %s", cmd.Flag("input").Value)

		diagram, err := _7.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input diagram", "err", err)
		}

		splitCount := diagram.Simulate()

		log.Infof("A simulation splits the beam %d times", splitCount)
	},
}

func init() {
	D07Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D07Cmd.MarkFlagRequired("input")
	D07Cmd.MarkFlagFilename("input")
}

// 9581313735384 low
