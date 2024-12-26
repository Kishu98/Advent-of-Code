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

	var list1 []int
	var list2 []int

	for _, line := range output {
		nums := strings.Fields(line)
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	sort(list1)
	sort(list2)
	var sum int

	for i := 0; i < len(list1); i++ {
		dist := list2[i] - list1[i]
		if dist < 0 {
			dist = dist * -1
		}
		sum = sum + dist
	}

    fmt.Println(sum)

	var similarity_score int

	for _, num := range list1 {
		similarity_score = similarity_score + search(num, list2)
	}

	fmt.Println(similarity_score)
}

func search(v int, l []int) int {
    var times int
    for _, val := range l {
        if v == val {
            times++
        }
    }

    return v * times
}

func sort(l []int) {
	for i := 0; i < len(l)-1; i++ {
		for j := 0; j < len(l)-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
