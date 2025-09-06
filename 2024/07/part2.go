package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		sum, nums := getSumAndNumbers(s)
		if checkSumPossible(sum, nums, 1, nums[0], true) {
			result += sum
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}
