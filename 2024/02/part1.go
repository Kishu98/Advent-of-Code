package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Fields(s)
		var levels []int
		for i := range parts {
			levels = append(levels, helpers.StrToInt(parts[i]))
		}

		if checkDecreasing(levels) || checkIncreasing(levels) {
			result++
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkDecreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i-1] <= levels[i] || abs(levels[i-1]-levels[i]) > 3 {
			return false
		}
	}
	return true
}

func checkIncreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i-1] >= levels[i] || abs(levels[i-1]-levels[i]) > 3 {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
