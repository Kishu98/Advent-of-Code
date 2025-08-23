package main

import (
	"log"
	"slices"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

var dirs = map[string][]int{
	"R": {0, 1},
	"L": {0, -1},
	"D": {-1, 0},
	"U": {1, 0},
}

func part1(filename string) int {
	hpos := []int{0, 0}
	tpos := []int{0, 0}
	result := make(map[any]bool)
	result[[2]int{0, 0}] = true

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		move := parts[0]
		steps := helpers.StrToInt(parts[1])
		for range steps {
			currX, currY := dirs[move][0], dirs[move][1]
			lastPos := []int{hpos[0], hpos[1]}
			hpos = []int{hpos[0] + currX, hpos[1] + currY}
			dirst := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {1, 1}, {1, -1}, {-1, -1}, {0, 0}}
			move := true
			for _, dir := range dirst {
				currtX, currtY := dir[0], dir[1]
				tmp := []int{tpos[0] + currtX, tpos[1] + currtY}
				if slices.Compare(tmp, hpos) == 0 {
					move = false
					break
				}
			}
			if move {
				tpos = lastPos
				if !result[[2]int{tpos[0], tpos[1]}] {
					result[[2]int{tpos[0], tpos[1]}] = true
				}
			}
		}
	}); err != nil {
		log.Fatal(err)
	}

	return len(result)
}
