package year2023

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day06Command() *cobra.Command {
	return &cobra.Command{
		Use:   "6",
		Short: "2023 - Day 6",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day06Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day06Solution() error {
	lines, err := aocInput.Strings(2023, 6)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}

	times, distances := make([]int, 4), make([]int, 4)
	for i, s := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		t, _ := strconv.Atoi(s)
		times[i] = t
	}

	for i, s := range strings.Fields(strings.Split(lines[1], ":")[1]) {
		d, _ := strconv.Atoi(s)
		distances[i] = d
	}

	permute := func(dist, time int) int {
		out := 0
		for i := 1; i <= time; i++ {
			if i*(time-i) > dist {
				out++
			}
		}
		return out
	}

	s1 := 1
	for i := 0; i < len(times); i++ {
		s1 *= permute(distances[i], times[i])
	}

	s2t, err := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	if err != nil {
		return err
	}

	s2d, err := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	if err != nil {
		return err
	}

	s2 := permute(s2d, s2t)

	fmt.Println("Solution 1:", s1)
	fmt.Println("Solution 2:", s2)
	return nil
}
