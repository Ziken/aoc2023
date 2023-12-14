package main

import (
	"bytes"
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
)

type Spring struct {
	Spring []byte
	Parts  []int
}

const (
	PART_OPERATIONAL = '.'
	PART_DAMAGED     = '#'
	PART_UNKNOWN     = '?'
)

func parseData(data [][]byte) []Spring {
	var numbersRegexp = regexp.MustCompile(`\d+`)
	var springs []Spring
	for row := 0; row < len(data); row++ {
		var splitted = bytes.Split(data[row], []byte(" "))
		var spring, rawParts = splitted[0], splitted[1]
		var s = Spring{Spring: spring}
		var parts = numbersRegexp.FindAll(rawParts, -1)
		for _, part := range parts {
			n, _ := strconv.Atoi(string(part))
			s.Parts = append(s.Parts, n)
		}
		springs = append(springs, s)
	}

	return springs
}

func getPermutationNumberOfParts(spring Spring, consideredPartIndex int, startingPos int) (out int) {
	if consideredPartIndex >= len(spring.Parts) {
		// No part to consider
		// check if there is any part different from PART_OPERATIONAL
		var damagedPartLeft = bytes.ContainsAny(spring.Spring[startingPos:], string(PART_DAMAGED))
		if !damagedPartLeft {
			//fmt.Println("Valid", string(spring.Spring))
			return 1
		} else {
			return 0
		}
	}
	if startingPos >= len(spring.Spring) {
		if consideredPartIndex == len(spring.Parts) {
			//fmt.Println("Valid", string(spring.Spring))
			return 1
		}
		return 0
	}

	var consideredPart = spring.Parts[consideredPartIndex]
	for i := startingPos; i < len(spring.Spring); i++ {
		if i > 0 && spring.Spring[i-1] == PART_DAMAGED {
			continue
		}
		var partCounter = 0

		for j := i; j < len(spring.Spring); j++ {
			if spring.Spring[j] == PART_UNKNOWN {
				partCounter += 1
			} else if spring.Spring[j] == PART_DAMAGED {
				partCounter += 1
			} else {
				// PART_OPERATIONAL
				break
			}
			if partCounter == consideredPart {
				// part counter is valid
				if (j + 1) >= len(spring.Spring) {
					out += getPermutationNumberOfParts(spring, consideredPartIndex+1, j+1)
				} else if spring.Spring[j+1] != PART_DAMAGED {
					// make sure that the next part is not damaged
					out += getPermutationNumberOfParts(spring, consideredPartIndex+1, j+2)
				}
				break
			}
		}
		if spring.Spring[i] == PART_DAMAGED {
			return out
		}
	}
	return
}

func assert(s Spring, expected int) {
	var r = getPermutationNumberOfParts(s, 0, 0)
	if r != expected {
		fmt.Println("[ERROR] Expected", expected, "got", r, "for", string(s.Spring))
	} else {
		fmt.Println("[OK] Expected", expected, "got", r, "for", string(s.Spring))
	}
}

func partOne(springs []Spring) (out int) {
	//assert(Spring{Spring: []byte("???.###"), Parts: []int{1, 1, 3}}, 1)
	//assert(Spring{Spring: []byte(".??..??...?##."), Parts: []int{1, 1, 3}}, 4)
	//assert(Spring{Spring: []byte("?#?#?#?#?#?#?#?"), Parts: []int{1, 3, 1, 6}}, 1)
	//assert(Spring{Spring: []byte("????.#...#..."), Parts: []int{4, 1, 1}}, 1)
	//assert(Spring{Spring: []byte("????.######..#####."), Parts: []int{1, 6, 5}}, 4)
	//assert(Spring{Spring: []byte("?###.???????"), Parts: []int{3, 2, 1}}, 10)
	//assert(Spring{Spring: []byte(".#??.???????"), Parts: []int{3, 2, 1}}, 10)
	//assert(Spring{Spring: []byte(".#??????????"), Parts: []int{3, 2, 1}}, 10)
	//assert(Spring{Spring: []byte("?###????????"), Parts: []int{3, 2, 1}}, 10)
	//assert(Spring{Spring: []byte("#??#"), Parts: []int{2}}, 0)
	//assert(Spring{Spring: []byte("#?????????#"), Parts: []int{2}}, 0)
	//assert(Spring{Spring: []byte("#?????????"), Parts: []int{2}}, 1)
	//assert(Spring{Spring: []byte("#????.?##??."), Parts: []int{3, 1, 2}}, 1)
	//assert(Spring{Spring: []byte("#????.?##??."), Parts: []int{3, 1, 2}}, 1)
	//assert(Spring{Spring: []byte(".??#?#."), Parts: []int{3}}, 1)
	//assert(Spring{Spring: []byte("?????#???.?"), Parts: []int{4, 1}}, 8)
	//assert(Spring{Spring: []byte("????#?#?####????."), Parts: []int{1, 11, 1}}, 4)
	//assert(Spring{Spring: []byte("????"), Parts: []int{2, 2}}, 0)
	//assert(Spring{Spring: []byte("#.....??"), Parts: []int{1, 1}}, 2)
	//assert(Spring{Spring: []byte("?.??.??.??#????."), Parts: []int{2, 1, 6}}, 4)
	//assert(Spring{Spring: []byte("#..?????#?"), Parts: []int{1, 1, 5}}, 1)
	//assert(Spring{Spring: []byte("??????.????."), Parts: []int{4, 1}}, 13)
	//assert(Spring{Spring: []byte("???#???????????????????#????????"), Parts: []int{2, 2}}, 4)
	//assert(Spring{Spring: []byte("???.#??.#?"), Parts: []int{1, 1, 1}}, 4)
	//assert(Spring{Spring: []byte("??.???.??????"), Parts: []int{1, 1}}, 47)
	//assert(Spring{Spring: []byte("????#??????#????"), Parts: []int{8}}, 1)
	//assert(Spring{Spring: []byte("????#??????#????"), Parts: []int{1, 8}}, 4)
	//assert(Spring{Spring: []byte("#???#??????#????"), Parts: []int{1, 8}}, 1)
	//assert(Spring{Spring: []byte("????#?????#????"), Parts: []int{1, 8}}, 2)
	for _, spring := range springs {
		//fmt.Println("Spring", string(spring.Spring), getPermutationNumberOfParts(spring, 0, 0))
		out += getPermutationNumberOfParts(spring, 0, 0)
	}
	return
}

func partTwo(springs []Spring) (out int) {
	var newSprings = make([]Spring, 0)
	for _, spring := range springs {
		var newSpring = Spring{}
		for i := 0; i < 5; i++ {
			newSpring.Spring = append(newSpring.Spring, spring.Spring...)
			if i != 4 {
				newSpring.Spring = append(newSpring.Spring, PART_UNKNOWN)
			}
			newSpring.Parts = append(newSpring.Parts, spring.Parts...)
		}
		newSprings = append(newSprings, newSpring)
	}

	for i, spring := range newSprings {
		fmt.Println("Spring", i+1, 1000)
		//fmt.Println("Spring", string(spring.Spring), getPermutationNumberOfParts(spring, 0, 0))
		out += getPermutationNumberOfParts(spring, 0, 0)
	}
	return
}

func main() {
	var rawData = utils.GetInput("day_12/input.txt")
	var springs = parseData(rawData)
	fmt.Println("Part one:", partOne(springs))
	fmt.Println("Part Two:", partTwo(springs))
}
