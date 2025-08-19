package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2() int {
	filename := "Input.txt"
	totalPairs := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		start1, end1, start2, end2 := parseRanges(s)
		if start1 <= end2 && start2 <= end1 {
			totalPairs++
		}
	}); err != nil {
		log.Fatal(err)
	}
	return totalPairs
}
