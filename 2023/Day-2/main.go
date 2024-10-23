package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	// var games []string
	sum := 0
    powerSum := 0

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), " ")

		for i, g := range game {
			game[i] = strings.TrimRight(g, ",;:")
		}

		gameNum, err := strconv.Atoi(game[1])
		if err != nil {
			log.Fatal(err)
		}

		game = slices.Delete(game, 0, 2)

		// fmt.Println(game)

		var possible = 0

		for i := 0; i < len(game); i = i + 2 {
			val, err := strconv.Atoi(game[i])
			if err != nil {
				log.Fatal(err)
			}
			switch game[i+1] {
			case "blue":
				if val > 14 {
					// fmt.Println("Increaseing possible value:", val, game[i+1])
					possible++
				}
			case "red":
				if val > 12 {
					// fmt.Println("Increaseing possible value:", val, game[i+1])
					possible++
				}
			case "green":
				if val > 13 {
					// fmt.Println("Increaseing possible value:", val, game[i+1])
					possible++
				}
			}

		}
		if possible == 0 {
			sum += gameNum
		}
		// fmt.Println(sum)
		maxBlue := 0
		maxRed := 0
		maxGreen := 0

		for i := 0; i < len(game); i = i + 2 {
			val, err := strconv.Atoi(game[i])
			if err != nil {
				log.Fatal(err)
			}
			switch game[i+1] {
			case "blue":
				if val > maxBlue {
					maxBlue = val
				}
			case "red":
				if val > maxRed {
					maxRed = val
				}
			case "green":
				if val > maxGreen {
					maxGreen = val
				}
			}
		}
        powerSum += maxBlue * maxRed * maxGreen
		// games = append(games, game)
	}
    fmt.Println(powerSum)
	fmt.Println(sum)
}
