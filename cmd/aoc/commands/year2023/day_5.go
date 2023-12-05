package year2023

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day05Command() *cobra.Command {
	return &cobra.Command{
		Use:   "5",
		Short: "2023 - Day 5",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day05Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day05Solution() error {
	lines, err := aocInput.Strings(2023, 5)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}
	input := strings.Join(lines, "\n")

	almanac := day5Almanac{}
	for _, m := range regexp.MustCompile(`\n{2,}`).Split(input, -1)[1:] {
		parts := regexp.MustCompile(`(.*) map: (.*)`).FindStringSubmatch(strings.ReplaceAll(m, "\n", " "))
		nums := strings.Fields(parts[2])
		for i := 0; i < len(nums); i += 3 {
			dest, _ := strconv.Atoi(nums[i])
			src, _ := strconv.Atoi(nums[i+1])
			lenRange, _ := strconv.Atoi(nums[i+2])
			almanac[parts[1]] = append(almanac[parts[1]], day5Mapping{
				Src:  Range{Min: src, Max: src + lenRange},
				Dest: Range{Min: dest, Max: dest + lenRange},
			})
		}
	}

	s1, s2 := math.MaxInt, math.MaxInt
	seedNumbers := strings.Fields(regexp.MustCompile(`seeds: (.*)\n`).FindStringSubmatch(input)[1])
	for i := 0; i < len(seedNumbers); i++ {
		lookup, _ := strconv.Atoi(seedNumbers[i])

		if loc := almanac.Location(lookup); loc < s1 {
			s1 = loc
		}
		// This is slow... I'm sure there's a better way to do this
		if i%2 == 0 {
			lenRange, _ := strconv.Atoi(seedNumbers[i+1])
			for j := lookup; j < lookup+lenRange; j++ {
				if loc := almanac.Location(j); loc < s2 {
					s2 = loc
				}
			}
		}
	}

	fmt.Printf("Solution 1: %d\n", s1)
	fmt.Printf("Solution 2: %d\n", s2)
	return nil
}

type Range struct {
	Min, Max int
}

type day5Mapping struct {
	Src, Dest Range
}

func (m day5Mapping) SourceIndex(val int) int {
	if m.Src.Min <= val && val <= m.Src.Max {
		return val - m.Src.Min
	}
	return -1
}

func (m day5Mapping) DestLookup(idx int) int {
	return m.Dest.Min + idx
}

var day5LookupOrder = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

type day5Almanac map[string][]day5Mapping

func (a day5Almanac) Location(lookup int) int {
	for _, l := range day5LookupOrder {
		for _, mapping := range a[l] {
			if idx := mapping.SourceIndex(lookup); idx != -1 {
				lookup = mapping.DestLookup(idx)
				break
			}
		}
	}
	return lookup
}
