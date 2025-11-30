package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	_025 "github.com/eliastre100/Advent-of-Code-2025/cmd/2025"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Personal AOC solution implementation",
	Long: `Advent of Code is a series of small programming puzzles for a variety of skill sets 
and skill levels that can be solved in any programming language you like.
This CLI is my own implementations to solves theses puzzles, and you should avoid 
using it to solve yours and rather give it a try yourself.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("verbose")
		if debug {
			log.SetLevel(log.DebugLevel)
		}
		cmd.SetContext(context.WithValue(cmd.Context(), "startedAt", time.Now()))
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		startedAt, ok := cmd.Context().Value("startedAt").(time.Time)
		if ok {
			log.Infof("Completed in %s", time.Since(startedAt).Round(time.Nanosecond))
		} else {
			log.Warn("Unable to get startedAt value from context. Cannot compute completion time.")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose mode")
	rootCmd.AddCommand(_025.Year2025Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
