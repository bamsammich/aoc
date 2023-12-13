package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day01Command() *cobra.Command {
	return &cobra.Command{
		Use:   "1",
		Short: "2023 - Day 1",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day01Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day01Solution() error {
	lines, err := aocInput.Strings(2023, 1)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}

	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	solution1 := 0
	solution2 := 0
	for _, l := range lines {
		solution1 += day01LineValue(digits[:9], l)
		solution2 += day01LineValue(digits, l)
	}
	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
	return nil
}

func day01LineValue(numbers []string, in string) int {
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	rgxFirst := regexp.MustCompile(fmt.Sprintf("(%s)", strings.Join(numbers, "|")))
	rgxLast := regexp.MustCompile(fmt.Sprintf(".*%s", rgxFirst))
	num, _ := strconv.Atoi(fmt.Sprintf("%s%s",
		replacer.Replace(rgxFirst.FindStringSubmatch(in)[1]),
		replacer.Replace(rgxLast.FindStringSubmatch(in)[1])),
	)

	return num
}
