package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Monkey struct {
	ID             int
	items          []int
	operation      func(int) int
	divisibleBy    int
	throwTo        [2]int
	totalInspected int
}

func part1(filename string) int {
	monkeys := getMonkeys(filename)

	round := 20
	throwItem(round, monkeys, func(i *int) {
		*i = *i / 3
	})

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].totalInspected > monkeys[j].totalInspected
	})

	return monkeys[0].totalInspected * monkeys[1].totalInspected
}

func getMonkeys(filename string) []Monkey {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(file), "\n\n")

	var monkeys []Monkey
	for i := range parts {
		monkeys = append(monkeys, createMonkey(parts[i]))
	}

	return monkeys
}

func throwItem(round int, monkeys []Monkey, reduceFunc func(*int)) {
	for range round {
		for _, monkey := range monkeys {
			if len(monkey.items) == 0 {
				continue
			}
			for _, item := range monkey.items {
				monkeys[monkey.ID].totalInspected++
				item = monkey.operation(item)
				reduceFunc(&item)
				if item%monkey.divisibleBy == 0 {
					monkeys[monkey.throwTo[0]].items = append(monkeys[monkey.throwTo[0]].items, item)
				} else {
					monkeys[monkey.throwTo[1]].items = append(monkeys[monkey.throwTo[1]].items, item)
				}
			}
			monkeys[monkey.ID].items = []int{}
		}
	}
}

func createMonkey(monkeyDetails string) Monkey {
	var monkey Monkey
	parts := strings.SplitSeq(monkeyDetails, "\n")
	for part := range parts {
		part = strings.Trim(part, " ")
		firstWord := strings.Split(part, " ")
		switch firstWord[0] {
		case "Monkey":
			fmt.Sscanf(part, "Monkey %d:", &monkey.ID)
		case "Starting":
			part = strings.TrimPrefix(part, "Starting items:")
			nums := strings.Split(part, ", ")
			for _, num := range nums {
				monkey.items = append(monkey.items, helpers.StrToInt(strings.Trim(num, " ")))
			}
		case "Operation:":
			var op string
			var num string
			fmt.Sscanf(part, "Operation: new = old %s %s", &op, &num)
			switch op {
			case "*":
				monkey.operation = func(old int) int {
					switch num {
					case "old":
						return old * old
					default:
						return old * helpers.StrToInt(num)
					}
				}
			case "+":
				monkey.operation = func(old int) int {
					switch num {
					case "old":
						return old * old
					default:
						return old + helpers.StrToInt(num)
					}
				}
			}
		case "Test:":
			fmt.Sscanf(part, "Test: divisible by %d", &monkey.divisibleBy)
		case "If":
			var boolVal bool
			var throwTo int
			fmt.Sscanf(part, "If %t: throw to monkey %d", &boolVal, &throwTo)
			if boolVal {
				monkey.throwTo[0] = throwTo
			} else {
				monkey.throwTo[1] = throwTo
			}
		}
	}
	return monkey
}
