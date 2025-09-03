package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func StrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func ProcessInput(filename string, lineProcessor func(string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineProcessor(scanner.Text())
	}
	return scanner.Err()
}

func MapAlphatoNum(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func ParseGrid[T any](filename string, conv func(rune) T) [][]T {
	var grid [][]T
	if err := ProcessInput(filename, func(s string) {
		row := make([]T, len(s))
		for y, c := range s {
			row[y] = conv(c)
		}
		grid = append(grid, row)
	}); err != nil {
		log.Fatal(err)
	}
	return grid
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}
