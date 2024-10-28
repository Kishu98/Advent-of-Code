package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(file))

	lines := strings.Split(string(file), "\n\n")

	seeds := strings.Split(lines[0], " ")
	seeds = append(seeds[:0], seeds[1:]...)
	lines = append(lines[:0], lines[1:]...)
	// fmt.Println(lines[0])

	for _, seed := range seeds {
		fmt.Println(seed)
		for _, line := range lines {
			mapNums := strings.Split(line, "\n")
			mapNums = append(mapNums[:0], mapNums[1:]...)
			// fmt.Println(mapNums, ", ")
            for _, mapNum := range(mapNums) {
                nums := strings.Split(mapNum, " ")
                fmt.Println(nums, ",")
            }
		}
	}

}
