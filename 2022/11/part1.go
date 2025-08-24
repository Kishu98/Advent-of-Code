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
	reduceWorry := func(i int) int {
		return i / 3
	}

	return calculateMonkeyBusiness(monkeys, round, reduceWorry)
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

func calculateMonkeyBusiness(monkeys []Monkey, round int, reduceWorry func(int) int) int {
	throwItem(round, monkeys, func(i int) int {
		return reduceWorry(i)
	})

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].totalInspected > monkeys[j].totalInspected
	})

	return monkeys[0].totalInspected * monkeys[1].totalInspected
}

func throwItem(round int, monkeys []Monkey, reduceFunc func(int) int) {
	for range round {
		for i := range monkeys {
			itemsToThrow := monkeys[i].items
			monkeys[i].items = []int{}
			for _, item := range itemsToThrow {
				monkeys[i].totalInspected++

				item = monkeys[i].operation(item)
				item = reduceFunc(item)

				var throwToMonkey int
				if item%monkeys[i].divisibleBy == 0 {
					throwToMonkey = monkeys[i].throwTo[0]
				} else {
					throwToMonkey = monkeys[i].throwTo[1]
				}
				monkeys[throwToMonkey].items = append(monkeys[throwToMonkey].items, item)
			}
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
						return old + old
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
