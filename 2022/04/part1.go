package main

import (
	"log"
	"slices"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	filename := "Input.txt"
	totalPairs := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, ",")
		range1, range2 := parts[0], parts[1]

		r1 := strings.Split(range1, "-")
		start1 := helpers.StrToInt(string(r1[0]))
		end1 := helpers.StrToInt(string(r1[1]))
		checkSet1 := make(map[int]bool)
		for i := start1; i <= end1; i++ {
			checkSet1[i] = true
		}

		r2 := strings.Split(range2, "-")
		if slices.Compare(r1, r2) == 0 {
			totalPairs++
			return
		}
		start2 := helpers.StrToInt(string(r2[0]))
		end2 := helpers.StrToInt(string(r2[1]))
		checkSet2 := make(map[int]bool)
		for i := start2; i <= end2; i++ {
			checkSet2[i] = true
		}

		fullyInside1 := true
		for i := start1; i <= end1; i++ {
			if !checkSet2[i] {
				fullyInside1 = false
				break
			}
		}
		if fullyInside1 {
			totalPairs++
		}

		fullyInside2 := true
		for i := start2; i <= end2; i++ {
			if !checkSet1[i] {
				fullyInside2 = false
				break
			}
		}
		if fullyInside2 {
			totalPairs++
		}
	}); err != nil {
		log.Fatal(err)
	}
	return totalPairs
}
