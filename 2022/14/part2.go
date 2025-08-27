package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	var result int

	grid := createGrid()
	if err := helpers.ProcessInput(filename, func(s string) {
		if len(s) == 0 {
			return
		}
		parseRocks(grid, s)
	}); err != nil {
		log.Fatal(err)
	}

	_, end := getBounds(grid)
	limit := end.X + 2

	for i := range grid[limit] {
		grid[limit][i] = '#'
	}

	for dropSand(grid, 0, 500) {
		result++
	}

	// printThisGrid(grid)

	return result
}
