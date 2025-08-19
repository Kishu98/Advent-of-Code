package main

func part2() string {
	move := func(stacks [][]rune, count, from, to int) {
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:]...)
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}
	return solver(move)
}
