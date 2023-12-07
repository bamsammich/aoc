package year2023

import (
	"fmt"

	"github.com/echojc/aocutil"
	"github.com/spf13/cobra"
)

var aocInput *aocutil.Input

func New2023Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2023",
		Short: "2023 Solutions",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			i, err := aocutil.NewInputFromFile("session_id")
			if err != nil {
				return fmt.Errorf("can't get AoC input: %w", err)
			}
			aocInput = i
			return nil
		},
	}

	cmd.AddCommand(
		New2023Day01Command(),
		New2023Day02Command(),
		New2023Day03Command(),
		New2023Day04Command(),
		New2023Day05Command(),
		New2023Day06Command(),
	)
	return cmd
}
