package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)

	var grid [][]string

	g := strings.Split(str, "\n")
	g = g[:len(g)-1]

	for i := range len(g) {
		grid = append(grid, append([]string{}, strings.Split(g[i], "")...))
	}

	word := "XMAS"

	// result := countXMAS(grid, word)
	fmt.Println(countXMAS(grid, word))
	fmt.Println(countXMASPattern(grid))

}

func countXMAS(grid [][]string, w string) int {
	rows, cols := len(grid), len(grid[0])
	wlen := len(w)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	isValid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	findWord := func(x, y, dx, dy int) bool {
		for i := 0; i < wlen; i++ {
			nx, ny := x+i*dx, y+i*dy
			if !isValid(nx, ny) || grid[nx][ny] != string(w[i]) {
				return false
			}
		}
		return true
	}

	count := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if findWord(x, y, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}

func countXMASPattern(grid [][]string) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	isValidXmas := func(x, y int) bool {
		if x-1 < 0 || x+1 >= rows || y-1 < 0 || y+1 >= cols {
			return false
		}

		tLeft, tRight := grid[x-1][y-1], grid[x-1][y+1]
		center, bLeft, bRight := grid[x][y], grid[x+1][y-1], grid[x+1][y+1]

		if center != "A" {
			return false
		}

		diag1 := tLeft + center + bRight
		diag2 := tRight + center + bLeft

		return (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM")
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if isValidXmas(x, y) {
				count++
			}
		}
	}

	return count
}
