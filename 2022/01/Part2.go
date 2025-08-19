package main

import (
	"bufio"
	"log"
	"os"
	"sort"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Toppers []int

func Part2() int {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	top_three := make(Toppers, 4)
	calories := 0

	for scanner.Scan() {
		token := scanner.Text()
		if len(token) < 1 {
			top_three[3] = calories
			calories = 0
			sort.Sort(top_three)
		} else {
			calorie := helpers.StrToInt(token)
			calories += calorie
		}
	}

	total_calorie := 0
	for i, calorie := range top_three {
		if i == 3 {
			break
		}
		total_calorie += calorie
	}

	return total_calorie
}

func (s Toppers) Len() int {
	return len(s)
}

func (s Toppers) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Toppers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
