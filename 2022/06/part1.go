package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	filename := "Input.txt"
	return solver(filename, 4)
}

func solver(filename string, distinct int) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		for i := 0; i <= len(s)-distinct; i++ {
			window := s[i : i+distinct]
			if !checkDuplicate(window) {
				result = i + distinct
				return
			}
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkDuplicate(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	return false
}
