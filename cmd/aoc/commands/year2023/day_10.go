package year2023

import (
	"fmt"
	"image"
	"slices"
	"strings"
	"time"

	"github.com/bamsammich/aoc/pkg/utils"
	"github.com/spf13/cobra"
)

func New2023Day10Command() *cobra.Command {
	return &cobra.Command{
		Use:   "10",
		Short: "2023 - Day 10",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day010Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day010Solution() error {
	lines, err := aocInput.Strings(2023, 10)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}
	diagram := map[image.Point][]image.Point{}
	var start image.Point
	for y := range lines {
		for x, c := range lines[y] {
			switch c {
			case 'S':
				start = image.Point{x, y}
			case '|':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x, y - 1},
					image.Point{x, y + 1},
				)
			case '-':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x - 1, y},
					image.Point{x + 1, y},
				)
			case 'L':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x, y + 1},
					image.Point{x + 1, y},
				)
			case 'J':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x, y + 1},
					image.Point{x - 1, y},
				)
			case '7':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x, y - 1},
					image.Point{x - 1, y},
				)
			case 'F':
				diagram[image.Point{x, y}] = bound(len(lines[y])-1, len(lines)-1,
					image.Point{x, y - 1},
					image.Point{x + 1, y},
				)
			}
		}
		lines[y] = strings.NewReplacer("7", "↰", "F", "↱", "J", "↲", "L", "↳").Replace(lines[y])
	}
	fmt.Println(start)
	// s1 := 0
	grid := updateGrid(make([][]string, len(lines)), len(lines[0]), start, "S")
	for _, p := range bound(len(lines[0])-1, len(lines)-1, surrounding(start)...) {
		if adj, ok := diagram[p]; ok && slices.Contains(adj, start) {
			grid = updateGrid(grid, len(lines)-1, p, ".")
			last := start
			current := p
			pipe := []image.Point{last, current}
			for current != start {
				grid = updateGrid(grid, len(lines)-1, current, strings.Split(lines[current.Y], "")[current.X])
				nextE := slices.IndexFunc(diagram[current], func(p image.Point) bool { return current != last })
				last = current
				current = diagram[current][nextE]
				pipe = append(pipe, current)

				utils.Printlns(grid[0:30]...)
			}
		}
	}
	return nil
}

func bound(xmax, ymax int, pts ...image.Point) []image.Point {
	out := []image.Point{}
	for _, p := range pts {
		if p.X <= xmax && p.Y <= ymax && p.X >= 0 && p.Y >= 0 {
			out = append(out, p)
		}
	}

	return out
}

func surrounding(pt image.Point) []image.Point {
	return []image.Point{
		{pt.X - 1, pt.Y + 1}, {pt.X, pt.Y + 1}, {pt.X + 1, pt.Y + 1},
		{pt.X - 1, pt.Y}, {pt.X + 1, pt.Y},
		{pt.X - 1, pt.Y - 1}, {pt.X, pt.Y - 1}, {pt.X + 1, pt.Y - 1},
	}
}

func day10findStart(last, current, goal image.Point, diagram map[image.Point][]image.Point) []image.Point {
	if current == goal {
		return []image.Point{current}
	}
	nextE := slices.IndexFunc(diagram[current], func(p image.Point) bool { return p != last })
	next := diagram[current][nextE]
	return append([]image.Point{current}, day10findStart(current, next, goal, diagram)...)
}

func updateGrid(grid [][]string, xmax int, pt image.Point, value string) [][]string {
	if grid[pt.Y] == nil {
		grid[pt.Y] = strings.Split(strings.Repeat(" ", xmax), "")
	}
	grid[pt.Y][pt.X] = value
	return grid
}
