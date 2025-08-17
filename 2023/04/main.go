package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.OpenFile("Input.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum = 0
	var lines = []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		Nums := strings.Split(line, "|")

		winningNums := strings.Split(strings.TrimSpace(Nums[0]), " ")
		winningNums = winningNums[2:]
		myNums := strings.Split(strings.TrimSpace(Nums[1]), " ")

		for i, winningNum := range winningNums {
			if winningNum == "" {
				winningNums = append(winningNums[:i], winningNums[i+1:]...)
				continue
			}
		}

		for i, myNum := range myNums {
			if myNum == "" {
				myNums = append(myNums[:i], myNums[i+1:]...)
				continue
			}
		}

		var total = 0

		for _, num1 := range myNums {
			myNum, _ := strconv.Atoi(num1)
			var counter = 0
			for _, num2 := range winningNums {
				winningNum, _ := strconv.Atoi(num2)
				if myNum == winningNum {
					counter++
				}
			}
			for counter > 0 {
				if total == 0 {
					total = 1
					// fmt.Println("Total = ", total)
					counter--
					continue
				}
				// fmt.Println("Total = ", total)
				total = total * 2
				counter--
			}
		}

		sum = sum + total

		// fmt.Println(total)
		// fmt.Println(winningNums)
		// fmt.Println(myNums)
	}
	fmt.Println(sum)
}
