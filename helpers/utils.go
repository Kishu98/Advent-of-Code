package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func StrToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
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
