package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
)

const INPUT_FILE = "day_02/input.txt"

func bothParts(input [][]byte) (part1, part2 int) {
	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14
	var gameRegexp = regexp.MustCompile("Game (\\d+)")
	var squareRegexp = regexp.MustCompile("(\\d+) (\\w+)")
	for _, row := range input {
		var gameID, _ = strconv.Atoi(string(gameRegexp.FindSubmatch(row)[1]))
		var squares = squareRegexp.FindAllSubmatch(row, -1)
		var isGamePossible = true
		var maxGameRed, maxGameGreen, maxGameBlue = 0, 0, 0
		for _, square := range squares {
			switch string(square[2]) {
			case "blue":
				n, _ := strconv.Atoi(string(square[1]))
				if n > maxGameBlue {
					maxGameBlue = n
				}
				if n > MAX_BLUE {
					isGamePossible = false
				}
				break
			case "red":
				n, _ := strconv.Atoi(string(square[1]))
				if n > maxGameRed {
					maxGameRed = n
				}
				if n > MAX_RED {
					isGamePossible = false
				}
				break
			case "green":
				n, _ := strconv.Atoi(string(square[1]))
				if n > maxGameGreen {
					maxGameGreen = n
				}
				if n > MAX_GREEN {
					isGamePossible = false
				}
				break
			default:
				panic("Unknown color")
			}
		}
		if isGamePossible {
			part1 += gameID
		}
		part2 += maxGameRed * maxGameGreen * maxGameBlue
	}

	return
}

func main() {
	var input = utils.GetInput(INPUT_FILE)
	var part1, part2 = bothParts(input)
	fmt.Println("Part One:", part1)
	fmt.Println("Part Two:", part2)
}
