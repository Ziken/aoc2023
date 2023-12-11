package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"math"
)

type Galaxy struct {
	X, Y int
}

func getGalaxies(data [][]byte, distanceMultiplier int) []Galaxy {
	var galaxies []Galaxy
	var columOffset = []int{0}

	for x := 0; x < len(data[0]); x++ {
		var isColEmpty = true
		for y := 0; y < len(data); y++ {
			if data[y][x] == '#' {
				isColEmpty = false
			}
		}
		var lastOffset = columOffset[len(columOffset)-1]
		if isColEmpty {
			columOffset = append(columOffset, lastOffset+distanceMultiplier)
		} else {
			columOffset = append(columOffset, lastOffset)
		}
	}
	// remove first mock element
	columOffset = columOffset[1:]

	var rowOffset = 0
	for y := 0; y < len(data); y++ {
		var isEmptyRow = true
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '#' {
				galaxies = append(galaxies, Galaxy{X: x + columOffset[x], Y: y + rowOffset})
				isEmptyRow = false
			}
		}
		if isEmptyRow {
			rowOffset += distanceMultiplier
		}
	}

	return galaxies
}

func getDistance(galaxies []Galaxy) (out int) {
	for i, galaxy := range galaxies {
		for _, otherGalaxy := range galaxies[i+1:] {
			var d = int(math.Abs(float64(galaxy.X-otherGalaxy.X)) + math.Abs(float64(galaxy.Y-otherGalaxy.Y)))
			out += d
		}
	}
	return out
}

func main() {
	var data = utils.GetInput("day_11/input.txt")
	var galaxiesPartOne = getGalaxies(data, 1)
	var galaxiesPartTwo = getGalaxies(data, 1000000-1)
	fmt.Println("Part one:", getDistance(galaxiesPartOne))
	fmt.Println("Part two:", getDistance(galaxiesPartTwo))
}
