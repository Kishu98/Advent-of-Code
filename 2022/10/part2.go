package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

var pixelRow = ""

func part2(filename string) int {
	register := 1
	cycle := 0

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")
		if len(parts) == 1 {
			cycle++
			pixelRow += checkCRT(register, cycle)
		} else {
			cycle++
			pixelRow += checkCRT(register, cycle)
			cycle++
			pixelRow += checkCRT(register, cycle)
			register += helpers.StrToInt(parts[1])
		}
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pixelRow)

	return 0
}

func checkCRT(register int, cycle int) string {
	c := cycle - 1
	c = c % 40
	if len(pixelRow) == 40 {
		fmt.Println(pixelRow)
		pixelRow = ""
	}
	if c == register-1 || c == register || c == register+1 {
		return "#"
	}
	return "."
}
