package _025

import (
	"strconv"

	"github.com/charmbracelet/log"
	_1 "github.com/eliastre100/Advent-of-Code-2025/internal/2025/01"
	"github.com/eliastre100/Advent-of-Code-2025/pkg/parsing"
	"github.com/spf13/cobra"
)

var D01Cmd = &cobra.Command{
	Use:   "01",
	Short: "Implementations of 2025's day 1 puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Solving day 01/2025 with input file %s", cmd.Flag("input").Value)

		method := cmd.Flag("method").Value.String() == "0x434C49434B"
		dial := _1.Dial{Position: 50}
		count := 0

		if err := parsing.ReadInputLines(cmd.Flag("input").Value.String(), func(line string) error {
			log.Debug("New instruction", "instruction", line)
			if len(line) == 0 {
				return nil
			}
			direction := _1.Left
			if line[0] == 'R' {
				direction = _1.Right
			}
			steps, err := strconv.Atoi(line[1:len(line)])
			if err != nil {
				log.Fatal("The step number is invalid", "steps", line[1:len(line)], "error", err)
			}
			log.Debug("Moving the dial", "direction", direction, "steps", steps)
			crossed := dial.Turn(direction, steps)
			log.Debug("Dial landed", "position", dial.Position, "crossed", crossed)
			if method {
				count += crossed
			} else {
				if dial.Position == 0 {
					count += 1
				}
			}
			log.Debug("Dial moved", "position", dial.Position)
			log.Debug("")
			return nil
		}); err != nil {
			log.Fatal(err)
		}

		log.Infof("The dial pointed %d times at 0", count)
	},
}

func init() {
	D01Cmd.Flags().StringP("input", "i", "", "Path to the input file")
	D01Cmd.Flags().String("method", "0x434C49434C", "Password method to use")

	D01Cmd.MarkFlagRequired("input")
	D01Cmd.MarkFlagFilename("input")
}
