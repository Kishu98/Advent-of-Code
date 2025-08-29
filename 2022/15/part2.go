package main

import (
	"fmt"

	"github.com/Kishu98/AdventOfCode/helpers"
)

func part2(filename string) int {
	var sensors []Sensor
	helpers.ProcessInput(filename, func(s string) {
		var sensor Sensor
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.p.x, &sensor.p.y, &sensor.b.p.x, &sensor.b.p.y)
		sensor.d = abs(sensor.p.x-sensor.b.p.x) + abs(sensor.p.y-sensor.b.p.y)
		sensors = append(sensors, sensor)
	})

	for rowY := 0; rowY <= 4000000; rowY++ {
		intervals := [][2]int{}
		for _, sensor := range sensors {
			if abs(sensor.p.y-rowY) > sensor.d {
				continue
			}
			interval := getInterval(sensor, rowY)
			intervals = append(intervals, interval)
		}

		merged := mergeIntervals(intervals)
		if len(merged) > 1 {
			x := merged[0][1] + 1
			return x*4000000 + rowY
		}
	}

	return 0
}
