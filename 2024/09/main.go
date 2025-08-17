package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	str = str[:len(str)-1]

	file := true
	i := 0
	var disk []string
	for _, c := range str {
		id := strconv.Itoa(i)
		if file {
			block, _ := strconv.Atoi(string(c))
			for range block {
				disk = append(disk, id)
			}
			file = false
			i++
		} else {
			block, _ := strconv.Atoi(string(c))
			for range block {
				disk = append(disk, ".")
			}
			file = true
		}
	}
	fmt.Println(disk)

	for j, char := range disk {
		if char == "." {
			for i := len(disk) - 1; i >= 0; i-- {
				if disk[i] != "." && i > j {
					disk[i], disk[j] = disk[j], disk[i]
					break
				}
			}
		}
	}

	var sum int
	for i, num := range disk {
		if num == "." {
			break
		}
		n, _ := strconv.Atoi(string(num))
		sum += i * n
	}

	fmt.Println(sum)
}

func swap(str string, i, j int) string {
	runes := []rune(str)
	runes[i], runes[j] = runes[j], runes[i]

	return string(runes)
}
