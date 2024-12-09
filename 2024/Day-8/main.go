package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

func main() {
	antennas := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	rows := strings.Split(str, "\n")
	rows = rows[:len(rows)-1]

	var grid [][]string

	for _, row := range rows {
		grid = append(grid, strings.Split(row, ""))
	}

	antennasPos := make(map[string][]Position)

    totalAntennas := 0

	for r, row := range grid {
		for c, cell := range row {

			if strings.Contains(antennas, cell) {
                totalAntennas++
				antennasPos[cell] = append(antennasPos[cell], Position{r, c})
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println(antennasPos)

	// dirs := []Position{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	var count int

	for key := range antennasPos {
		fmt.Println(key)
		for _, pos := range antennasPos[key] {
			for _, pos2 := range antennasPos[key] {
				if pos == pos2 {
					continue
				}
				dr, dc := pos2.row-pos.row, pos2.col-pos.col
				nr, nc := pos.row-dr, pos.col-dc
				for nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
					if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
						if grid[nr][nc] == "." {
							grid[nr][nc] = "#"
							count++
						}
						// if grid[nr][nc] != key && grid[nr][nc] != "#" {
						// 	fmt.Println("-->", grid[nr][nc])
						// 	count++
						// }
					}
					nr, nc = nr-dr, nc-dc
				}
				// } else if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] == key {
				// 	continue
				// }
			}
		}
		for _, row := range grid {
			fmt.Println(row)
		}
	}

	fmt.Println(count+totalAntennas)

}
