package main

import (
	"adventOfCode2023/day8"
	_ "embed"
)

//go:embed day1/input.txt
var day1Input string

//go:embed day2/input.txt
var day2Input string

//go:embed day3/input.txt
var day3Input string

//go:embed day4/input.txt
var day4Input string

//go:embed day5/input.txt
var day5Input string

//go:embed day6/input.txt
var day6Input string

//go:embed day7/input.txt
var day7Input string

//go:embed day8/input.txt
var day8Input string

func main() {
	//day1.Part1(&day1Input)
	//day1.Part2(&day1Input)
	//day2.Part1(day2Input)
	//day2.Part2(day2Input)
	//day3.Part1(day3Input)
	//day3.Part2(day3Input)
	//day4.Part1(day4Input)
	//day4.Part2(day4Input)
	//day5.Part1(day5Input)
	//day5.Part2(day5Input)
	//day6.Part1(day6Input)
	//day6.Part2(day6Input)
	//day7.Part1(day7Input)
	//day7.Part2(day7Input)
	day8.Part1(day8Input)
	day8.Part2(day8Input)
}
