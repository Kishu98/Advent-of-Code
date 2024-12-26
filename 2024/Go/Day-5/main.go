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
	file1, err := os.OpenFile("InputA.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.OpenFile("InputB.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.OpenFile("InputB.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)
    scanner3 := bufio.NewScanner(file3)

	ruleMap := make(map[string][]string)

	for scanner1.Scan() {
		rule := strings.Split(scanner1.Text(), "|")
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	sum := 0

	for scanner2.Scan() {
		update := strings.Split(scanner2.Text(), ",")
		if !checkUpdate(update, ruleMap) {
			continue
		}
		num, _ := strconv.Atoi(update[(len(update)-1)/2])
		sum += num
	}

	for scanner3.Scan() {
		update1 := strings.Split(scanner3.Text(), ",")
		if !checkUpdateCorrect(update1, ruleMap) {
			continue
		}
		num1, _ := strconv.Atoi(update1[(len(update1)-1)/2])
		sum -= num1
	}

	fmt.Println(sum)

}

func checkUpdateCorrect(update []string, ruleMap map[string][]string) bool {
	for i, page := range update {
		for j := 0; j < i; j++ {
			for k := 0; k < len(ruleMap[page]); k++ {
				// fmt.Println("Checking", update[j], "with ruleMap number:",page, ruleMap[page][k])
				if update[j] == ruleMap[page][k] {
					return false
				}
			}

		}
	}

	return true
}

func checkUpdate(update []string, ruleMap map[string][]string) bool {
	for i, page := range update {
		for j := 0; j < i; j++ {
			for k := 0; k < len(ruleMap[page]); k++ {
				// fmt.Println("Checking", update[j], "with ruleMap number:",page, ruleMap[page][k])
				if update[j] == ruleMap[page][k] {
					fixUpdate(update, ruleMap)
				}
			}

		}
	}

	return true
}

func fixUpdate(update []string, ruleMap map[string][]string) {
	for i, page := range update {
		for j := 0; j < i; j++ {
			for k := 0; k < len(ruleMap[page]); k++ {
				if update[j] == ruleMap[page][k] {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}

	checkUpdate(update, ruleMap)
}
