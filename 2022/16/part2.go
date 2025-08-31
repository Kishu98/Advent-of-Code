package main

func part2(filename string) int {
	getAllValves(filename)
	getDists()
	for v := range valves {
		dist[v] = bfs(v)
	}

	usefulValves := getUsefulValves(valves)

	best := 0

	allSubsets := generateSubsets(usefulValves)

	for _, subset := range allSubsets {
		complement := difference(usefulValves, subset)

		openedMe := make(map[string]bool)
		openedElephant := make(map[string]bool)

		memo = make(map[string]int)
		meScore := dfs("AA", 26, openedMe, subset)

		memo = make(map[string]int)
		elephantScore := dfs("AA", 26, openedElephant, complement)

		total := meScore + elephantScore
		if total > best {
			best = total
		}
	}

	return best
}

func difference(vals, subset []string) []string {
	m := make(map[string]bool)
	for _, v := range subset {
		m[v] = true
	}

	var res []string
	for _, v := range vals {
		if !m[v] {
			res = append(res, v)
		}
	}

	return res
}

func generateSubsets(vals []string) [][]string {
	var res [][]string
	n := len(vals)
	var helper func(int, []string)
	helper = func(i int, curr []string) {
		if i == n {
			copyCurr := append([]string{}, curr...)
			res = append(res, copyCurr)
			return
		}
		helper(i+1, curr)
		helper(i+1, append(curr, vals[i]))
	}

	helper(0, []string{})
	return res
}
