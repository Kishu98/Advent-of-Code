package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1(filename string) int {
	result := 0
	if err := helpers.ProcessInput(filename, func(s string) {
		sum, nums := getSumAndNumbers(s)
		if checkSumPossible(sum, nums, 1, nums[0], false) {
			result += sum
		}
	}); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkSumPossible(sum int, nums []int, index int, current int, toConcat bool) bool {
	if index == len(nums) {
		return sum == current
	}

	next := nums[index]
	if checkSumPossible(sum, nums, index+1, current+next, toConcat) {
		return true
	}

	if checkSumPossible(sum, nums, index+1, current*next, toConcat) {
		return true
	}

	if toConcat {
		concat := helpers.StrToInt(fmt.Sprintf("%d%d", current, next))
		if checkSumPossible(sum, nums, index+1, concat, toConcat) {
			return true
		}
	}

	return false
}

func getSumAndNumbers(s string) (int, []int) {
	parts := strings.Split(s, ":")
	sum := helpers.StrToInt(parts[0])
	nums := getNumbers(strings.Split(parts[1], " "))

	return sum, nums
}

func getNumbers(s []string) []int {
	var nums []int
	for i := range s {
		if strings.TrimSpace(s[i]) == "" {
			continue
		}
		nums = append(nums, helpers.StrToInt(s[i]))
	}

	return nums
}
