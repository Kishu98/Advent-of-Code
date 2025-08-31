package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Valve struct {
	rate    int
	tunnels []string
}

var (
	dist   = make(map[string]map[string]int)
	valves = make(map[string]Valve)
)

var lineRegex = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func part1(filename string) int {
	getAllValves(filename)
	getDists()
	usefulValves := getUsefulValves(valves)

	opened := make(map[string]bool)
	result := dfs("AA", 30, opened, usefulValves)

	return result
}

func getDists() {
	for v := range valves {
		dist[v] = bfs(v)
	}
}

func getAllValves(filename string) {
	if err := helpers.ProcessInput(filename, func(s string) {
		matches := lineRegex.FindStringSubmatch(s)
		if matches == nil {
			return
		}
		name := matches[1]
		rate := helpers.StrToInt(matches[2])
		tunnels := strings.Split(matches[3], ", ")
		valves[name] = Valve{rate, tunnels}
	}); err != nil {
		log.Fatal(err)
	}
}

func getUsefulValves(valves map[string]Valve) []string {
	var usefulValves []string
	for v, valve := range valves {
		if valve.rate > 0 {
			usefulValves = append(usefulValves, v)
		}
	}

	return usefulValves
}

var memo = make(map[string]int)

func dfs(curr string, time int, opened map[string]bool, usefulValves []string) int {
	if time <= 0 {
		return 0
	}

	key := stateKey(curr, time, opened)
	if v, ok := memo[key]; ok {
		return v
	}

	best := 0
	for _, v := range usefulValves {
		if opened[v] {
			continue
		}

		d := dist[curr][v]
		remain := time - d - 1
		if remain <= 0 {
			continue
		}

		opened[v] = true
		gain := valves[v].rate * remain
		total := gain + dfs(v, remain, opened, usefulValves)
		if total > best {
			best = total
		}
		opened[v] = false
	}

	memo[key] = best
	return best
}

func stateKey(curr string, time int, opened map[string]bool) string {
	var keys []string
	for v, ok := range opened {
		if ok {
			keys = append(keys, v)
		}
	}
	sort.Strings(keys)
	return fmt.Sprintf("%s|%d|%s", curr, time, strings.Join(keys, ","))
}

func bfs(v string) map[string]int {
	queue := []string{v}
	visited := map[string]int{v: 0}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nxt := range valves[curr].tunnels {
			if _, ok := visited[nxt]; !ok {
				visited[nxt] = visited[curr] + 1
				queue = append(queue, nxt)
			}
		}
	}

	return visited
}
