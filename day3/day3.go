package day3

import (
	"adventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func buildArray(lines []string) [][]string {
	var ls [][]string
	for _, line := range lines {
		ls = append(ls, strings.Split(line, ""))
	}
	return ls
}

func isSymbol(str string) bool {
	return str != "." && !utils.IsNumber(str)
}

func hasAdjacent(row int, col int, data *[][]string, check func(string, int, int) bool, returnOnFirst bool) []string {
	startRow := row - 1
	if startRow < 0 {
		startRow = row
	}
	endRow := row + 1
	if endRow > len((*data))-1 {
		endRow = row
	}
	startCol := col - 1
	if startCol < 0 {
		startCol = col
	}
	endCol := col + 1
	if endCol > len((*data)[row])-1 {
		endCol = col
	}

	var adj []string
	for r := startRow; r <= endRow; r++ {
		if len(adj) > 0 && returnOnFirst {
			break
		}
		for c := startCol; c <= endCol; c++ {
			if check((*data)[r][c], r, c) {
				adj = append(adj, (*data)[r][c])
				if returnOnFirst {
					break
				}
			}
		}
	}
	return adj
}

func getAdjacentNumbers(row int, col int, data *[][]string) []string {
	startRow := row - 1
	if startRow < 0 {
		startRow = row
	}
	endRow := row + 1
	if endRow > len((*data))-1 {
		endRow = row
	}
	startCol := col - 1
	if startCol < 0 {
		startCol = col
	}
	endCol := col + 1
	if endCol > len((*data)[row])-1 {
		endCol = col
	}

	var adj, ranges []string
	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			isNum := utils.IsNumber((*data)[r][c])
			if isNum {
				startIndex := getStartIndex(c, &(*data)[r])
				endIndex := getEndIndex(c, &(*data)[r])
				ra := strconv.Itoa(r) + ";" + strconv.Itoa(startIndex) + ":" + strconv.Itoa(endIndex)
				if !slices.Contains(ranges, ra) {
					adj = append(adj, strings.Join((*data)[r][startIndex:endIndex+1], ""))
					ranges = append(ranges, ra)
				}
			}
		}
	}
	return adj
}

func getEndIndex(startCol int, row *[]string) int {
	for i := startCol; i < len(*row); i++ {
		if !utils.IsNumber((*row)[i]) {
			return i - 1
		}
	}
	return len(*row) - 1
}

func getStartIndex(startCol int, row *[]string) int {
	for i := startCol; i >= 0; i-- {
		if !utils.IsNumber((*row)[i]) {
			return i + 1
		}
	}
	return 0
}

func Part1(input string) {
	defer utils.Timing("Day3-Part1")()
	lines := utils.SplitLines(&input)
	ls := buildArray(lines)
	total := 0

	for row := 0; row < len(ls); row++ {
		for col := 0; col < len(ls[row]); col++ {
			if utils.IsNumber(ls[row][col]) {
				endIndex := getEndIndex(col, &ls[row])
				predicate := func(input string, row int, col int) bool {
					return isSymbol(input)
				}
				startAdj := hasAdjacent(row, col, &ls, predicate, true)
				endAdj := hasAdjacent(row, endIndex, &ls, predicate, true)
				if len(startAdj) > 0 || len(endAdj) > 0 {
					res, _ := strconv.Atoi(strings.Join(ls[row][col:endIndex+1], ""))
					total += res
				}
				col = endIndex
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func Part2(input string) {
	defer utils.Timing("Day3-Part2")()
	lines := utils.SplitLines(&input)
	ls := buildArray(lines)
	total := 0

	for row := 0; row < len(ls); row++ {
		for col := 0; col < len(ls[row]); col++ {
			if ls[row][col] == "*" {
				adj := getAdjacentNumbers(row, col, &ls)
				if len(adj) == 2 {
					v1, _ := strconv.Atoi(adj[0])
					v2, _ := strconv.Atoi(adj[1])
					total += v1 * v2
				}
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}
