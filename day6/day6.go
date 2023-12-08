package day6

import (
	"adventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func lineToSlice(input string) []int {
	res := []int{}
	split := strings.Split(input, " ")
	for _, i := range split {
		st := strings.TrimSpace(i)
		if st != "" {
			val, _ := strconv.Atoi(st)
			res = append(res, val)
		}
	}
	return res
}

func Part1(input string) {
	defer utils.Timing("Day6-Part1")()
	lines := utils.SplitLines(&input)
	time := lineToSlice(strings.Split(lines[0], ":")[1])
	distance := lineToSlice(strings.Split(lines[1], ":")[1])

	total := 1
	for i := 0; i < len(time); i++ {
		t := time[i]
		d := distance[i]
		count := 0
		for j := 0; j < t; j++ {
			if (t-j)*j > d {
				count++
			}
		}
		total = total * count
	}

	fmt.Printf("Result: %v\n", total)
}

func Part2(input string) {
	defer utils.Timing("Day6-Part2")()
	lines := utils.SplitLines(&input)

	time, _ := strconv.Atoi(strings.Join(strings.Split(strings.Split(lines[0], ":")[1], " "), ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Split(strings.Split(lines[1], ":")[1], " "), ""))

	count := 0
	for j := 0; j < time; j++ {
		if (time-j)*j > distance {
			count++
		}
	}

	fmt.Printf("Result: %v\n", count)
}
