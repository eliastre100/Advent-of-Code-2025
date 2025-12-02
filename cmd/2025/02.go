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

		invalidDuplicate := 0
		invalidRepeated := 0
		ranges, err := _2.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		for _, r := range ranges {
			invalidDuplicate += invalidSum(r, _2.NotDuplicate, "DUPLICATE")
			invalidRepeated += invalidSum(r, _2.NotRepeating, "REPEATED")
		}

		log.Infof("The sum of invalid values by duplication is %d and %d by repetition", invalidDuplicate, invalidRepeated)
	},
}

func invalidSum(r _2.Range, validator _2.Validator, scope string) int {
	sum := 0
	invalidValues := _2.InvalidValuesInRange(r, _2.NotDuplicate)

	if len(invalidValues) > 0 {
		log.Debugf("[%s] Range: %+v have invalid values %+v", scope, r, invalidValues)
		for _, v := range invalidValues {
			sum += v
		}
	}
	return sum
}

func init() {
	D02Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D02Cmd.MarkFlagRequired("input")
	D02Cmd.MarkFlagFilename("input")
}
