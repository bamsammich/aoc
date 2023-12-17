package utils

import (
	"fmt"
	"strings"
)

type Grid map[int][]string

func ToGrid(in []string) Grid {
	out := make(Grid)
	for i, s := range in {
		out[i] = strings.Split(s, "")
	}
	return out
}

func (g Grid) String() string {
	var out string
	for i := 0; i < len(g); i++ {
		out += fmt.Sprintln(strings.Join(g[i], ""))
	}
	return out
}

func (g Grid) Update(x, y int, value string) {
	g[y][x] = value
}
