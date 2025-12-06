package _025

import (
	"github.com/charmbracelet/log"
	_6 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/06"
	"github.com/spf13/cobra"
)

var D06Cmd = &cobra.Command{
	Use:   "06",
	Short: "Implementations of 2025's day 6 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 06/2025 with input file %s", cmd.Flag("input").Value)

		problems, err := _6.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		total := 0
		for i, problem := range problems {
			result, err := _6.Solve(problem)
			if err != nil {
				log.Fatal("Unable to solve problem", "err", err)
			}
			log.Debugf("The solution to problem %d is %d", i+1, result)
			total += result
		}

		log.Infof("The total of all solutions is %d", total)
	},
}

func init() {
	D06Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D06Cmd.MarkFlagRequired("input")
	D06Cmd.MarkFlagFilename("input")
}

// 9581313735384 low
