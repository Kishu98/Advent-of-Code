package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	guard := findGuard(grid)
	guardPath := tracePath(grid, guard)

	result := 0
	for path := range guardPath {
		grid[path.x][path.y] = '#'
		if checkCycle(grid, guard) {
			result++
		}
		grid[path.x][path.y] = '.'
	}

	return result
}

func checkCycle(grid [][]rune, guard Guard) bool {
	seen := make(map[Guard]bool)
	for guard.nextMove(grid) {
		state := Guard{guard.pos, guard.facing}
		if seen[state] {
			return true
		}
		seen[state] = true
	}

	return false
}
