package main

import (
	"bufio"
	"log"
	"os"
)

var scores = map[string]int{
	"A X": 4, "A Y": 8, "A Z": 3,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 7, "C Y": 2, "C Z": 6,
}

func part1() int {
	totalPoints := 0
	processInput(func(line string) {
		points := calculatePoints(line)
		totalPoints += points
	})

	return totalPoints
}

func calculatePoints(line string) int {
	return scores[line]
}

func processInput(lineProcessor func(string)) {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineProcessor(scanner.Text())
	}
}
