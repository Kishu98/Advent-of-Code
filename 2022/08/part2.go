package main

import "github.com/Kishu98/AdventOfCode/helpers"

func part2(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) int {
		return int(r - '0')
	})

	score := 0
	for i, row := range grid {
		for j := range row {
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
				continue
			}

			leftSide, rightSide, topSide, bottomSide := 0, 0, 0, 0
			// Check till top
			for currX := i - 1; currX >= 0; currX-- {
				if grid[currX][j] < grid[i][j] {
					topSide++
				} else {
					topSide++
					break
				}
			}
			// Check till bot
			for currX := i + 1; currX < len(grid); currX++ {
				if grid[currX][j] < grid[i][j] {
					bottomSide++
				} else {
					bottomSide++
					break
				}
			}
			// Check till left
			for currY := j - 1; currY >= 0; currY-- {
				if grid[i][currY] < grid[i][j] {
					leftSide++
				} else {
					leftSide++
					break
				}
			}
			// Check till right
			for currY := j + 1; currY < len(row); currY++ {
				if grid[i][currY] < grid[i][j] {
					rightSide++
				} else {
					rightSide++
					break
				}
			}
			score = max(score, leftSide*rightSide*topSide*bottomSide)
		}
	}

	return score
}
