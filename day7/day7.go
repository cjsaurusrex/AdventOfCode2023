package day7

import (
	"adventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type round struct {
	cardValues []int
	bid        int
	score      int
}

var mappingsP1 = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var mappingsP2 = map[string]int{
	"J": 1,
	"T": 10,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func parseLines(input []string, mapping map[string]int, jokers bool) []round {
	r := make([]round, len(input))
	for k, i := range input {
		s := strings.Split(i, " ")
		b, _ := strconv.Atoi(s[1])
		cards := strings.Split(s[0], "")
		nc := make([]int, len(cards))
		m := map[string]int{}
		mostCommonCard := ""
		mostCommonCardCount := 0

		for _, s1 := range cards {
			nc = append(nc, getValue(s1, mapping))
			m[s1]++
			if m[s1] > mostCommonCardCount && s1 != "J" {
				mostCommonCard = s1
				mostCommonCardCount = m[s1]
			}
		}

		if jokers && mostCommonCard != "" {
			if val, ok := m["J"]; ok {
				m[mostCommonCard] += val
				delete(m, "J")
			}
		}

		r[k] = round{nc, b, getCardScore(m)}
	}
	return r
}

func getValue(input string, mapping map[string]int) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		return mapping[input]
	}
	return val
}

func getCardScore(cards map[string]int) int {
	cmLen := len(cards)

	// High card
	if cmLen == 5 {
		return 1
	}
	// One Pair
	if cmLen == 4 {
		return 2
	}
	// 5 of a kind
	if cmLen == 1 {
		return 7
	}

	for _, v := range cards {
		// 3 of kind
		if v == 3 && cmLen == 3 {
			return 4
		}
		// Full house
		if v == 3 && cmLen == 2 {
			return 5
		}
		// 4 of Kind
		if v == 4 {
			return 6
		}
	}

	return 3
}

func Part1(input string) {
	defer utils.Timing("Day7-Part1")()
	lines := utils.SplitLines(&input)
	rounds := parseLines(lines, mappingsP1, false)

	slices.SortFunc(rounds, func(a, b round) int {
		aScore := a.score
		bScore := b.score
		if aScore == bScore {
			for i := 0; i < len(a.cardValues); i++ {
				av := a.cardValues[i]
				bv := b.cardValues[i]
				if av-bv != 0 {
					return av - bv
				}
			}
		} else {
			return aScore - bScore
		}
		return 0
	})

	total := 0
	for i := 0; i < len(rounds); i++ {
		total += rounds[i].bid * (i + 1)
	}

	fmt.Printf("Total: %v\n", total)
}

func Part2(input string) {
	defer utils.Timing("Day7-Part2")()
	lines := utils.SplitLines(&input)
	rounds := parseLines(lines, mappingsP2, true)

	slices.SortFunc(rounds, func(a, b round) int {
		aScore := a.score
		bScore := b.score
		if aScore == bScore {
			for i := 0; i < len(a.cardValues); i++ {
				av := a.cardValues[i]
				bv := b.cardValues[i]
				if av-bv != 0 {
					return av - bv
				}
			}
		} else {
			return aScore - bScore
		}
		return 0
	})

	total := 0
	for i := 0; i < len(rounds); i++ {
		total += rounds[i].bid * (i + 1)
	}

	fmt.Printf("Total: %v\n", total)
}
