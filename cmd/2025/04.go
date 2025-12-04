package _025

import (
	"github.com/charmbracelet/log"
	_4 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/04"
	"github.com/spf13/cobra"
)

var D04Cmd = &cobra.Command{
	Use:   "04",
	Short: "Implementations of 2025's day 4 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 04/2025 with input file %s", cmd.Flag("input").Value)

		m, err := _4.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Failed to parse map: ", err)
		}

		accessibleRollCount := 0
		for {
			log.Debugf("MAP:\n")
			print(m.String())
			removedRolls := m.Cleanup()
			log.Debugf("Removed %d rolls", removedRolls)
			if removedRolls == 0 {
				break
			}
			accessibleRollCount += removedRolls
		}

		log.Infof("Number of accessible rolls: %d", accessibleRollCount)
	},
}

func init() {
	D04Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D04Cmd.MarkFlagRequired("input")
	D04Cmd.MarkFlagFilename("input")
}
