package main

import (
	"log"
	"slices"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	listA, listB := getLists(filename)

	result := 0
	for i := range listA {
		result += abs(listA[i] - listB[i])
	}

	return result
}

func getLists(filename string) ([]int, []int) {
	var listA []int
	var listB []int
	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Fields(s)
		listA = append(listA, helpers.StrToInt(parts[0]))
		listB = append(listB, helpers.StrToInt(parts[1]))
	}); err != nil {
		log.Fatal(err)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	return listA, listB
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
