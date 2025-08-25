package main

func part2(filename string) int {
	grid := parseGrid(filename)
	start, end := getPos(grid, "S", "E")
	grid[start.X][start.Y] = "a"
	grid[end.X][end.Y] = "z"

	aPos := getAPos(grid)

	minSteps := 9999999999999999
	for _, apos := range aPos {
		steps := findMinSteps(grid, apos, end)
		if steps != -1 {
			minSteps = min(minSteps, steps)
		}
	}
	return minSteps
}

func getAPos(grid [][]string) []Point {
	var apos []Point
	for i, row := range grid {
		for j, col := range row {
			if col == "a" {
				apos = append(apos, Point{i, j})
			}
		}
	}
	return apos
}
