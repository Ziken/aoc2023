package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
)

const (
	MIRROR_SLASH              = '/'
	MIRROR_BACKSLASH          = '\\'
	MIRROR_REFLECT_VERTICAL   = '|'
	MIRROR_REFLECT_HORIZONTAL = '-'
)
const (
	DIRECTION_UP = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

type VisitedPoint struct {
	X, Y, Direction int
}

func followLightPath(data [][]byte, visitedPoints map[VisitedPoint]bool, currentX, currentY int, direction int) {
	for currentX >= 0 && currentX < len(data[0]) && currentY >= 0 && currentY < len(data) {

		if visitedPoints[VisitedPoint{X: currentX, Y: currentY, Direction: direction}] {
			// found a loop
			return
		}
		visitedPoints[VisitedPoint{X: currentX, Y: currentY, Direction: direction}] = true
		switch data[currentY][currentX] {
		case MIRROR_SLASH: // /
			switch direction {
			case DIRECTION_UP:
				direction = DIRECTION_RIGHT
			case DIRECTION_DOWN:
				direction = DIRECTION_LEFT
			case DIRECTION_LEFT:
				direction = DIRECTION_DOWN
			case DIRECTION_RIGHT:
				direction = DIRECTION_UP
			}
		case MIRROR_BACKSLASH: // \
			switch direction {
			case DIRECTION_UP:
				direction = DIRECTION_LEFT
			case DIRECTION_DOWN:
				direction = DIRECTION_RIGHT
			case DIRECTION_LEFT:
				direction = DIRECTION_UP
			case DIRECTION_RIGHT:
				direction = DIRECTION_DOWN
			}
		case MIRROR_REFLECT_VERTICAL: // |
			if direction == DIRECTION_LEFT || direction == DIRECTION_RIGHT {
				followLightPath(data, visitedPoints, currentX, currentY-1, DIRECTION_UP)
				followLightPath(data, visitedPoints, currentX, currentY+1, DIRECTION_DOWN)
				return
			}
		case MIRROR_REFLECT_HORIZONTAL: // -
			if direction == DIRECTION_UP || direction == DIRECTION_DOWN {
				followLightPath(data, visitedPoints, currentX-1, currentY, DIRECTION_LEFT)
				followLightPath(data, visitedPoints, currentX+1, currentY, DIRECTION_RIGHT)
				return
			}
		}

		switch direction {
		case DIRECTION_UP:
			currentY--
		case DIRECTION_DOWN:
			currentY++
		case DIRECTION_LEFT:
			currentX--
		case DIRECTION_RIGHT:
			currentX++
		}
	}
}

func partOne(data [][]byte) int {
	var visitedPoints = make(map[VisitedPoint]bool)
	followLightPath(data, visitedPoints, 53, 31, 3)
	var uniquePoints = make(map[VisitedPoint]bool)
	for point := range visitedPoints {
		uniquePoints[VisitedPoint{X: point.X, Y: point.Y}] = true
	}
	return len(uniquePoints)
}

func partTwo(data [][]byte) int {
	var bestResult = 0
	var globalVisitedPoints = make(map[VisitedPoint]bool)
	for i := 0; i < len(data); i++ {
		var points = [][]int{
			{0, i},
			{i, 0},
			{len(data[0]) - 1, 0},
			{0, len(data) - 1},
		}
		for _, p := range points {
			var x, y = p[0], p[1]
			for direction := 0; direction < 4; direction++ {
				if globalVisitedPoints[VisitedPoint{X: x, Y: y, Direction: direction}] {
					continue
				}
				var visitedPoints = make(map[VisitedPoint]bool)
				followLightPath(data, visitedPoints, x, y, direction)
				var uniquePoints = make(map[VisitedPoint]bool)
				for point := range visitedPoints {
					uniquePoints[VisitedPoint{X: point.X, Y: point.Y}] = true
				}
				if len(uniquePoints) > bestResult {
					bestResult = len(uniquePoints)
				}
				for point := range visitedPoints {
					globalVisitedPoints[point] = true
				}
			}
		}
	}

	return bestResult
}
func main() {
	var data = utils.GetInput("day_16/input.txt")
	fmt.Println("Part one", partOne(data))
	fmt.Println("Part two", partTwo(data))
}
