package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

const INPUT_FILE = "day_01/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getInput() (in [][]byte) {
	file, errFile := os.Open(INPUT_FILE)
	check(errFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		in = append(in, []byte(row))
	}

	return
}
func reverseString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)
	// Get the length of the slice
	n := len(runes)

	// Swap the runes from the ends towards the center
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}
func reverseBytes(s []byte) []byte {
	// Convert string to a slice of runes
	// Get the length of the slice
	n := len(s)

	// Swap the runes from the ends towards the center
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}

	// Convert the slice of runes back to a string
	return s
}
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
		numMap[reverseString(key)] = numMap[key]
		preparedRegex += "|" + key
		reversedRegex += "|" + reverseString(key)
	}
	var normalRe = regexp.MustCompile(preparedRegex)
	var reversedRe = regexp.MustCompile(reversedRegex)

	for _, row := range input {
		var normalOrderNums = normalRe.FindAll(row, -1)
		var reversedOrderNums = reversedRe.FindAll(reverseBytes(row), -1)
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
	var input = getInput()
	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}
