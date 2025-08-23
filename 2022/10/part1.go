package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	register := 1
	totalCycles := 0
	signalSum := 0

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		if len(parts) == 1 {
			totalCycles++
			signalSum += checkCycle(register, totalCycles)
		} else {
			totalCycles++
			signalSum += checkCycle(register, totalCycles)
			totalCycles++
			signalSum += checkCycle(register, totalCycles)
			register += helpers.StrToInt(parts[1])
		}
	}); err != nil {
		log.Fatal(err)
	}

	return signalSum
}

func checkCycle(register, totalCycles int) int {
	if totalCycles == 20 || ((totalCycles-20)%40) == 0 {
		return register * totalCycles
	}

	return 0
}
