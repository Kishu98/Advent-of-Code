package main

import (
	"log"
	"regexp"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	result := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	if err := helpers.ProcessInput(filename, func(s string) {
		nums := re.FindAllStringSubmatch(s, -1)
		for _, num := range nums {
			n1 := helpers.StrToInt(num[1])
			n2 := helpers.StrToInt(num[2])
			result += n1 * n2
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}
