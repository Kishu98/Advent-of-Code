package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2() int {
	var group []string
	totalPriority := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		group = append(group, s)
		if len(group) > 2 {
			var charSet = make(map[rune]bool)
			for _, c := range group[0] {
				charSet[c] = true
			}

			var commonSet = make(map[rune]bool)
			for _, c := range group[1] {
				if charSet[c] {
					commonSet[c] = true
				}
			}

			for _, c := range group[2] {
				if commonSet[c] {
					totalPriority += helpers.MapAlphatoNum(c)
					group = group[:0]
					break
				}
			}
		}
	}); err != nil {
		log.Fatal(err)
	}
	return totalPriority
}
