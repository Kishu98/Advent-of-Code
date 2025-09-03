package main

import "github.com/Kishu98/AdventOfCode/helpers"

func part1(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) int {
		return int(r - '0')
	})

	totalVisible := len(grid) * len(grid[0])
	for i, row := range grid {
		for j := range row {
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
				continue
			}

			visibleTop := true
			visibleBottom := true
			visibleLeft := true
			visibleRight := true

			// Check till top
			for currX := i - 1; currX >= 0; currX-- {
				if grid[currX][j] >= grid[i][j] {
					visibleTop = false
				}
			}
			// Check till bot
			for currX := i + 1; currX < len(grid); currX++ {
				if grid[currX][j] >= grid[i][j] {
					visibleBottom = false
				}
			}
			// Check till left
			for currY := j - 1; currY >= 0; currY-- {
				if grid[i][currY] >= grid[i][j] {
					visibleLeft = false
				}
			}
			// Check till right
			for currY := j + 1; currY < len(row); currY++ {
				if grid[i][currY] >= grid[i][j] {
					visibleRight = false
				}
			}

			if !visibleLeft && !visibleBottom && !visibleRight && !visibleTop {
				totalVisible--
			}
		}
	}

	return totalVisible
}
