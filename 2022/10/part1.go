package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type CPU struct {
	cycle    int
	register int
}

func part1(filename string) int {
	cpu := CPU{cycle: 0, register: 1}
	signalSum := 0

	cycleFunc := func(cycle, register int) {
		if cycle == 20 || ((cycle-20)%40) == 0 {
			signalSum += cycle * register
		}
	}

	if err := helpers.ProcessInput(filename, func(s string) {
		cpu.runInstruction(s, cycleFunc)
	}); err != nil {
		log.Fatal(err)
	}

	return signalSum
}

func (c *CPU) runInstruction(instructions string, cycleFunc func(cycle, register int)) {
	parts := strings.Split(instructions, " ")
	if len(parts) == 1 {
		c.cycle++
		cycleFunc(c.cycle, c.register)
	} else {
		c.cycle++
		cycleFunc(c.cycle, c.register)
		c.cycle++
		cycleFunc(c.cycle, c.register)
		c.register += helpers.StrToInt(parts[1])
	}
}
