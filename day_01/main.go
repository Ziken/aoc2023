package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
	"unicode"
)

const INPUT_FILE = "day_01/input.txt"

func partOne(input [][]byte) (sum int) {
	var re = regexp.MustCompile("[0-9]")
	for _, row := range input {
		var foundNums = re.FindAllString(string(row), -1)
		var n1, _ = strconv.Atoi(foundNums[0] + foundNums[len(foundNums)-1])

		sum += n1

	}
	return sum
}

func partTwo(input [][]byte) (sum int) {
	var numMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var preparedRegex = "[0-9]"
	var reversedRegex = "[0-9]"
	for key, _ := range numMap {
		numMap[utils.ReverseString(key)] = numMap[key]
		preparedRegex += "|" + key
		reversedRegex += "|" + utils.ReverseString(key)
	}
	var normalRe = regexp.MustCompile(preparedRegex)
	var reversedRe = regexp.MustCompile(reversedRegex)

	for _, row := range input {
		var normalOrderNums = normalRe.FindAll(row, -1)
		var reversedOrderNums = reversedRe.FindAll(utils.ReverseBytes(row), -1)
		var composedNumber = ""
		if unicode.IsDigit(rune(normalOrderNums[0][0])) {
			composedNumber = string(normalOrderNums[0])
		} else {
			composedNumber = numMap[string(normalOrderNums[0])]
		}

		if unicode.IsDigit(rune(reversedOrderNums[0][0])) {
			composedNumber += string(reversedOrderNums[0])
		} else {
			composedNumber += numMap[string(reversedOrderNums[0])]
		}

		var convertedNumber, _ = strconv.Atoi(composedNumber)
		sum += convertedNumber
	}

	return
}

func main() {
	var input = utils.GetInput(INPUT_FILE)
	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}
