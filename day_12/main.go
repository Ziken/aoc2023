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

func getPermutationNumberOfParts(spring Spring, currentPartIndex int, startingPos int) (validOrder int) {
	if startingPos == len(spring.Spring) {
		if currentPartIndex == len(spring.Parts) {
			fmt.Println("Valid", string(spring.Spring))
			return 1
		}
		// End of Spring, not parts to match
		return 0
	}
	if currentPartIndex == len(spring.Parts) {
		// End of parts
		// check if there is any part different than PART_OPERATIONAL
		var b = bytes.ContainsAny(spring.Spring[startingPos:], string(PART_UNKNOWN)+string(PART_DAMAGED))
		if b {
			return 0
		}
		fmt.Println("Valid", string(spring.Spring))
		return 1
	}
	var currentPartCounter = spring.Parts[currentPartIndex]
	var i = startingPos
	for ; i < len(spring.Spring); i++ {
		var cpSpring = make([]byte, len(spring.Spring))
		copy(cpSpring, spring.Spring)
		var usedParts = 0
		var j = i
		for z := j; z > 0 && cpSpring[z-1] == PART_DAMAGED; z++ {
			j++
		}
		for j < len(cpSpring) {
			var prevJ = j
			//if j > 0 && cpSpring[j-1] == PART_DAMAGED {
			//	j++
			//	continue
			//}
			for ; j < len(cpSpring) && usedParts < currentPartCounter && cpSpring[j] == PART_UNKNOWN; j++ {
				usedParts++
				cpSpring[j] = PART_DAMAGED
			}
			for ; j < len(cpSpring) && cpSpring[j] == PART_DAMAGED && usedParts < currentPartCounter; j++ {
				usedParts++
			}
			if prevJ == j {
				// no modification done. Exit this variant
				break
			}
			if usedParts > currentPartCounter {
				// to many parts used. Exit this variant
				break
			}
			if usedParts == currentPartCounter {
				if j < len(cpSpring) {
					if cpSpring[j] == PART_DAMAGED {
						// reverse, invalid end
						break
					} else {
						cpSpring[j] = PART_OPERATIONAL
						j++
						validOrder += getPermutationNumberOfParts(Spring{Spring: cpSpring, Parts: spring.Parts}, currentPartIndex+1, j)
						break
					}
				} else {
					// end of spring if (currentPartIndex + 1) == len(spring.Parts) {
					validOrder += getPermutationNumberOfParts(Spring{Spring: cpSpring, Parts: spring.Parts}, currentPartIndex+1, j)
					break
				}
			}
		}

	}
	return
}

func partOne(springs []Spring) {
	for _, spring := range springs {
		fmt.Println("Spring", string(spring.Spring), getPermutationNumberOfParts(spring, 0, 0))
		//break
	}
}

func main() {
	var rawData = utils.GetInput("day_12/input.txt")
	var springs = parseData(rawData)
	partOne(springs)
}
