package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Fields(s)
		var levels []int
		for i := range parts {
			levels = append(levels, helpers.StrToInt(parts[i]))
		}

		if checkDecreasingWithDampner(levels) || checkIncreasingWithDampner(levels) {
			result++
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkDecreasingWithDampner(levels []int) bool {
	if !checkDecreasing(levels) {
		for i := range levels {
			newLevel := make([]int, 0, len(levels)-1)
			newLevel = append(newLevel, levels[:i]...)
			newLevel = append(newLevel, levels[i+1:]...)
			if checkDecreasing(newLevel) {
				return true
			}
		}
	} else {
		return true
	}

	return false
}

func checkIncreasingWithDampner(levels []int) bool {
	if !checkIncreasing(levels) {
		for i := range levels {
			newLevel := make([]int, 0, len(levels)-1)
			newLevel = append(newLevel, levels[:i]...)
			newLevel = append(newLevel, levels[i+1:]...)
			if checkIncreasing(newLevel) {
				return true
			}
		}
	} else {
		return true
	}

	return false
}
