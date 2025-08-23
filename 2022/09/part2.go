package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := make(map[Position]bool)
	knots := make([]Position, 10)

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		move := parts[0]
		steps := helpers.StrToInt(parts[1])

		for range steps {
			knots[0] = moveHead(knots[0], move)

			for i := 1; i < len(knots); i++ {
				knots[i] = moveTail(knots[i-1], knots[i])
			}

			result[knots[9]] = true
		}
	}); err != nil {
		log.Fatal(err)
	}

	return len(result)
}
