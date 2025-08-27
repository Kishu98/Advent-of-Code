package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part2(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Read Error:", err)
	}

	grid := make([][]rune, 1000)
	for i := range grid {
		grid[i] = make([]rune, 1000)
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	rockPaths := strings.SplitSeq(string(file), "\n")
	var result int
	for rockPath := range rockPaths {
		if len(rockPath) == 0 {
			continue
		}
		parseRocks(grid, rockPath)
	}

	limit := getEnd(grid) + 2
	fmt.Println(limit)

	for i := range grid[limit] {
		grid[limit][i] = '#'
	}

	for dropSand(grid, 0, 500) {
		result++
	}
	// dropSand(grid, 0, 500)
	// printThisGrid(grid)
	return result
}

func getEnd(grid [][]rune) int {
	var start Point
	var end Point
	got := false
	for i := range grid {
		for j := range len(grid[i]) {
			if (grid[i][j] == 'o' || grid[i][j] == '#') && !got {
				start.X = i
				start.Y = j
				end.X = i
				end.Y = j
				got = true
			}
			if grid[i][j] == '#' || grid[i][j] == 'o' {
				start.X = min(i, start.X)
				start.Y = min(j, start.Y)
				end.X = max(i, end.X)
				end.Y = max(j, end.Y)
			}
		}
	}

	return end.X
}
