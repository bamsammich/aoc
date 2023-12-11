package year2023

import (
	"fmt"
	"image"
	"time"

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
	}
	fmt.Println(diagram[start])
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
