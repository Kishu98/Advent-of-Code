package main

import (
	"log"
	"strings"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part1() int {
	filename := "TestInput.txt"
	cache := make(map[string]int)
	dirsSize := getDirs(filename, cache)

	result := 0
	for _, val := range dirsSize {
		if val <= 100000 {
			result += val
		}
	}

	return result
}

func getSize(dir string, dirsSize map[string]int, dirs map[string][]string, cache map[string]int) int {
	if size, ok := cache[dir]; ok {
		return size
	}

	size := dirsSize[dir]
	for _, child := range dirs[dir] {
		size += getSize(child, dirsSize, dirs, cache)
	}

	cache[dir] = size
	dirsSize[dir] = size
	return size
}

func getDirs(filename string, cache map[string]int) map[string]int {
	dirsSize := make(map[string]int)
	dirs := make(map[string][]string)

	var stack []string
	currentDir := "/"

	if err := helpers.ProcessInput(filename, func(s string) {
		parts := strings.Split(s, " ")

		if parts[0] == "$" && parts[1] == "cd" {
			switch parts[2] {
			case "/":
				stack = []string{"/"}
			case "..":
				stack = stack[:len(stack)-1]
			default:
				stack = append(stack, parts[2])
			}
			currentDir = strings.Join(stack, "/")
			return
		}

		if parts[0] != "$" {
			if parts[0] == "dir" {
				child := currentDir + "/" + parts[1]
				dirs[currentDir] = append(dirs[currentDir], child)
			} else {
				dirsSize[currentDir] += helpers.StrToInt(parts[0])
			}
		}
	}); err != nil {
		log.Fatal(err)
	}

	for dir := range dirs {
		getSize(dir, dirsSize, dirs, cache)
	}

	return dirsSize
}
