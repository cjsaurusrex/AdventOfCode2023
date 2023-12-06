package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(string(data), "\n")
}

func SplitLines(input *string) []string {
	return strings.Split(*input, "\n")
}

func Timing(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s: %v\n", name, time.Since(start))
	}
}

func IsNumber(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
