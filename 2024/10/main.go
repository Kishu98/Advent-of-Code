package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Row int
	Col int
}

func main() {
	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var arr []int
		for _, l := range line {
			arr = append(arr, str_to_int(l))
		}
		grid = append(grid, arr)
	}

	trailStartPos := findAllTrailHeads(grid)

	var score int
	for _, trailhead := range trailStartPos {
		score += checkIfNice(grid, trailhead)
	}

	fmt.Println(score)
}

func checkIfNice(grid [][]int, trailhead Position) int {
	rows, cols := len(grid), len(grid[0])
	directions := []Position{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	visited := make(map[Position]bool)
	queue := []Position{trailhead}
	reachableNines := make(map[Position]bool)
	var count int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			count++
			continue
		}

		visited[node] = true

		r, c := node.Row, node.Col
		if grid[r][c] == 9 {
			reachableNines[node] = true
			count++
		}

		for _, dir := range directions {
			nr, nc := r+dir.Row, c+dir.Col
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if grid[r][c]+1 == grid[nr][nc] && !visited[Position{nr, nc}] {
					queue = append(queue, Position{nr, nc})
				}
			}
		}
	}

	return count
}

func findAllTrailHeads(grid [][]int) []Position {
	var pos []Position
	for r, row := range grid {
		for c, cell := range row {
			if cell == 0 {
				pos = append(pos, Position{r, c})
			}
		}
	}

	return pos
}

func str_to_int(s string) int {
	istr, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return istr
}
