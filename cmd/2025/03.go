package _025

import (
	"github.com/charmbracelet/log"
	_3 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/03"
	"github.com/spf13/cobra"
)

var D03Cmd = &cobra.Command{
	Use:   "03",
	Short: "Implementations of 2025's day 3 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 03/2025 with input file %s", cmd.Flag("input").Value)

		batteries, err := _3.Parse(cmd.Flag("input").Value.String())
		if err != nil {
			log.Fatal("Failed to parse batteries: ", err)
		}

		powerOf2 := 0
		powerOf12 := 0
		for i, bat := range batteries {
			batteryPowerOf2 := _3.MaxBatteryJoltage(bat, 2)
			batteryPowerOf12 := _3.MaxBatteryJoltage(bat, 12)

			log.Debugf("Power of 2(%d) and 12(%d) for Battery %d", batteryPowerOf2, batteryPowerOf12, i)
			powerOf2 += batteryPowerOf2
			powerOf12 += batteryPowerOf12
		}

		log.Infof("The total output Joltage is 2(%d) and 12(%d)", powerOf2, powerOf12)
	},
}

func init() {
	D03Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D03Cmd.MarkFlagRequired("input")
	D03Cmd.MarkFlagFilename("input")
}
