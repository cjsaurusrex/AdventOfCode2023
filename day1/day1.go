package day1

import (
	"adventOfCode2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var maps = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func Part1(input *string) {
	defer utils.Timing("Day1-Part1")()
	lines := utils.SplitLines(input)

	var total int

	for _, line := range lines {
		var first, last string
		for _, char := range line {
			if unicode.IsNumber(char) {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			}
		}
		combined, _ := strconv.Atoi(first + last)
		total += combined
	}

	fmt.Printf("Total: %v\n", total)
}

func Part2(input *string) {
	defer utils.Timing("Day1-Part2")()
	lines := utils.SplitLines(input)
	var total int

	for _, line := range lines {
		var first, last string
		highest := math.MaxInt32
		lowest := math.MinInt32
		for s, s1 := range maps {
			in := strings.Index(line, s)
			in2 := strings.LastIndex(line, s)
			if in != -1 && in < highest {
				highest = in
				first = s1
			}
			if in2 != -1 && in2 > lowest {
				lowest = in2
				last = s1
			}
		}

		combined, _ := strconv.Atoi(first + last)
		total += combined
	}

	fmt.Printf("Total: %v\n", total)
}
