package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type PriorityQueueItem struct {
	Point    Point
	Distance int
	Index    int
}

type PriorityQueue []PriorityQueueItem

func main() {
	content, err := os.ReadFile("Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var points []Point
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			continue
		}

		c := strings.Split(l, ",")
		points = append(points, Point{str_to_int(c[0]), str_to_int(c[1])})
	}
    // Part - 2
	// for i := 0; i < len(points); i++ {
	// 	grid := CreateGrid(points[:i])
	// 	_, err := Dijkstra(grid, Point{0, 0}, Point{70, 70})
	//        if err != nil {
	//            fmt.Println(points[i-1])
	//            break
	//        }
	// }

    // Part - 1
	grid := CreateGrid(points[:1024])


	path, err := Dijkstra(grid, Point{0,0}, Point{70,70})

    VisualizeGrid(grid, path)

    fmt.Println(len(path)- 1)
}

func VisualizeGrid(grid [][]string, path []Point) {
    for r, row := range grid {
        for c, cell := range row {
            point := Point{r, c}
            if containsPoint(path, point) {
                fmt.Print("\033[31m", ".", "\033[0m")
            } else {
                fmt.Print(cell)
            }
        }
        fmt.Println()
    }
}

func containsPoint(path []Point, p Point) bool {
    for _, point := range path {
        if point == p {
            return true
        }
    }

    return false
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(PriorityQueueItem)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
    n := len(old)
    item := old[n-1]
	item.Index = -1
	*pq = old[:n-1]
	return item
}

func Dijkstra(grid [][]string, start Point, end Point) ([]Point, error) {
	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)
	prev := make(map[Point]Point)
	directions := []Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

    for i := range dist {
        dist[i] = make([]int, cols)
        for j := range dist[i] {
            dist[i][j] = math.MaxInt32
        }
    }

    dist[start.X][start.Y] = 0

    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, PriorityQueueItem{Point: start, Distance: 0})

    for pq.Len() > 0 {
        current := heap.Pop(pq).(PriorityQueueItem)
        currPoint := current.Point

        if currPoint == end {
            break
        }

        for _, dir := range directions {
            neighbour := Point{currPoint.X + dir.X, currPoint.Y + dir.Y}
            if neighbour.X < 0 || neighbour.Y < 0 || neighbour.X >= rows || neighbour.Y >= cols {
                continue
            }
            if grid[neighbour.X][neighbour.Y] == "#" {
                continue
            }

            newDist := dist[currPoint.X][currPoint.Y] + 1
            if newDist < dist[neighbour.X][neighbour.Y] {
                dist[neighbour.X][neighbour.Y] = newDist
                prev[neighbour] = currPoint
                heap.Push(pq, PriorityQueueItem{Point: neighbour, Distance: newDist})
            }
        }
    }

    path := []Point{}
    for at := end; at != start; at = prev[at] {
        path = append([]Point{at}, path...)
        if _, found := prev[at]; !found {
            return nil, fmt.Errorf("No path found")
        }
    }

    path = append([]Point{start}, path...)

    return path, nil
}

func CreateGrid(points []Point) [][]string {
	gridSize := 71
	grid := make([][]string, gridSize)

	for i := 0; i < gridSize; i++ {
		grid[i] = make([]string, gridSize)
		for j := 0; j < gridSize; j++ {
			if corruptedPoint(i, j, points) {
				grid[i][j] = "#"
			} else {
				grid[i][j] = "."
			}
		}
	}

	return grid
}

func corruptedPoint(x, y int, points []Point) bool {
	for _, point := range points {
		if point.X == y && point.Y == x {
			return true
		}
	}

	return false
}

func str_to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
