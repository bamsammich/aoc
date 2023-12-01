package commands

import (
	"github.com/bamsammich/aoc/cmd/aoc/commands/year2023"
	"github.com/spf13/cobra"
)

func NewAOCCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aoc",
		Short: "bamsammich's Adevent of Code solutions by year and day",
	}
	cmd.AddCommand(
		year2023.New2023Command(),
	)
	return cmd
}
