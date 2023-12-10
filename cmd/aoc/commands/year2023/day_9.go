package year2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day09Command() *cobra.Command {
	return &cobra.Command{
		Use:   "9",
		Short: "2023 - Day 9",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day09Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day09Solution() error {
	lines, err := aocInput.Strings(2023, 9)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}

	s1, s2 := 0, 0
	for _, l := range lines {
		data := strings.Fields(l)
		nums := make([]int, len(data))
		for i, d := range data {
			n, _ := strconv.Atoi(d)
			nums[i] = n
		}
		s1 += day9Diffs(nums)
		slices.Reverse(nums)
		s2 += day9Diffs(nums)
	}

	fmt.Println("Solution 1:", s1)
	fmt.Println("Solution 2:", s2)
	return nil
}

func day9Diffs(nums []int) int {
	next := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		next[i] = nums[i+1] - nums[i]
	}
	if !slices.ContainsFunc(next, func(i int) bool { return i != 0 }) {
		return nums[len(nums)-1]
	}
	return nums[len(nums)-1] + day9Diffs(next)
}
