package day2

import (
	"adventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (int, []map[string]int) {
	split := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
	combinations := strings.Split(split[1], ";")
	var cr []map[string]int

	for _, combination := range combinations {
		result := make(map[string]int)
		c := strings.Split(combination, ",")
		for _, s := range c {
			sp := strings.Split(strings.TrimSpace(s), " ")
			r, _ := strconv.Atoi(sp[0])
			result[sp[1]] = r
		}
		cr = append(cr, result)
	}

	return id, cr
}

func Part1(input string) {
	defer utils.Timing("Day2-Part1")()
	lines := utils.SplitLines(&input)
	maxR, maxG, maxB := 12, 13, 14
	total := 0

	for _, line := range lines {
		id, c := parseLine(line)
		failed := false
		for _, co := range c {
			if failed {
				break
			}
			for k, v := range co {
				if k == "red" && v > maxR {
					failed = true
					break
				}
				if k == "blue" && v > maxB {
					failed = true
					break
				}
				if k == "green" && v > maxG {
					failed = true
					break
				}
			}
		}
		if !failed {
			total += id
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func Part2(input string) {
	defer utils.Timing("Day2-Part2")()
	lines := utils.SplitLines(&input)
	power := 0
	for _, line := range lines {
		r, g, b := 0, 0, 0
		_, c := parseLine(line)
		for _, m := range c {
			for k, v := range m {
				if k == "red" && v > r {
					r = v
				}
				if k == "blue" && v > b {
					b = v
				}
				if k == "green" && v > g {
					g = v
				}
			}
		}
		res := r * g * b
		power += res
	}

	fmt.Printf("Power: %v\n", power)
}
