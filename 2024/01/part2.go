package main

func part2(filename string) int {
	result := 0
	listA, listB := getLists(filename)

	freq := make(map[int]int)
	for _, elemB := range listB {
		freq[elemB]++
	}

	for _, elemA := range listA {
		if similar, found := freq[elemA]; found {
			result += elemA * similar
		}
	}

	return result
}
