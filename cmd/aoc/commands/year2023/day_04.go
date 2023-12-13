package year2023

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day04Command() *cobra.Command {
	return &cobra.Command{
		Use:   "4",
		Short: "2023 - Day 4",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day04Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day04Solution() error {
	lines, err := aocInput.Strings(2023, 4)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}
	copies := make(map[int]int, len(lines))
	s1, s2 := 0, 0
	for i, l := range lines {
		var (
			nums    = strings.Split(strings.Split(l, ":")[1], "|")
			matches = 0
		)

		for _, n := range strings.Fields(nums[1]) {
			if slices.Contains(strings.Fields(nums[0]), n) {
				matches++
			}
		}

		copies[i]++
		for j := i + 1; j <= i+matches; j++ {
			copies[j] += copies[i]
		}

		if matches > 0 {
			s1 += int(math.Pow(2, float64(matches-1)))
		}
		s2 += copies[i]
	}

	fmt.Printf("Solution 1: %d\n", s1)
	fmt.Printf("Solution 2: %d\n", s2)
	return nil
}
