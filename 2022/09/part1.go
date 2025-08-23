package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Position struct {
	X, Y int
}

var dirs = map[string][]int{
	"R": {0, 1},
	"L": {0, -1},
	"D": {-1, 0},
	"U": {1, 0},
}

func part1(filename string) int {
	result := make(map[Position]bool)
	head := Position{0, 0}
	tail := Position{0, 0}

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		move := parts[0]
		steps := helpers.StrToInt(parts[1])

		for range steps {
			head = moveHead(head, move)
			tail = moveTail(head, tail)
			result[tail] = true
		}
	}); err != nil {
		log.Fatal(err)
	}

	return len(result)
}

func moveHead(head Position, move string) Position {
	head.X += dirs[move][0]
	head.Y += dirs[move][1]

	return head
}

func moveTail(head, tail Position) Position {
	dx := head.X - tail.X
	dy := head.Y - tail.Y

	if abs(dx) > 1 || abs(dy) > 1 {
		if dx != 0 {
			tail.X += dx / abs(dx)
		}
		if dy != 0 {
			tail.Y += dy / abs(dy)
		}
	}
	return tail
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
