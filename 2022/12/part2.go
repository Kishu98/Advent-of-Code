package main

func part2(filename string) int {
	grid := parseGrid(filename)
	start, end := getPos(grid, 'S', 'E')
	grid[start.X][start.Y] = 'a'
	grid[end.X][end.Y] = 'z'

	steps := findMinSteps(
		grid,
		end,
		func(p Point, r rune) bool { return r == 'a' },
		func(r1, r2 rune) bool { return r2-r1 <= 1 },
	)

	return steps
}
