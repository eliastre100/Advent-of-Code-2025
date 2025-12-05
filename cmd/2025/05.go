package _025

import (
	"github.com/charmbracelet/log"
	_5 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/05"
	"github.com/spf13/cobra"
)

var D05Cmd = &cobra.Command{
	Use:   "05",
	Short: "Implementations of 2025's day 5 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 05/2025 with input file %s", cmd.Flag("input").Value)

		db, ingredients, err := _5.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Unable to parse input file", "err", err)
		}

		freshCount := 0
		for _, ingredient := range ingredients {
			if db.IsFresh(ingredient) {
				freshCount += 1
			}
		}

		log.Infof("There are %d ingredients that are still fresh", freshCount)
	},
}

func init() {
	D05Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D05Cmd.MarkFlagRequired("input")
	D05Cmd.MarkFlagFilename("input")
}
