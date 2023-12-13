package year2023

import (
	"fmt"
	"image"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func New2023Day03Command() *cobra.Command {
	return &cobra.Command{
		Use:   "3",
		Short: "2023 - Day 3",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day03Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day03Solution() error {
	lines, err := aocInput.Strings(2023, 3)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}
	s1 := 0
	s2 := 0
	grid := toGrid(lines)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if regexp.MustCompile(`[^\.0-9]`).MatchString(grid[y][x]) {
				nums, err := findAdjacentNumbers(grid, image.Point{X: x, Y: y})
				if err != nil {
					return err
				}
				for _, n := range nums {
					s1 += n
				}
				if grid[y][x] == "*" && len(nums) == 2 {
					s2 += nums[0] * nums[1]
				}
			}
		}
	}
	fmt.Printf("Solution 1: %d\n", s1)
	fmt.Printf("Solution 2: %d\n", s2)
	return nil
}

func toGrid(in []string) map[int][]string {
	out := make(map[int][]string)
	for i, s := range in {
		out[i] = strings.Split(s, "")
	}
	return out
}

func findAdjacentNumbers(grid map[int][]string, point image.Point) ([]int, error) {
	search := []image.Point{
		{X: point.X - 1, Y: point.Y + 1}, {X: point.X, Y: point.Y + 1}, {X: point.X + 1, Y: point.Y + 1},
		{X: point.X - 1, Y: point.Y}, {X: point.X + 1, Y: point.Y},
		{X: point.X - 1, Y: point.Y - 1}, {X: point.X, Y: point.Y - 1}, {X: point.X + 1, Y: point.Y - 1},
	}
	var out []int
	for _, p := range search {
		if isOutOfBounds(grid, p) {
			continue
		}
		if isNumber(grid[p.Y][p.X]) {
			points := slices.Compact(
				append(
					// Seek left
					seekNumber(grid, p, func(pt image.Point) image.Point {
						return image.Point{X: pt.X - 1, Y: pt.Y}
					}),
					// Seek right
					seekNumber(grid, p, func(pt image.Point) image.Point {
						return image.Point{X: pt.X + 1, Y: pt.Y}
					})...,
				),
			)
			var numStr string
			for _, p := range points {
				numStr += grid[p.Y][p.X]
			}
			i, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			if !slices.Contains(out, i) {
				out = append(out, i)
			}
		}
	}
	return out, nil
}

func seekNumber(grid map[int][]string, p image.Point, seekFunc func(pt image.Point) image.Point) []image.Point {
	if !isNumber(grid[p.Y][p.X]) {
		return nil
	}
	points := []image.Point{p}
	next := seekFunc(p)
	if isOutOfBounds(grid, next) || !isNumber(grid[next.Y][next.X]) {
		return points
	}

	points = append(points, seekNumber(grid, next, seekFunc)...)
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	return points
}

func isNumber(s string) bool {
	return regexp.MustCompile(`\d`).MatchString(s)
}

func isOutOfBounds(grid map[int][]string, p image.Point) bool {
	return p.X < 0 || p.X >= len(grid[0]) || p.Y < 0 || p.Y >= len(grid)
}
