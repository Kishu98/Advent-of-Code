package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	cpu := CPU{cycle: 0, register: 1}
	var crtOutput strings.Builder

	cycleFunc := func(cycle, register int) {
		pixelPos := (cycle - 1) % 40
		if pixelPos >= register-1 && pixelPos <= register+1 {
			crtOutput.WriteString("#")
		} else {
			crtOutput.WriteString(".")
		}
		if cycle%40 == 0 {
			crtOutput.WriteString("\n")
		}
	}

	if err := helpers.ProcessInput(filename, func(s string) {
		cpu.runInstruction(s, cycleFunc)
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(crtOutput.String())

	return 0
}
