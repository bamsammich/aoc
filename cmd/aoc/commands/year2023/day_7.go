package year2023

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bamsammich/aoc/pkg/utils"
	"github.com/spf13/cobra"
)

func New2023Day07Command() *cobra.Command {
	return &cobra.Command{
		Use:   "7",
		Short: "2023 - Day 7",
		RunE: func(cmd *cobra.Command, args []string) error {
			start := time.Now()
			if err := day07Solution(); err != nil {
				return err
			}
			fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
			return nil
		},
	}
}

func day07Solution() error {
	lines, err := aocInput.Strings(2023, 7)
	if err != nil {
		return fmt.Errorf("failed to get puzzle input: %w", err)
	}

	cardVals := map[rune]int{
		'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}

	rounds := []round{}
	for _, l := range lines {
		var (
			parts  = strings.Fields(l)
			bid, _ = strconv.Atoi(parts[1])
			rnd    = round{
				hand: make([]rune, 5),
				bid:  bid,
			}
		)

		for i, c := range parts[0] {
			rnd.hand[i] = c
		}

		rounds = append(rounds, rnd)
	}

	sortFn := func(i, j int) bool {
		powI := rounds[i].Pow(cardVals)
		powJ := rounds[j].Pow(cardVals)
		if powI != powJ {
			return powI < powJ
		}
		for i, v := range rounds[i].hand {
			if cardVals[v] != cardVals[rounds[j].hand[i]] {
				return cardVals[v] < cardVals[rounds[j].hand[i]]
			}
		}
		return false
	}
	sort.Slice(rounds, sortFn)

	s1 := 0
	for rank, round := range rounds {
		s1 += round.bid * (rank + 1)
	}

	cardVals['J'] = 2
	sort.Slice(rounds, sortFn)
	s2 := 0
	for rank, round := range rounds {
		// fmt.Println(round.Pow(cardVals), round)
		s2 += round.bid * (rank + 1)
	}

	fmt.Println("Solution 1:", s1)
	fmt.Println("Solution 2:", s2)
	return nil
}

type round struct {
	hand []rune
	bid  int
}

func (r round) String() string {
	letters := []string{}
	for _, c := range r.hand {
		letters = append(letters, string(c))
	}
	return fmt.Sprintf("{%v %d}", letters, r.bid)
}

func (rnd round) Pow(cardVals map[rune]int) int {
	groups := map[rune]int{}
	for _, i := range rnd.hand {
		groups[i]++
	}

	if _, has := groups['J']; has && cardVals['J'] == 2 {
		groups[utils.MaxValueKey(groups)] += groups['J']
		delete(groups, 'J')
	}

	counts := []int{}
	for _, g := range groups {
		counts = append(counts, g)
	}

	switch len(groups) {
	case 1:
		return 6 // five of a kind
	case 2:
		if slices.Contains(counts, 4) {
			return 5 // four of a kind
		}
		return 4 // full house
	case 3:
		if slices.Contains(counts, 3) {
			return 3 // three of a kind
		}
		return 2 // two pair
	case 4:
		return 1 // one pair
	default:
		return 0
	}

}

// 254035588 too high
// 254078957 too high
