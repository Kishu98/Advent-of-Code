package main

import (
	"slices"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	x, y int
}

var Dirs = map[rune]Point{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var facingOrder = []rune{'^', '>', 'v', '<'}

func part1(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	gridMap := buildGridMap(grid)
	guardPos, guardFacing := findGuard(grid)

	return len(moveGuard(grid, gridMap, guardPos, guardFacing))
}

func moveGuard(grid [][]rune, gridMap map[Point]rune, pos Point, facing rune) map[Point]bool {
	visited := make(map[Point]bool)
	visited[pos] = true

	for {
		dir := Dirs[facing]
		next := Point{pos.x + dir.x, pos.y + dir.y}
		if !inBound(next, grid) {
			break
		}
		if gridMap[next] != '#' {
			pos = next
			visited[pos] = true
		} else {
			facing = turnRight(facing)
		}
	}

	return visited
}

func inBound(guardPos Point, grid [][]rune) bool {
	return (guardPos.x >= 0 && guardPos.y >= 0 && guardPos.x < len(grid) && guardPos.y < len(grid[0]))
}

func turnRight(f rune) rune {
	return facingOrder[(slices.Index(facingOrder, f)+1)%len(facingOrder)]
}

func findGuard(grid [][]rune) (Point, rune) {
	for x, row := range grid {
		for y, col := range row {
			if strings.ContainsRune("V<>^", col) {
				return Point{x, y}, col
			}
		}
	}

	return Point{}, '^'
}

func buildGridMap(grid [][]rune) map[Point]rune {
	gridMap := make(map[Point]rune)
	for x, row := range grid {
		for y, col := range row {
			gridMap[Point{x, y}] = col
		}
	}

	return gridMap
}
