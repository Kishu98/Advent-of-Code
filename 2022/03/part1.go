package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	totalPriority := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		l := len(s) / 2
		compartmentA := s[:l]
		compartmentB := s[l:]

		var charSet = make(map[rune]bool)
		for _, c := range compartmentA {
			charSet[c] = true
		}

		for _, c := range compartmentB {
			if charSet[c] {
				totalPriority += helpers.MapAlphatoNum(c)
				break
			}
		}
	}); err != nil {
		log.Fatal(err)
	}
	return totalPriority
}
