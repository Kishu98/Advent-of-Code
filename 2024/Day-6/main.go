package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	row, col int
}

func main() {
	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	directions := map[rune]Position{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	turnRight := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	simulate := func(grid [][]rune, startPos Position, startDir rune) (map[Position]bool, bool) {
		currentPos := startPos
		currentDir := startDir
		visitedStates := make(map[struct {
			pos Position
			dir rune
		}]bool)
		visitedPositions := make(map[Position]bool)
		rows, cols := len(grid), len(grid[0])

		for {
			state := struct {
				pos Position
				dir rune
			}{currentPos, currentDir}
			if visitedStates[state] {
				return visitedPositions, true
			}
			visitedStates[state] = true
			visitedPositions[currentPos] = true

			dr, dc := directions[currentDir].row, directions[currentDir].col
			nextPos := Position{currentPos.row + dr, currentPos.col + dc}

			if nextPos.row < 0 || nextPos.row >= rows || nextPos.col < 0 || nextPos.col >= cols {
				return visitedPositions, false
			}

			if grid[nextPos.row][nextPos.col] == '#' {
				currentDir = turnRight[currentDir]
			} else {
				currentPos = nextPos
			}
		}
	}

	var startPos Position
	var startDir rune
	for r, row := range grid {
		for c, cell := range row {
			if _, ok := directions[cell]; ok {
				startPos = Position{r, c}
				startDir = cell
				break
			}
		}
	}

	originalVisited, _ := simulate(grid, startPos, startDir)

	loopPositions := make(map[Position]bool)
	for r, row := range grid {
		for c, cell := range row {
			pos := Position{r, c}
			if _, visited := originalVisited[pos]; visited && cell == '.' {
				grid[r][c] = '#'
				_, isLoop := simulate(grid, startPos, startDir)
				if isLoop {
					printGrid(grid)
					loopPositions[pos] = true
				}
				grid[r][c] = '.'
			}
		}
	}

	fmt.Println(len(loopPositions))
}

func printGrid(grid [][]rune) {
	fmt.Println("*********************************************")
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

