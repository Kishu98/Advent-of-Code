package main

func part2(filename string) int {
	result := 0
	listA, listB := getLists(filename)

	for _, elemA := range listA {
		similar := 0
		for _, elemB := range listB {
			if elemA == elemB {
				similar++
			}
		}
		result += elemA * similar
	}

	return result
}
