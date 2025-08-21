package main

func part2() int {
	filename := "TestInput.txt"
	cache := make(map[string]int)
	dirsSize := getDirs(filename, cache)
	total := 70000000
	requiredSpace := 30000000
	leftSpace := total - dirsSize["/"]
	result := dirsSize["/"]
	for _, val := range dirsSize {
		spaceAfterDeletetion := leftSpace + val
		if spaceAfterDeletetion >= requiredSpace {
			result = min(result, val)
		}
	}

	return result
}
