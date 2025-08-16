package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() int {
	// Reading the Input File
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	max_calories := 0
	calories := 0
	for scanner.Scan() {
		token := scanner.Text()
		if len(token) < 1 {
			max_calories = max(max_calories, calories)
			calories = 0
		} else {
			calories += str_to_int(token)
		}
	}

	return max_calories
}

func str_to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
