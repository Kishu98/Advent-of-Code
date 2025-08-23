package main

import (
	"log"
	"slices"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	knotsPos := make([][]int, 10)
	for i := range 10 {
		knotsPos[i] = []int{0, 0}
	}
	result := make(map[any]bool)
	result[[2]int{0, 0}] = true

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		move := parts[0]
		steps := helpers.StrToInt(parts[1])
		for range steps {
			currX, currY := dirs[move][0], dirs[move][1]
			knotsPos[0] = []int{knotsPos[0][0] + currX, knotsPos[0][1] + currY}
			dirst := [][]int{{0, 0}, {0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
			for i := 1; i < len(knotsPos); i++ {
				move := true
				for _, dir := range dirst {
					currtX, currtY := dir[0], dir[1]
					tmp := []int{knotsPos[i][0] + currtX, knotsPos[i][1] + currtY}
					if slices.Compare(tmp, knotsPos[i-1]) == 0 {
						move = false
						break
					}
				}
				if move {
					dx := knotsPos[i-1][0] - knotsPos[i][0]
					dy := knotsPos[i-1][1] - knotsPos[i][1]
					if abs(dx) > 1 || abs(dy) > 1 {
						if dx != 0 {
							knotsPos[i][0] += dx / abs(dx)
						}
						if dy != 0 {
							knotsPos[i][1] += dy / abs(dy)
						}
					}

					result[[2]int{knotsPos[9][0], knotsPos[9][1]}] = true
				}
			}
		}
	}); err != nil {
		log.Fatal(err)
	}

	return len(result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
