package day4

import (
	"adventOfCode2023/utils"
	"fmt"
	"slices"
	"strings"
)

func parseCard(input string) (card []string, r []string) {
	sc := strings.Split(input, "|")
	result := strings.Split(strings.TrimSpace(strings.Split(sc[0], ":")[1]), " ")
	c := strings.Split(sc[1], " ")
	return c, result
}

func Part1(input string) {
	defer utils.Timing("Day4-Part1")()
	lines := utils.SplitLines(&input)
	total := 0

	for _, line := range lines {
		card, res := parseCard(line)
		points := 0
		for _, s := range card {
			if strings.TrimSpace(s) == "" {
				continue
			}
			if slices.Contains(res, s) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
		total += points
	}

	fmt.Printf("Total %v\n", total)
}

func Part2(input string) {
	defer utils.Timing("Day4-Part2")()
	lines := utils.SplitLines(&input)
	total := 0
	cards := make(map[int]int)

	for n, line := range lines {
		card, res := parseCard(line)
		if val, ok := cards[n]; !ok {
			cards[n] = 1
		} else {
			cards[n] = val + 1
		}
		winnings := 0
		for _, s := range card {
			if strings.TrimSpace(s) == "" {
				continue
			}
			if slices.Contains(res, s) {
				winnings++
			}
		}
		for i := 0; i < winnings; i++ {
			if val, ok := cards[n+i+1]; !ok {
				cards[n+i+1] = cards[n] * 1
			} else {
				cards[n+i+1] = (cards[n] * 1) + val
			}
		}
	}

	for k, i := range cards {
		if k <= len(lines) {
			total += i
		}
	}

	fmt.Printf("Total: %v\n", total)
}
