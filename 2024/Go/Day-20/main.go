package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	Y int
	X int
}

func main() {
	file, err := os.OpenFile("Input.txt", os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var l []rune
		for _, c := range line {
			l = append(l, c)
		}
		grid = append(grid, l)
	}

	var start Point

	for r, row := range grid {
		for c, col := range row {
			if col == 'S' {
				start = Point{r, c}
			} else {
				continue
			}
		}
	}

	dist := make([][]int, len(grid))

	for i := range len(grid) {
		dist[i] = make([]int, len(grid[0]))
		for j := range len(grid[0]) {
			dist[i][j] = -1
		}
	}

	rows := len(grid)
	cols := len(grid[0])

	dist[start.Y][start.X] = 0

	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	r := start.Y
	c := start.X
	for grid[r][c] != 'E' {
		cr, cc := r, c
		for _, dir := range dirs {
			nr, nc := cr+dir.Y, cc+dir.X
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if grid[nr][nc] == '#' {
				continue
			}
			if dist[nr][nc] != -1 {
				continue
			}
			dist[nr][nc] = dist[cr][cc] + 1
			r = nr
			c = nc
		}
	}

	count := 0
    cheats := make(map[int]int)

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '#' {
				continue
			}
			for _, p := range []Point{{r + 2, c}, {r + 1, c + 1}, {r, c + 2}, {r - 1, c + 1}} {
				nr, nc := p.Y, p.X
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if grid[nr][nc] == '#' {
					continue
				}
                diff := abs(dist[r][c] - dist[nr][nc])
				if diff >= 102 {
                    cheats[diff] += 1
					count++
				}
			}
		}
	}

	fmt.Println(count)
    // for k := range cheats {
    //     fmt.Println(k, "-->", cheats[k])
    // }
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
