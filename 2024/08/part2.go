package main

import (
	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	grid := helpers.ParseGrid(filename, func(r rune) rune { return r })
	antennasPos := getAllAntennas(grid)
	antiNodes := getAllAntiNodes(grid, antennasPos, true)

	totalAntennas := 0
	for _, val := range antennasPos {
		totalAntennas += len(val)
	}

	return len(antiNodes) + totalAntennas
}
