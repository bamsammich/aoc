package year2023

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day08Command() *cobra.Command {
	return &cobra.Command{
		Use:   "8",
		Short: "2023 - Day 8",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day08Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day08Solution() error {
	lines, err := aocInput.Strings(2023, 8)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}

	turns := strings.Split(lines[0], "")
	nodes := map[string]map[string]string{}
	for _, input := range lines[2:] {
		parts := strings.Split(input, " = ")
		fields := strings.Fields(strings.NewReplacer("(", "", ")", "", ",", "").Replace(parts[1]))
		nodes[parts[0]] = map[string]string{"L": fields[0], "R": fields[1]}
	}

	findEnd := func(start, end string) int {
		result := 1
		for n := range nodes {
			if !strings.HasSuffix(n, start) {
				continue
			}

			steps := 0
			for !strings.HasSuffix(n, end) {
				n = nodes[n][turns[steps%len(turns)]]
				steps++
			}

			result = lcm(result, steps)
		}
		return result
	}

	fmt.Println("Solution 1:", findEnd("AAA", "ZZZ"))
	fmt.Println("Solution 2:", findEnd("A", "Z"))

	return nil
}

// Stolen from github.com/mnml/aoc; Was the missing piece in my solution.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
