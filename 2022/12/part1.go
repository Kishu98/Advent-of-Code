package main

import (
	"log"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	X, Y int
}

func part1(filename string) int {
	grid := parseGrid(filename)
	start, end := getPos(grid, 'S', 'E')
	grid[start.X][start.Y] = 'a'
	grid[end.X][end.Y] = 'z'

	steps := findMinSteps(
		grid,
		start,
		func(p Point, val rune) bool { return p == end },
		func(r1, r2 rune) bool { return r1-r2 <= 1 },
	)

	return steps
}

func findMinSteps(grid [][]rune, start Point, goalFunc func(Point, rune) bool, canStep func(rune, rune) bool) int {
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[Point]bool)
	queue := []struct {
		p     Point
		steps int
	}{{start, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if goalFunc(curr.p, grid[curr.p.X][curr.p.Y]) {
			return curr.steps
		}

		for _, dir := range dirs {
			dx, dy := curr.p.X+dir[0], curr.p.Y+dir[1]
			if dx < 0 || dy < 0 || dx >= len(grid) || dy >= len(grid[0]) {
				continue
			}

			next := Point{dx, dy}
			if visited[next] {
				continue
			}

			if canStep(grid[dx][dy], grid[curr.p.X][curr.p.Y]) {
				visited[next] = true
				queue = append(queue, struct {
					p     Point
					steps int
				}{
					next,
					curr.steps + 1,
				})
			}
		}
	}

	return -1
}

func parseGrid(filename string) [][]rune {
	var grid [][]rune
	if err := helpers.ProcessInput(filename, func(s string) {
		grid = append(grid, []rune(s))
	}); err != nil {
		log.Fatal(err)
	}

	return grid
}

func getPos(grid [][]rune, start rune, end rune) (Point, Point) {
	var s, e Point
	for i, row := range grid {
		for j, col := range row {
			if col == start {
				s.X, s.Y = i, j
			}
			if col == end {
				e.X, e.Y = i, j
			}
		}
	}
	return s, e
}
