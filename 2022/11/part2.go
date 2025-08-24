package main

import (
	"sort"
)

func part2(filename string) int {
	monkeys := getMonkeys(filename)

	modulus := 1
	for _, m := range monkeys {
		modulus *= m.divisibleBy
	}

	round := 10000
	throwItem(round, monkeys, func(i *int) {
		*i = *i % modulus
	})

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].totalInspected > monkeys[j].totalInspected
	})

	return monkeys[0].totalInspected * monkeys[1].totalInspected
}
