package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/Kishu98/AdventOfCode/helpers"
)

type Point struct {
	x, y int
}

type Sensor struct {
	p Point
	b Beacon
	d int
}

type Beacon struct {
	p Point
}

func part1(filename string) int {
	rowY := 2000000
	intervals := [][2]int{}
	if err := helpers.ProcessInput(filename, func(s string) {
		var sensor Sensor
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.p.x, &sensor.p.y, &sensor.b.p.x, &sensor.b.p.y)
		sensor.d = abs(sensor.p.x-sensor.b.p.x) + abs(sensor.p.y-sensor.b.p.y)
		if abs(sensor.p.y-rowY) > sensor.d {
			return
		}
		interval := getInterval(sensor, rowY)
		intervals = append(intervals, interval)
	}); err != nil {
		log.Fatal(err)
	}

	merged := mergeIntervals(intervals)

	result := merged[0][1] - merged[0][0]

	return result
}

func mergeIntervals(intervals [][2]int) [][2]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][2]int{}
	cur := intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] <= cur[1]+1 {
			if interval[1] > cur[1] {
				cur[1] = interval[1]
			}
		} else {
			merged = append(merged, cur)
			cur = interval
		}
	}
	merged = append(merged, cur)

	return merged
}

func getInterval(sensor Sensor, rowY int) [2]int {
	coverStart := sensor.p.x - (sensor.d - abs(sensor.p.y-rowY))
	coverEnd := sensor.p.x + (sensor.d - abs(sensor.p.y-rowY))

	return [2]int{coverStart, coverEnd}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
