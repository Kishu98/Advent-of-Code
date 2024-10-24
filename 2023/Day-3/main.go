package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {

	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	// var coor:= []struct{
	//     x:0,
	//     y:0,
	//
	var coors [][]int
	var input [][]string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		input = append(input, line)
		for i, char := range line {
			// matrix[lineNum][i] := char
			// charNum := 0
			// fmt.Printf("%s ", char)
			if !strings.ContainsAny(char, "1234567890.") {
				// fmt.Printf("Found %s at (%d,%d)", char, lineNum, i)
				// x,y := lineNum, i
				coors = append(coors, []int{lineNum, i})
			}
		}
		// fmt.Println()
		lineNum++

	}
	// fmt.Println(input)

	for _, coor := range coors {
		// symbol := input[coor[0]][coor[1]]
        if unicode.IsDigit([]rune(input[coor[0]-1][coor[1]-1])[0]) {
			fmt.Println(input[coor[0]-1][coor[1]-1])
		}
		if unicode.IsDigit([]rune(input[coor[0]-1][coor[1]])[0]) {
			fmt.Println(input[coor[0]-1][coor[1]])
		}
		if unicode.IsDigit([]rune(input[coor[0]-1][coor[1]+1])[0]) {
			fmt.Println(input[coor[0]-1][coor[1]+1])
		}
		if unicode.IsDigit([]rune(input[coor[0]][coor[1]-1])[0]) {
			fmt.Println(input[coor[0]][coor[1]-1])
		}
		if unicode.IsDigit([]rune(input[coor[0]][coor[1]+1])[0]) {
			fmt.Println(input[coor[0]][coor[1]+1])
		}
		if unicode.IsDigit([]rune(input[coor[0]+1][coor[1]-1])[0]) {
			fmt.Println(input[coor[0]+1][coor[1]-1])
		}
		if unicode.IsDigit([]rune(input[coor[0]+1][coor[1]])[0]) {
			fmt.Println(input[coor[0]+1][coor[1]])
		}
		if unicode.IsDigit([]rune(input[coor[0]+1][coor[1]+1])[0]) {
			fmt.Println(input[coor[0]+1][coor[1]+1])
		}
	}
}
