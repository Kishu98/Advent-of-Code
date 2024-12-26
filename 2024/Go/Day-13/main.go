package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split(string(content), "\n\n")

	// fmt.Println(str)
	for _, s := range str {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println(s)
        fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	}

	for _, s := range str {
        if s == "\n" {
            continue
        }
	    spl := strings.Split(s, "\n")
	    for _, k := range spl {
	        eq := strings.Split(k, ":")[1]
	        fmt.Println(eq)
	    }
	}

}
