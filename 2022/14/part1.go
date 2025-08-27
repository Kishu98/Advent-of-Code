package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	X, Y int
}

func part1(filename string) int {
	var result int

	grid := createGrid()
	if err := helpers.ProcessInput(filename, func(s string) {
		if len(s) == 0 {
			return
		}
		parseRocks(grid, s)
	}); err != nil {
		log.Fatal(err)
	}

	for dropSand(grid, 0, 500) {
		result++
	}

	// printThisGrid(grid)

	return result
}

func dropSand(grid [][]rune, i, j int) bool {
	for start := i; start < len(grid); start++ {
		if grid[start][j] == '#' || grid[start][j] == 'o' {
			if start == 0 || j < 0 || j >= len(grid) {
				return false
			}
			// Move Left
			if grid[start][j-1] == '.' {
				return dropSand(grid, start, j-1)
			}
			// Move Right
			if grid[start][j+1] == '.' {
				return dropSand(grid, start, j+1)
			}
			// Sand Settles
			grid[start-1][j] = 'o'
			return true
		}
	}

	return false
}

func printThisGrid(grid [][]rune) {
	start, end := getBounds(grid)
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

func getBounds(grid [][]rune) (Point, Point) {
	start := Point{X: len(grid), Y: len(grid[0])}
	end := Point{X: 0, Y: 0}

	for i := range grid {
		for j := range len(grid[i]) {
			if grid[i][j] == '#' || grid[i][j] == 'o' {
				start.X = min(i, start.X)
				start.Y = min(j, start.Y)
				end.X = max(i, end.X)
				end.Y = max(j, end.Y)
			}
		}
	}

	return start, end
}

func parseRocks(grid [][]rune, rockPath string) {
	var points []Point
	coordinates := strings.SplitSeq(rockPath, " -> ")
	for coor := range coordinates {
		parts := strings.Split(coor, ",")
		points = append(points, Point{helpers.StrToInt(parts[1]), helpers.StrToInt(parts[0])})
	}

	for i := 0; i < len(points)-1; i++ {
		if points[i].X-points[i+1].X != 0 {
			start := min(points[i].X, points[i+1].X)
			end := max(points[i].X, points[i+1].X)
			for j := start; j <= end; j++ {
				grid[j][points[i].Y] = '#'
			}
		}
		if points[i].Y-points[i+1].Y != 0 {
			start := min(points[i].Y, points[i+1].Y)
			end := max(points[i].Y, points[i+1].Y)
			for j := start; j <= end; j++ {
				grid[points[i].X][j] = '#'
			}
		}
	}
}

func createGrid() [][]rune {
	grid := make([][]rune, 1000)
	for i := range grid {
		grid[i] = make([]rune, 1000)
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}
