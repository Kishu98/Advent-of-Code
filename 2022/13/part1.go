package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func part1(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	blocks := strings.Split(string(file), "\n\n")

	sum := 0
	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		if len(lines) < 2 {
			continue
		}

		left := mustParse(lines[0])
		right := mustParse(lines[1])
		if cmp(left, right) < 0 {
			sum += i + 1
		}
	}

	return sum
}

func cmp(a, b any) int {
	switch av := a.(type) {
	case float64:
		switch bv := b.(type) {
		case float64:
			if av < bv {
				return -1
			}
			if bv < av {
				return 1
			}
			return 0
		default:
			return cmp([]any{av}, b)
		}
	case []any:
		switch bv := b.(type) {
		case float64:
			return cmp(a, []any{bv})
		case []any:
			n := min(len(av), len(bv))
			for i := range n {
				if r := cmp(av[i], bv[i]); r != 0 {
					return r
				}
			}
			if len(av) < len(bv) {
				return -1
			}
			if len(av) > len(bv) {
				return 1
			}
			return 0
		default:
			panic("what the fuck is this?")
		}
	default:
		panic("WHat the fuck is this?")
	}
}

func mustParse(s string) any {
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		log.Fatal(err)
	}

	return v
}
