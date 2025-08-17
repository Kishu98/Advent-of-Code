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
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var safe bool
	var cnt int
	cnt = 0

	for _, line := range lines {
		l := strings.Fields(line)
		// n1, _ := strconv.Atoi(string(l[0]))
		// n2, _ := strconv.Atoi(string(l[1]))

		safe = checkSafe(l) || checkTolerateSafe(l)
		if safe {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func checkTolerateSafe(l []string) bool {
    for i := 0; i < len(l); i++ {
        data := append([]string{}, l[:i]...)
        data = append(data, l[i+1:]...)
        if checkSafe(data) {
            return true
        }
    }
    return false
}

func checkSafe(l []string) bool {
    if (checkDecr(l)||checkIncr(l)) && checkRange(l) {
        return true
    }
    return false
}

func checkDecr(l []string) bool {
	for i := 0; i < len(l)-1; i++ {
		n1, _ := strconv.Atoi(l[i])
		n2, _ := strconv.Atoi(l[i+1])
		if n1 < n2 {
			return false
		}
	}
	return true
}

func checkIncr(l []string) bool {
	for i := 0; i < len(l)-1; i++ {
		n1, _ := strconv.Atoi(l[i])
		n2, _ := strconv.Atoi(l[i+1])
		if n2 < n1 {
			return false
		}
	}

	return true
}

func checkRange(l []string) bool {
	for i := 0; i < len(l)-1; i++ {
		n1, _ := strconv.Atoi(l[i])
		n2, _ := strconv.Atoi(l[i+1])
		if absVal(n2-n1) > 3 || absVal(n2-n1) < 1 {
			return false
		}
	}

	return true
}

func absVal(n int) int {
	if n < 0 {
		return n * -1
	}
    return n
}

