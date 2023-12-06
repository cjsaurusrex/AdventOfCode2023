package day5

import (
	"adventOfCode2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type mapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	length                int
}

type seed struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func createMapping(input string) mapping {
	s := strings.Split(input, " ")
	d, _ := strconv.Atoi(s[0])
	sr, _ := strconv.Atoi(s[1])
	l, _ := strconv.Atoi(s[2])
	return mapping{d, sr, l}
}

func getSeeds(input *string) []seed {
	i := strings.Split((*input)[7:], " ")
	s := []seed{}
	for _, se := range i {
		v, _ := strconv.Atoi(se)
		s = append(s, seed{seed: v})
	}
	return s
}

func parseInput(lines []string) map[string][]mapping {
	//var seeds []seed
	mappings := make(map[string][]mapping)

	cm := ""
	current := []mapping{}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			mappings[cm] = current
			current = []mapping{}
			cm = strings.Split(line, " ")[0]
			continue
		}
		m := createMapping(line)
		current = append(current, m)
	}
	mappings[cm] = current
	return mappings
}

func getValue(source int, mappings []mapping) int {
	for _, v := range mappings {
		if source >= v.sourceRangeStart && source <= (v.sourceRangeStart+v.length) {
			return (v.destinationRangeStart + v.length) - ((v.sourceRangeStart + v.length) - source)
		}
	}
	return source
}

func getLocation(seed seed, mappings map[string][]mapping) int {
	seed.soil = getValue(seed.seed, mappings["seed-to-soil"])
	seed.fertilizer = getValue(seed.soil, mappings["soil-to-fertilizer"])
	seed.water = getValue(seed.fertilizer, mappings["fertilizer-to-water"])
	seed.light = getValue(seed.water, mappings["water-to-light"])
	seed.temperature = getValue(seed.light, mappings["light-to-temperature"])
	seed.humidity = getValue(seed.temperature, mappings["temperature-to-humidity"])
	return getValue(seed.humidity, mappings["humidity-to-location"])
}

func getLocationFromSeed(seed int, mappings map[string][]mapping) int {
	soil := getValue(seed, mappings["seed-to-soil"])
	fertilizer := getValue(soil, mappings["soil-to-fertilizer"])
	water := getValue(fertilizer, mappings["fertilizer-to-water"])
	light := getValue(water, mappings["water-to-light"])
	temperature := getValue(light, mappings["light-to-temperature"])
	humidity := getValue(temperature, mappings["temperature-to-humidity"])
	return getValue(humidity, mappings["humidity-to-location"])
}

func getLowestLocation(seeds []seed, mappings map[string][]mapping) int {
	lowest := math.MaxInt

	for i := 0; i < len(seeds); i++ {
		seeds[i].soil = getValue(seeds[i].seed, mappings["seed-to-soil"])
		seeds[i].fertilizer = getValue(seeds[i].soil, mappings["soil-to-fertilizer"])
		seeds[i].water = getValue(seeds[i].fertilizer, mappings["fertilizer-to-water"])
		seeds[i].light = getValue(seeds[i].water, mappings["water-to-light"])
		seeds[i].temperature = getValue(seeds[i].light, mappings["light-to-temperature"])
		seeds[i].humidity = getValue(seeds[i].temperature, mappings["temperature-to-humidity"])
		seeds[i].location = getValue(seeds[i].humidity, mappings["humidity-to-location"])
		if seeds[i].location < lowest {
			lowest = seeds[i].location
		}

	}
	return lowest
}

func Part1(input string) {
	defer utils.Timing("Day5-Part1")()
	lines := utils.SplitLines(&input)
	seeds := getSeeds(&lines[0])
	mappings := parseInput(lines)

	lowest := getLowestLocation(seeds, mappings)

	fmt.Printf("Lowest: %v\n", lowest)
}

func Part2(input string) {
	defer utils.Timing("Day5-Part2")()
	lines := utils.SplitLines(&input)
	mappings := parseInput(lines)

	largestMappingVal := math.MinInt
	for _, m := range mappings {
		for _, m2 := range m {
			if m2.destinationRangeStart+m2.length > largestMappingVal {
				largestMappingVal = m2.destinationRangeStart + m2.length
			}
			if m2.sourceRangeStart+m2.length > largestMappingVal {
				largestMappingVal = m2.sourceRangeStart + m2.length
			}
		}
	}
	seedRanges := strings.Split(lines[0][7:], " ")
	lowest := math.MaxInt
	for i := 0; i < len(seedRanges); i += 2 {
		start, _ := strconv.Atoi(seedRanges[i])
		length, _ := strconv.Atoi(seedRanges[i+1])
		if start+length > largestMappingVal {
			length = (start + length) - largestMappingVal
		}

		for j := start; j < start+length; j++ {
			cl := getLocationFromSeed(j, mappings)
			if cl < lowest {
				lowest = cl
			}
		}
	}

	fmt.Printf("Lowest: %v\n", lowest)
}
