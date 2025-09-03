package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune {
		return r
	})

	result := 0
	for x, row := range grid {
		for y := range row {
			if row[y] == 'A' {
				result += checkNewXMAS(grid, x, y)
			}
		}
	}

	return result
}

func checkNewXMAS(grid [][]rune, x, y int) int {
	result := 0
	dx1, dy1, dx2, dy2 := x-1, y-1, x+1, y+1
	if dx1 >= 0 && dx2 < len(grid) && dy1 >= 0 && dy2 < len(grid) {
		if isPair(grid[dx1][dy1], grid[dx2][dy2]) && isPair(grid[dx1][dy2], grid[dx2][dy1]) {
			result++
		}
	}

	return result
}

func isPair(a, b rune) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}
