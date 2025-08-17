package main

import "github.com/Kishu98/AdventOfCode/helpers"

var myHand = map[string]string{
	"A X": "Z", "A Y": "X", "A Z": "Y",
	"B X": "X", "B Y": "Y", "B Z": "Z",
	"C X": "Y", "C Y": "Z", "C Z": "X",
}

func part2() int {
	totalPoints := 0
	helpers.ProcessInput(func(line string) {
		elfHand := string(line[0])
		mySuggestedHand := calculateHand(line)
		points := calculatePoints(elfHand + " " + mySuggestedHand)
		totalPoints += points
	})
	return totalPoints
}

func calculateHand(line string) string {
	return myHand[line]
}
