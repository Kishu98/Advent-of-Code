package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() string {
	filename := "Input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(file), "\n\n")
	stacksInput := strings.Split(parts[0], "\n")
	moveInput := strings.Split(parts[1], "\n")

	totalStacks := (len(stacksInput[len(stacksInput)-1]) + 1) / 4
	stacks := make([][]rune, totalStacks)

	for i := len(stacksInput) - 2; i >= 0; i-- {
		line := stacksInput[i]
		for s := range totalStacks {
			pos := 1 + s*4
			if line[pos] != ' ' {
				stacks[s] = append(stacks[s], rune(line[pos]))
			}
		}
	}

	for _, line := range moveInput {
		if strings.TrimSpace(line) == "" {
			continue
		}
		count, from, to := 0, 0, 0
		fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)

		from--
		to--

		for i := 0; i < count; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	result := ""
	for i := range totalStacks {
		top := len(stacks[i]) - 1
		result += string(stacks[i][top])
	}

	return result
}
