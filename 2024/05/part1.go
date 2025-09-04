package main

import (
	"log"
	"os"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	result := 0
	ruleMap, updates := parseInput(filename)

	for _, update := range updates {
		if checkUpdate(ruleMap, update) {
			result += helpers.StrToInt(update[len(update)/2])
		}
	}

	return result
}

func checkUpdate(ruleMap map[string][]string, currUpdate []string) bool {
	seen := make(map[string]bool)
	for _, currPage := range currUpdate {
		for _, pages := range ruleMap[currPage] {
			if _, exists := seen[pages]; exists {
				return false
			}
		}
		seen[currPage] = true
	}

	return true
}

func getRuleMap(r string) map[string][]string {
	ruleMap := make(map[string][]string)
	rules := strings.SplitSeq(r, "\n")
	for rule := range rules {
		pages := strings.Split(rule, "|")
		ruleMap[pages[0]] = append(ruleMap[pages[0]], pages[1])
	}

	return ruleMap
}

func parseInput(filename string) (map[string][]string, [][]string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(file), "\n\n")
	ruleMap := getRuleMap(parts[0])
	updatesRaw := strings.Split(parts[1], "\n")

	var updates [][]string
	for _, update := range updatesRaw {
		if len(update) == 0 {
			continue
		}
		updates = append(updates, strings.Split(update, ","))
	}

	return ruleMap, updates
}
