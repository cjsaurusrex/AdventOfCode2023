package day8

import (
	"adventOfCode2023/utils"
	"fmt"
	"strings"
)

func parseInput(lines []string) ([]rune, map[string][]string) {
	d := []rune(lines[0])
	nodes := map[string][]string{}
	for _, line := range lines[2:] {
		sp := line[0:3]
		nodes[sp] = []string{line[7:10], line[12:15]}
	}
	return d, nodes
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lowestCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func Part1(input string) {
	defer utils.Timing("Day8-Part1")()
	lines := utils.SplitLines(&input)
	directions, nodes := parseInput(lines)
	totalMoves := 0
	pos := "AAA"
	for i := 0; i < len(directions); i++ {
		d := directions[i]
		if d == 'L' {
			pos = nodes[pos][0]
		} else if d == 'R' {
			pos = nodes[pos][1]
		}
		totalMoves++
		if pos == "ZZZ" {
			break
		}
		if i == len(directions)-1 {
			i = -1
			continue
		}
	}

	fmt.Printf("Total moves: %v\n", totalMoves)
}

func Part2(input string) {
	defer utils.Timing("Day8-Part2")()
	lines := utils.SplitLines(&input)
	directions, nodes := parseInput(lines)
	nodePositions := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			nodePositions = append(nodePositions, k)
		}
	}
	lcm := 0

	for i, np := range nodePositions {
		moves := 0
		for j := 0; j < len(directions); j++ {
			d := directions[j]
			if d == 'L' {
				np = nodes[np][0]
			} else if d == 'R' {
				np = nodes[np][1]
			}
			moves++
			if strings.HasSuffix(np, "Z") {
				break
			}
			if j == len(directions)-1 {
				j = -1
				continue
			}
		}
		if i == 0 {
			lcm = moves
		} else {
			lcm = lowestCommonMultiple(lcm, moves)
		}
	}

	fmt.Printf("Total moves: %v\n", lcm)
}
