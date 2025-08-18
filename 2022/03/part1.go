package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	totalPriority := 0
	helpers.ProcessInput(filename, func(s string) {
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
	})
	return totalPriority
}
