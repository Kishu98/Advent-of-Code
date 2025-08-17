package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re2 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	// re2 := regexp.MustCompile(`do\(\)|don't\(\)`)

	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)

	mult := re2.FindAllString(str, -1)
	var nums [][]string

	reNum := regexp.MustCompile(`\d{1,3}`)
	flag := true

	for _, mul := range mult {
		if mul == "don't()" {
			flag = false
		} else if mul == "do()" {
			flag = true
		} else {
			if flag {
				nums = append(nums, reNum.FindAllString(mul, -1))
			}
		}
	}

	sum := 0
	for _, n := range nums {
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		sum += n1 * n2
	}

	fmt.Println(sum)
}
