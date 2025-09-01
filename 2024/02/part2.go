package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		levels := getLevels(s)

		if checkMonotonicWithDampner(levels) {
			result++
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkMonotonicWithDampner(levels []int) bool {
	if checkMonotonic(levels, false) || checkMonotonic(levels, true) {
		return true
	}

	for i := range levels {
		newLevel := make([]int, 0, len(levels)-1)
		newLevel = append(newLevel, levels[:i]...)
		newLevel = append(newLevel, levels[i+1:]...)
		if checkMonotonic(newLevel, false) || checkMonotonic(newLevel, true) {
			return true
		}
	}

	return false
}
