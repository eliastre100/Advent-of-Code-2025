package _025

import (
	"github.com/spf13/cobra"
)

var Year2025Cmd = &cobra.Command{
	Use:   "2025",
	Short: "Implementations of 2025's puzzles.",
}

func init() {
	Year2025Cmd.AddCommand(D01Cmd)
	Year2025Cmd.AddCommand(D02Cmd)
	Year2025Cmd.AddCommand(D03Cmd)
	Year2025Cmd.AddCommand(D04Cmd)
	Year2025Cmd.AddCommand(D05Cmd)
	Year2025Cmd.AddCommand(D06Cmd)
	Year2025Cmd.AddCommand(D07Cmd)
}
