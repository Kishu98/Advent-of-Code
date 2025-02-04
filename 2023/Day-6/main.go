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


    // Here we read the file
	file, err := os.OpenFile("Input.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var time []string
	var distance []string

    // Here I'm separating time and distance from the file
	for scanner.Scan() {
		if len(time) == 0 {
			time = append(time, scanner.Text())
			continue
		}
		distance = append(distance, scanner.Text())
	}

	time = strings.Split(time[0], " ")

    // Removing whitespaces from the time slice
	for i := 0; i < len(time); i++ {
		if time[i] == "" {
			time = append(time[:i], time[i+1:]...)
			i--
		}
	}

	distance = strings.Split(distance[0], " ")

    // Removing whitespaces from the distance slice
	for i := 0; i < len(distance); i++ {
		if distance[i] == "" {
			distance = append(distance[:i], distance[i+1:]...)
			i--
		}
	}

	time = time[1:]
	distance = distance[1:]
	finalTime, _ := strconv.Atoi(strings.Join(time, ""))
	finalDistance,_ := strconv.Atoi(strings.Join(distance, ""))
	c := 0

	// for i := 0; i < len(finalTime); i++ {
	t := finalTime
	d := finalDistance
	count := 0

    // Here the main this is happening which is I am calculating the distance and the checking if it 
    // greater than the distance. This works, it was easy too.
	for j := 1; j <= t; j++ {
		dist := j * (t - j)
		if dist > d {
			count++
		}
	}
	if c == 0 {
		c = count
	} else {
		c = c * count
	}
	// }

	fmt.Println(c)
	// for _, t := range time {
	// 	fmt.Println(t)
	// }
	//
	//    for _, d := range(distance) {
	//        fmt.Println(d)
	//    }
}
