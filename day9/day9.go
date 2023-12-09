package day9

import (
	"adventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInput(lines []string) [][]int {
	res := [][]int{}
	for _, line := range lines {
		current := []int{}
		split := strings.Split(line, " ")
		for _, s := range split {
			val, _ := strconv.Atoi(s)
			current = append(current, val)
		}
		res = append(res, current)
	}
	return res
}

func getDifferences(sequence []int) []int {
	res := []int{}
	for i := 1; i < len(sequence); i++ {
		res = append(res, sequence[i]-sequence[i-1])
	}
	return res
}

func isAllZeroes(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}

func Part1(input string) {
	defer utils.Timing("Day9-Part1")()
	lines := utils.SplitLines(&input)
	histories := parseInput(lines)
	result := 0
	for _, history := range histories {
		split := [][]int{history}
		for !isAllZeroes(split[len(split)-1]) {
			split = append(split, getDifferences(split[len(split)-1]))
		}

		split[len(split)-1] = append(split[len(split)-1], 0)
		for i := len(split) - 2; i >= 0; i-- {
			previous := split[i+1][len(split[i+1])-1]
			current := split[i][len(split[i])-1]
			split[i] = append(split[i], current+previous)
		}
		result += split[0][len(split[0])-1]
	}
	fmt.Printf("Result: %v\n", result)
}

func Part2(input string) {
	defer utils.Timing("Day9-Part2")()
	lines := utils.SplitLines(&input)
	histories := parseInput(lines)
	result := 0
	for _, history := range histories {
		slices.Reverse(history)
		split := [][]int{history}
		for !isAllZeroes(split[len(split)-1]) {
			split = append(split, getDifferences(split[len(split)-1]))
		}

		split[len(split)-1] = append(split[len(split)-1], 0)
		for i := len(split) - 2; i >= 0; i-- {
			previous := split[i+1][len(split[i+1])-1]
			current := split[i][len(split[i])-1]
			split[i] = append(split[i], current+previous)
		}
		result += split[0][len(split[0])-1]
	}
	fmt.Printf("Result: %v\n", result)
}
