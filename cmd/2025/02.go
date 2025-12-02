package _025

import (
	"github.com/charmbracelet/log"
	_2 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/02"
	"github.com/spf13/cobra"
)

var D02Cmd = &cobra.Command{
	Use:   "02",
	Short: "Implementations of 2025's day 2 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 02/2025 with input file %s", cmd.Flag("input").Value)

		invalidValuesSum := 0
		ranges, err := _2.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		for _, r := range ranges {
			invalidValues := _2.InvalidValuesInRange(r)
			if len(invalidValues) > 0 {
				for _, v := range invalidValues {
					invalidValuesSum += v
				}
			}
			log.Debugf("Range: %+v have invalid values %+v", r, invalidValues)
		}

		log.Infof("The sum of invalid values is %d", invalidValuesSum)
	},
}

func init() {
	D02Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D02Cmd.MarkFlagRequired("input")
	D02Cmd.MarkFlagFilename("input")
}
