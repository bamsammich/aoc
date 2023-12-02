package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day02Command() *cobra.Command {
	return &cobra.Command{
		Use:   "2",
		Short: "2023 - Day 2",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day02Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day02Solution() error {
	lines, err := aocInput.Strings(2023, 2)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}
	s1 := 0
	s2 := 0
	for _, l := range lines {
		rgx := regexp.MustCompile(`([\d]+) (red|green|blue)`)
		id, _ := strconv.Atoi(regexp.MustCompile(`Game ([\d]+):`).FindStringSubmatch(l)[1])
		maxes := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, draws := range strings.Split(strings.Split(l, ": ")[1], "; ") {
			for _, draw := range strings.Split(draws, ", ") {
				matches := rgx.FindStringSubmatch(draw)
				count, _ := strconv.Atoi(matches[1])
				if val := maxes[matches[2]]; val < count {
					maxes[matches[2]] = count
				}
			}
		}
		if maxes["red"] <= 12 && maxes["green"] <= 13 && maxes["blue"] <= 14 {
			s1 += id
		}
		pow := 1
		for _, v := range maxes {
			if v != 0 {
				pow *= v
			}
		}
		s2 += pow
	}
	fmt.Printf("Solution 1: %d\n", s1)
	fmt.Printf("Solution 2: %d\n", s2)
	return nil
}

type day2Game struct {
	ID    int
	Draws []map[string]int
}
