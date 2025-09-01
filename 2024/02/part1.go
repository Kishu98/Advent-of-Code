package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		levels := getLevels(s)
		if checkMonotonic(levels, true) || checkMonotonic(levels, false) {
			result++
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkMonotonic(levels []int, isIncreasing bool) bool {
	for i := 1; i < len(levels); i++ {
		if abs(levels[i-1]-levels[i]) > 3 {
			return false
		}

		if isIncreasing {
			if levels[i-1] >= levels[i] {
				return false
			}
		} else {
			if levels[i-1] <= levels[i] {
				return false
			}
		}
	}

	return true
}

func getLevels(s string) []int {
	parts := strings.Fields(s)
	var levels []int
	for i := range parts {
		levels = append(levels, helpers.StrToInt(parts[i]))
	}

	return levels
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
