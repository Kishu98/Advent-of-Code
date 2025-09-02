package main

import (
	"log"
	"regexp"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := 0
	flag := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do|don't)\(\)`)
	if err := helpers.ProcessInput(filename, func(s string) {
		nums := re.FindAllStringSubmatch(s, -1)
		for _, num := range nums {
			if num[0] == "don't()" {
				flag = false
				continue
			}
			if num[0] == "do()" {
				flag = true
				continue
			}
			if flag {
				n1 := helpers.StrToInt(num[1])
				n2 := helpers.StrToInt(num[2])
				result += n1 * n2
			}
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}
