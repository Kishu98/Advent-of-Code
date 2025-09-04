package main

import (
	"slices"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	result := 0
	ruleMap, updates := parseInput(filename)
	for _, update := range updates {
		if !checkUpdate(ruleMap, update) {
			sortUpdate(ruleMap, update)
			result += helpers.StrToInt(update[len(update)/2])
		}
	}

	return result
}

func sortUpdate(ruleMap map[string][]string, update []string) {
	before := make(map[string]map[string]bool)
	for a, afterList := range ruleMap {
		if before[a] == nil {
			before[a] = make(map[string]bool)
		}
		for _, b := range afterList {
			before[a][b] = true
		}
	}

	slices.SortFunc(update, func(a, b string) int {
		if _, ok := before[a][b]; ok {
			return -1
		}
		if _, ok := before[b][a]; ok {
			return 1
		}
		return 0
	})
}
