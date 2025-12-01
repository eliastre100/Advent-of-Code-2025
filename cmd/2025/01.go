package _025

import (
	"github.com/charmbracelet/log"
	_1 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/01"
	"github.com/spf13/cobra"
)

var D01Cmd = &cobra.Command{
	Use:   "01",
	Short: "Implementations of 2025's day 1 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 01/2025 with input file %s", cmd.Flag("input").Value)

		dial := _1.Dial{Position: 50}
		exactCount := 0
		passingCount := 0

		instructions, err := _1.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		for _, instruction := range instructions {
			passingCount += dial.Turn(instruction.Direction, instruction.Steps)
			if dial.Position == 0 {
				exactCount += 1
			}
		}

		log.Infof("The dial pointed exactly zero %d times, and passed it %d times", exactCount, passingCount)
	},
}

func init() {
	D01Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D01Cmd.MarkFlagRequired("input")
	D01Cmd.MarkFlagFilename("input")
}
