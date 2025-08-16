package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	var sum int

	for _, word := range output {
		firstDigit := searchFront(word)
		lastDigit := searchBack(word)

		calibrationNumber := firstDigit + lastDigit
		// fmt.Printf("first= %s \t\t last= %s\n", firstDigit, lastDigit)
		// fmt.Println(word, calibrationNumber)
		sum += getIntNum(calibrationNumber)
	}
	fmt.Println(sum)
}

func getIntNum(s string) int {
	num, _ := strconv.Atoi(s)
	// if err != nil {
	//        return 0
	// }
	return num
}

var strNums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func searchFront(w string) string {
	s := ""
	for _, letter := range w {
		s += string(letter)
		isNumber, index := isNum(s)
		if isNumber {
			return strconv.Itoa(index)
		}
		if unicode.IsDigit(letter) {
			return string(letter)
		}
	}
	return ""
}

func isNum(s string) (bool, int) {
	for i, num := range strNums {
		if strings.Contains(s, num) {
			return true, i + 1
		}
	}
	return false, 0
}

func searchBack(w string) string {
	s := ""
	for i := len(w) - 1; i >= 0; i-- {
		s = string(w[i]) + s
		isNumber, index := isNum(s)
		if isNumber {
			return strconv.Itoa(index)
		}
		if unicode.IsDigit(rune(w[i])) {
			return string(w[i])
		}
	}
	return ""
}
