package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	x, y int
}

func part1(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	antennasPos := getAllAntennas(grid)
	antiNodes := getAllAntiNodes(grid, antennasPos, false)

	return len(antiNodes)
}

func getAllAntiNodes(grid [][]rune, antennasPos map[rune][]Point, extend bool) map[Point]bool {
	antiNodes := make(map[Point]bool)
	for _, val := range antennasPos {
		for i := range val {
			currP := val[i]
			for _, p := range val {
				if p == currP {
					continue
				}
				dx, dy := currP.x-p.x, currP.y-p.y
				antiNodePos := Point{currP.x + dx, currP.y + dy}
				for antiNodePos.x >= 0 && antiNodePos.y >= 0 && antiNodePos.x < len(grid) && antiNodePos.y < len(grid[0]) {
					if extend && grid[antiNodePos.x][antiNodePos.y] != '.' {
						antiNodePos = Point{antiNodePos.x + dx, antiNodePos.y + dy}
						continue
					}
					antiNodes[antiNodePos] = true
					if !extend {
						break
					}
					antiNodePos = Point{antiNodePos.x + dx, antiNodePos.y + dy}
				}
			}
		}
	}

	return antiNodes
}

func getAllAntennas(grid [][]rune) map[rune][]Point {
	antennasPos := make(map[rune][]Point)

	for x, row := range grid {
		for y, col := range row {
			if col != '.' {
				antennasPos[col] = append(antennasPos[col], Point{x, y})
			}
		}
	}

	return antennasPos
}
