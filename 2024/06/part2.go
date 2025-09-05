package main

import (
	"slices"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	gridMap := buildGridMap(grid)
	startPos, startingFacing := findGuard(grid)

	result := 0
	visited := moveGuard(grid, gridMap, startPos, startingFacing)
	for pos := range visited {
		if pos == startPos {
			continue
		}
		gridMap[pos] = '#'
		if inCycle(grid, gridMap, startPos, startingFacing) {
			result++
		}
		gridMap[pos] = '.'
	}

	return result
}

func inCycle(grid [][]rune, gridMap map[Point]rune, pos Point, facing rune) bool {
	seen := make(map[[3]int]bool)

	for {
		dir := Dirs[facing]
		next := Point{pos.x + dir.x, pos.y + dir.y}

		if !inBound(next, grid) {
			return false
		}

		state := [3]int{pos.x, pos.y, slices.Index(facingOrder, facing)}
		if seen[state] {
			return true
		}
		seen[state] = true

		if gridMap[next] != '#' {
			pos = next
		} else {
			facing = turnRight(facing)
		}
	}
}
