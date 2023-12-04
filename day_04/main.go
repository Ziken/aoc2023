package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"math"
	"regexp"
	"strconv"
)

const INPUT_FILE = "day_04/input.txt"

func bothParts(data [][]byte) (out1, totalScratchcardAmount int) {
	var integerRegexp = regexp.MustCompile("([\\d|]+)")
	var globalWinnings = map[int]int{}

	for _, line := range data {
		var winningNumbers = map[string]bool{}
		var validNumbersAmount = 0
		var dividerIndex = 0

		var foundNumbers = integerRegexp.FindAll(line, -1)

		// omit first number (Card number)
		for _, rawNum := range foundNumbers[1:] {
			var s = string(rawNum)
			if s == "|" {
				break
			}
			winningNumbers[s] = true
			dividerIndex++
		}
		// omit divider
		for _, rawNum := range foundNumbers[dividerIndex+1:] {
			if _, ok := winningNumbers[string(rawNum)]; ok {
				validNumbersAmount++
			}
		}
		if validNumbersAmount >= 0 {
			out1 += int(math.Pow(2, float64(validNumbersAmount)-1))
			var cardID, _ = strconv.Atoi(string(foundNumbers[0]))
			globalWinnings[cardID] = validNumbersAmount
		}
	}
	// Part 2
	for cardID, validNumbersAmount := range globalWinnings {
		totalScratchcardAmount++
		for i := 1; i <= validNumbersAmount; i++ {
			traverseCards(globalWinnings, cardID+i, &totalScratchcardAmount)
		}
	}
	return
}

func traverseCards(globalWinnings map[int]int, cardID int, totalScratchcardAmount *int) {
	*totalScratchcardAmount++
	var validNumbersAmount, ok = globalWinnings[cardID]
	if !ok {
		return
	}
	for i := 1; i <= validNumbersAmount; i++ {
		traverseCards(globalWinnings, cardID+i, totalScratchcardAmount)
	}
}

func main() {
	var data = utils.GetInput(INPUT_FILE)
	var out1, out2 = bothParts(data)
	fmt.Println("Part 1:", out1)
	fmt.Println("Part 2:", out2)

}
