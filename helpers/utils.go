package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Str_to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func ProcessInput(lineProcessor func(string)) {
	file, err := os.Open("Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineProcessor(scanner.Text())
	}
}
