package main

import (
	"bufio"
	"log"
	"os"
)

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	result := 0
	for x, row := range grid {
		for y := range row {
			if row[y] == 'A' {
				result += checkNewXMAS(grid, x, y)
			}
		}
	}

	return result
}

func checkNewXMAS(grid [][]rune, x, y int) int {
	result := 0
	dx1, dy1, dx2, dy2 := x-1, y-1, x+1, y+1
	if dx1 >= 0 && dy1 >= 0 && dx2 >= 0 && dy2 >= 0 && dx1 < len(grid) && dy2 < len(grid) && dx2 < len(grid) && dy2 < len(grid) {
		if grid[dx1][dy1] == 'M' && grid[dx2][dy2] == 'S' {
			if grid[dx1][dy2] == 'M' && grid[dx2][dy1] == 'S' {
				result++
			} else if grid[dx1][dy2] == 'S' && grid[dx2][dy1] == 'M' {
				result++
			}
		} else if grid[dx1][dy1] == 'S' && grid[dx2][dy2] == 'M' {
			if grid[dx1][dy2] == 'M' && grid[dx2][dy1] == 'S' {
				result++
			} else if grid[dx1][dy2] == 'S' && grid[dx2][dy1] == 'M' {
				result++
			}
		}
	}

	return result
}
