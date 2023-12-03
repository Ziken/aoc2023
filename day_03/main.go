package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
	"unicode"
)

const INPUT_DATA = "day_03/input.txt"

type Gear struct {
	x, y int
}

func indexExists(input [][]byte, x, y int) bool {
	return x >= 0 && y >= 0 && len(input) > y && len(input[y]) > x
}

func isSpecialCharAtIndex(input [][]byte, x, y int) bool {
	return indexExists(input, x, y) && input[y][x] != '.' && !unicode.IsDigit(rune(input[y][x]))
}

func partOne(input [][]byte) (out int) {
	var integerRegexp = regexp.MustCompile("\\d+")
	var checkIndexes = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	for y, row := range input {
		var foundIndexes = integerRegexp.FindAllIndex(row, -1)
		for _, foundIndex := range foundIndexes {
			var start, end = foundIndex[0], foundIndex[1]
			var isInvalid = true
			for x := start; x < end; x++ {
				for _, p := range checkIndexes {
					if isSpecialCharAtIndex(input, x+p[0], y+p[1]) {
						isInvalid = false
						break
					}
				}
			}
			if !isInvalid {
				var n, _ = strconv.Atoi(string(input[y][start:end]))
				out += n
			}
		}
	}

	return
}

func partTwo(input [][]byte) (out int) {
	var integerRegexp = regexp.MustCompile("\\d+")
	var checkIndexes = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	var usedGears = map[Gear][]string{}

	for y, row := range input {
		var foundIndexes = integerRegexp.FindAllIndex(row, -1)
		for _, foundIndex := range foundIndexes {
			var start, end = foundIndex[0], foundIndex[1]
			for x := start; x < end; x++ {
				for _, p := range checkIndexes {
					// append number if any of its digits is near the gear
					if isSpecialCharAtIndex(input, x+p[0], y+p[1]) && input[y+p[1]][x+p[0]] == '*' {
						usedGears[Gear{x + p[0], y + p[1]}] = append(usedGears[Gear{x + p[0], y + p[1]}], string(input[y][start:end]))
					}
				}
			}
		}
	}

	// Get two unique parts of gear
	for _, values := range usedGears {
		var part1 = values[0]
		var part2 = ""
		for i := range values {
			if part1 != values[i] {
				part2 = values[i]
				break
			}
		}
		if part2 != "" {
			var n1, _ = strconv.Atoi(part1)
			var n2, _ = strconv.Atoi(part2)
			out += n1 * n2
		}
	}

	return
}

func main() {
	var inputData = utils.GetInput(INPUT_DATA)
	fmt.Println("Part 1: ", partOne(inputData))
	fmt.Println("Part 2: ", partTwo(inputData))
}
