package main

import (
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

var (
	Directions = map[rune]Point{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	facingOrder = "^>v<"
)

type Point struct {
	x, y int
}

type Guard struct {
	pos    Point
	facing rune
}

func part1(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	guard := findGuard(grid)

	guardPath := tracePath(grid, guard)
	return len(guardPath)
}

func tracePath(grid [][]rune, guard Guard) map[Point]bool {
	visited := make(map[Point]bool)
	visited[guard.pos] = true
	for guard.nextMove(grid) {
		visited[guard.pos] = true
	}

	return visited
}

func (g *Guard) turnRight() {
	currentFacing := strings.IndexRune(facingOrder, g.facing)
	g.facing = rune(facingOrder[(currentFacing+1)%4])
}

func (g *Guard) nextMove(grid [][]rune) bool {
	dir := Directions[g.facing]
	dx, dy := g.pos.x+dir.x, g.pos.y+dir.y
	next := Point{dx, dy}
	if !inBounds(grid, next) {
		return false
	}
	if grid[next.x][next.y] != '#' {
		g.pos = next
	} else {
		g.turnRight()
	}
	return true
}

func inBounds(grid [][]rune, p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(grid) && p.y < len(grid[0])
}

func findGuard(grid [][]rune) Guard {
	var guard Guard
	for x, row := range grid {
		for y, col := range row {
			if strings.ContainsRune(facingOrder, col) {
				guard.pos = Point{x, y}
				guard.facing = col
			}
		}
	}

	return guard
}
