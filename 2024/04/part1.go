package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune {
		return r
	})

	result := 0
	for x, row := range grid {
		for y := range row {
			if row[y] == 'X' {
				result += checkXMAS(grid, x, y)
			}
		}
	}

	return result
}

func checkXMAS(grid [][]rune, x, y int) int {
	dirs := [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	var result int
	for _, dir := range dirs {
		word := ""
		for k := range 4 {
			dx, dy := x+(k*dir[0]), y+(k*dir[1])
			if dx >= 0 && dx < len(grid) && dy >= 0 && dy < len(grid[dx]) {
				word += string(grid[dx][dy])
			} else {
				break
			}
		}

		if word == "XMAS" {
			result++
		}
	}

	return result
}
