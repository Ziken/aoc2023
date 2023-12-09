package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
)

func parseInput(data [][]byte) (out [][]int) {
	var regexp = regexp.MustCompile(`-?\d+`)
	for _, row := range data {
		var rawNums = regexp.FindAll(row, -1)
		var nums = make([]int, len(rawNums))
		for i, rawNum := range rawNums {
			nums[i], _ = strconv.Atoi(string(rawNum))
		}
		out = append(out, nums)
	}
	return
}

func generateHistory(row []int) [][]int {
	var history = make([][]int, 0)
	history = append(history, row)
	var currentRow = 0
	var allZeros = false
	for !allZeros {
		allZeros = true
		newHistory := make([]int, 0)
		for i, _ := range history[currentRow][1:] {
			var newNum = history[currentRow][i+1] - history[currentRow][i]
			if newNum != 0 {
				allZeros = false
			}
			newHistory = append(newHistory, newNum)
		}
		if len(newHistory) > 0 {
			history = append(history, newHistory)
		}
		currentRow++
	}
	return history
}

func bothParts(numbers [][]int) (outPartOne, outPartTwo int) {
	for _, row := range numbers {
		var history = generateHistory(row)
		for i := range history {
			outPartOne += history[i][len(history[i])-1]
		}
		var extrapolatedValue = 0
		for i := len(history) - 1; i > 0; i-- {
			extrapolatedValue = history[i-1][0] - extrapolatedValue
		}
		outPartTwo += extrapolatedValue
	}
	return
}

func main() {
	var data = utils.GetInput("day_09/input.txt")
	var numbers = parseInput(data)
	var part1, part2 = bothParts(numbers)
	fmt.Println("Part one:", part1)
	fmt.Println("Part two:", part2)
}
