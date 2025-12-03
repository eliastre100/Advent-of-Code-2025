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

		power := 0
		for i, bat := range batteries {
			batteryPower := _3.MaxBatteryJoltage(bat)
			log.Debugf("Power of %d for Battery %d", batteryPower, i)
			power += batteryPower
		}

		log.Infof("The total output Joltage is %d", power)
	},
}

func init() {
	D03Cmd.Flags().StringP("input", "i", "", "Path to the input file")

	D03Cmd.MarkFlagRequired("input")
	D03Cmd.MarkFlagFilename("input")
}
