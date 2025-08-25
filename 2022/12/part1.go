package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	X, Y int
}

func part1(filename string) int {
	grid := parseGrid(filename)
	start, end := getPos(grid, "S", "E")
	grid[start.X][start.Y] = "a"
	grid[end.X][end.Y] = "z"

	steps := findMinSteps(grid, start, end)
	return steps
}

func findMinSteps(grid [][]string, start, end Point) int {
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[Point]bool)
	queue := []struct {
		p     Point
		steps int
	}{{start, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.p == end {
			return curr.steps
		}

		if visited[curr.p] {
			continue
		}

		visited[curr.p] = true
		for _, dir := range dirs {
			dx, dy := curr.p.X+dir[0], curr.p.Y+dir[1]
			if dx < 0 || dy < 0 || dx >= len(grid) || dy >= len(grid[0]) {
				continue
			}

			if []rune(grid[dx][dy])[0]-[]rune(grid[curr.p.X][curr.p.Y])[0] <= 1 {
				queue = append(queue, struct {
					p     Point
					steps int
				}{
					Point{dx, dy},
					curr.steps + 1,
				})
			}
		}
	}

	return -1
}

func parseGrid(filename string) [][]string {
	var grid [][]string
	if err := helpers.ProcessInput(filename, func(s string) {
		grid = append(grid, strings.Split(s, ""))
	}); err != nil {
		log.Fatal(err)
	}

	return grid
}

func getPos(grid [][]string, start string, end string) (Point, Point) {
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
