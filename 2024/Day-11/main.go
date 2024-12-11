package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split(string(content)[:len(string(content))-1], " ")

    stoneCount := make(map[int]int)
    for _, s := range str {
        stoneCount[str_to_int(s)] = 1
    }

    for i := 0; i < 75; i++ {
        stoneCount = Blink(stoneCount)
    }

    var ans int
    for _, count := range stoneCount {
        ans += count
    }

    fmt.Println(ans)
}

func Blink(stoneCount map[int]int) map[int]int {
    newStoneCount := make(map[int]int)

    for stone, count := range stoneCount {
        if stone == 0 {
            newStoneCount[1] += count
        } else if len(strconv.Itoa(stone)) % 2 == 0 {
            left, right := splitNumber(stone)
            newStoneCount[left] += count
            newStoneCount[right] += count
        } else {
            newStoneCount[stone * 2024] += count
        }
    }

    return newStoneCount
}

func splitNumber(n int) (int, int) {
    s := strconv.Itoa(n)
    mid := len(s)/2
    left, _ := strconv.Atoi(s[:mid])
    right, _ := strconv.Atoi(s[mid:])

    return left, right
}

func str_to_int(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

