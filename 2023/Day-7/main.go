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

    scanner := bufio.NewScanner(file)

	var hands []string
	var bids []int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hands = append(hands, line[0])
        bid, err := strconv.Atoi(line[1])
        if err != nil{
            fmt.Println("Error converting bid:", err)
        }

		bids = append(bids, bid)
	}

    
}
