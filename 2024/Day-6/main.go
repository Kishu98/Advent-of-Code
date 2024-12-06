package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)
	var grid [][]string

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	currPos := getPosition(grid)

	going = "up"

	for checkPos(grid, getPosition(grid), &going) {
		fmt.Println("*********************************************")
		printGrid(grid)
		fmt.Println("*********************************************")
	}

	fmt.Println("*********************************************")
	printGrid(grid)
	fmt.Println("*********************************************")

    allPos := getAllPos(grid)

	fmt.Println(len(allPos))

    grid[currPos[0]][currPos[1]] = "^"
    going = "up"
    printGrid(grid)
}

func getAllPos(grid [][]string) [][]int {
	count := [][]int{}

	for i, row := range grid {
		for j, col := range row {
			if col == "X" {
                grid[i][j] = "."
                count = append(count, []int{i, j})
			}
		}
	}

	return count
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("Going -->", going)
}

var going string
var area int

func checkPos(grid [][]string, pos []int, going *string) bool {
	row, col := &pos[0], &pos[1]
	switch *going {
	case "up":
		if *row-1 < 0 {
			grid[*row][*col] = "X"
			return false
		}
		if grid[*row-1][*col] == "#" {
			*going = "right"
		} else {
			grid[*row][*col] = "X"
			area++
			*row = *row - 1
			grid[*row][*col] = "^"
		}
	case "right":
		if *col+1 > len(grid[*row])-1 {
			grid[*row][*col] = "X"
			return false
		}
		if grid[*row][*col+1] == "#" {
			*going = "down"
		} else {
			area++
			grid[*row][*col] = "X"
			*col = *col + 1
			grid[*row][*col] = "^"
		}
	case "down":
		if *row+1 > len(grid)-1 {
			grid[*row][*col] = "X"
			return false
		}
		if grid[*row+1][*col] == "#" {
			*going = "left"
		} else {
			grid[*row][*col] = "X"
			area++
			*row = *row + 1
			grid[*row][*col] = "^"
		}
	case "left":
		if *col-1 < 0 {
			grid[*row][*col] = "X"
			return false
		}
		if grid[*row][*col-1] == "#" {
			*going = "up"
		} else {
			grid[*row][*col] = "X"
			area++
			*col = *col - 1
			grid[*row][*col] = "^"
		}
	}

	// if *row > len(grid)-1 || *row < 0 || *col > len(grid[*row])-1 || *col < 0 {
	// 	return false
	// }

	return true
}

func getPosition(grid [][]string) []int {
	for i, row := range grid {
		for j, col := range row {
			if col == "^" {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
