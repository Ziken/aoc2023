package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
)

type Rock struct {
	X, Y int
}

const (
	ROUNDED_ROCK     = 'O'
	CUBE_SHAPED_ROCK = '#'
	EMPTY_SPACE      = '.'
)

func partOne(data [][]byte) (out int) {
	var rocks []Rock
	for x := 0; x < len(data[0]); x++ {
		var posX, posY = x, 0
		for y := 0; y < len(data); y++ {
			if data[y][x] == ROUNDED_ROCK {
				rocks = append(rocks, Rock{X: posX, Y: posY})
				posY++
				out += len(data) - posY + 1
			} else if data[y][x] == CUBE_SHAPED_ROCK {
				posY = y + 1
			}
		}
	}
	return
}

func partTwo(data [][]byte) (out int) {
	return
}

func main() {
	var data = utils.GetInput("day_14/input.txt")
	fmt.Println("Part one:", partOne(data))
}
