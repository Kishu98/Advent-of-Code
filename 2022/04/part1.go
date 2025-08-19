package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	filename := "Input.txt"
	totalPairs := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		start1, end1, start2, end2 := parseRanges(s)
		if (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1) {
			totalPairs++
		}
	}); err != nil {
		log.Fatal(err)
	}
	return totalPairs
}

func parseRanges(s string) (int, int, int, int) {
	parts := strings.Split(s, ",")
	range1, range2 := parts[0], parts[1]

	r1 := strings.Split(range1, "-")
	start1 := helpers.StrToInt(r1[0])
	end1 := helpers.StrToInt(r1[1])

	r2 := strings.Split(range2, "-")
	start2 := helpers.StrToInt(r2[0])
	end2 := helpers.StrToInt(r2[1])

	return start1, end1, start2, end2
}
