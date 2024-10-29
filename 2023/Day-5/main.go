package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var totalSeeds []int

	lines := strings.Split(string(file), "\n\n")
	seeds := strings.Split(lines[0], " ")
	seeds = seeds[1:]
	for i := 0; i < len(seeds); i = i + 2 {
		start, _ := strconv.Atoi(seeds[i])
        // fmt.Println("Start -->", start)
		end, _ := strconv.Atoi(seeds[i+1])
        end = start + end
        // fmt.Println("End -->", end)
		for j := start; j < end; j++ {
			totalSeeds = append(totalSeeds, j)
		}
	}
	blocks := lines[1:]
	// fmt.Println(totalSeeds)
	// seed := 79
	minLocation := 9999999999999
	for _, seed := range totalSeeds {

		for _, block := range blocks {
			// fmt.Println(block)
			// fmt.Println("----------------------------")
			blk := strings.Split(block, "\n")
			for i := 1; i < len(blk); i++ {
				if blk[i] == "" {
					continue
				}
				// fmt.Println("-->", blk[i])
				nums := strings.Split(blk[i], " ")
				dest, _ := strconv.Atoi(nums[1])
				source, _ := strconv.Atoi(nums[0])
				ran, _ := strconv.Atoi(nums[2])
				if seed >= dest && seed < (dest+ran) {
					// fmt.Println(seed, "+", dest, "-", source)
					seed = seed + source - dest
					break
				}
				// for j := 0; j < 3; j++ {
				// 	nums := strings.Split(blk[i], " ")
				// 	fmt.Println("\t-->", nums[j])
				// }
			}
			// fmt.Println(seed)
		}
		if minLocation > seed {
			minLocation = seed
		}
		// fmt.Println(seed)
	}
	fmt.Println(minLocation)
}
