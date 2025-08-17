package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

var scores = map[string]int{
	"A X": 4, "A Y": 8, "A Z": 3,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 7, "C Y": 2, "C Z": 6,
}

func part1() int {
	totalPoints := 0
	helpers.ProcessInput(func(line string) {
		points := calculatePoints(line)
		totalPoints += points
	})

	return totalPoints
}

func calculatePoints(line string) int {
	return scores[line]
}
