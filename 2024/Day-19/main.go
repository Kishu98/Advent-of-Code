package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split(string(content), "\n\n")

	towels := strings.Split(str[0], ", ")
	patterns := strings.Split(str[1], "\n")

	result := 0
    totalWays := 0
	for _, pattern := range patterns {
		p := strings.TrimSpace(pattern)
		if p == "" {
			continue
		}

		if checkPossible(p, towels, map[string]bool{}) {
			result++
		}
        ways := countWays(p ,towels, map[string]int{})
        totalWays += ways
	}

	fmt.Println(result)
    fmt.Println(totalWays)
}

func countWays(s string, towels []string, memo map[string]int) int {
    if s == "" {
        return 1
    }

    if val, exists := memo[s]; exists {
        return val
    }

    total := 0
    for _, towel := range towels {
        if strings.HasPrefix(s, towel) {
            remaining := strings.TrimPrefix(s, towel)
            total += countWays(remaining, towels, memo)
        }
    }

    memo[s] = total

    return total
}

func checkPossible(s string, towels []string, memo map[string]bool) bool {
	if s == "" {
		return true
	}

	if val, exists := memo[s]; exists {
		return val
	}

	for _, towel := range towels {
		if strings.HasPrefix(s, towel) {
			remaining := strings.TrimPrefix(s, towel)
			if checkPossible(remaining, towels, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false

	return false
}
