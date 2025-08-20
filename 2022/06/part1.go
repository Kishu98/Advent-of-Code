package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	filename := "Input.txt"
	return solver(filename, 4)
}

func solver(filename string, distinct int) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		start := 0
		end := distinct - 2
		for i := start; i <= end; i++ {
			if !checkDuplicate(s[start : end+2]) {
				result = end + 2
				return
			}
			start++
			end++
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkDuplicate(s string) bool {
	for i, c := range s {
		if strings.Contains(s[:i]+s[i+1:], string(c)) {
			return true
		}
	}
	return false
}
