package main

import (
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

func part2(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var packets []any
	for line := range strings.SplitSeq(string(file), "\n") {
		if line == "" {
			continue
		}
		packets = append(packets, mustParse(line))
	}

	div1 := mustParse("[[2]]")
	div2 := mustParse("[[6]]")
	packets = append(packets, div1, div2)

	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	pos1, pos2 := 0, 0
	for i, p := range packets {
		if reflect.DeepEqual(p, div1) {
			pos1 = i + 1
		} else if reflect.DeepEqual(p, div2) {
			pos2 = i + 1
		}
	}

	return pos1 * pos2
}
