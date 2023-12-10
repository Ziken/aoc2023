package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
)

const (
	GROUND          = '.'
	VERTICAL_PIPE   = '|'
	HORIZONTAL_PIPE = '-'
	NORTH_EAST_PIPE = 'L'
	NORTH_WEST_PIPE = 'J'
	SOUTH_EAST_PIPE = 'F'
	SOUTH_WEST_PIPE = '7'
	STARTING_POINT  = 'S'
)

type Coordinates struct {
	X, Y int
}

func getStartCoords(data [][]byte) (int, int) {
	for y, row := range data {
		for x, char := range row {
			if char == STARTING_POINT {
				return x, y
			}
		}
	}
	panic("No start found")
}
func getNextStep(data [][]byte, startX, startY int) (int, int) {
	if p := data[startY+1][startX]; p == VERTICAL_PIPE || p == NORTH_EAST_PIPE {
		return startX, startY + 1
	}
	if p := data[startY][startX-1]; p == HORIZONTAL_PIPE || p == NORTH_EAST_PIPE || p == SOUTH_EAST_PIPE {
		return startX - 1, startY
	}
	if p := data[startY-1][startX]; p == VERTICAL_PIPE || p == SOUTH_WEST_PIPE || p == SOUTH_EAST_PIPE {
		return startX, startY - 1
	}
	if p := data[startY][startX+1]; p == HORIZONTAL_PIPE || p == NORTH_WEST_PIPE || p == SOUTH_WEST_PIPE {
		return startX + 1, startY
	}
	panic("No next step found")
}

func getNextPosition(data [][]byte, currentX, currentY, prevX, prevY int, prevDirection Coordinates) (bool, int, int, Coordinates) {
	var direction = Coordinates{}
	if currentX < 0 || currentY < 0 || currentX >= len(data[0]) || currentY >= len(data) {
		return false, 0, 0, direction
	}
	var directionTurn = '0'
	switch data[currentY][currentX] {
	case GROUND:
		return false, 0, 0, direction
	case VERTICAL_PIPE:
		if currentY > prevY {
			// down
			currentY++
		} else {
			// up
			currentY--
		}
		direction = prevDirection
	case HORIZONTAL_PIPE:
		if currentX > prevX {
			// right
			currentX++
		} else {
			// left
			currentX--
		}
		direction = prevDirection
	case NORTH_EAST_PIPE:
		if currentY > prevY {
			currentX++
			directionTurn = 'L'
		} else if currentX < prevX {
			currentY--
			directionTurn = 'R'
		} else {
			return false, 0, 0, direction
		}
	case NORTH_WEST_PIPE:
		if currentX > prevX {
			currentY--
			directionTurn = 'L'
		} else if currentY > prevY {
			currentX--
			directionTurn = 'R'
		} else {
			return false, 0, 0, direction
		}
	case SOUTH_WEST_PIPE:
		if currentY < prevY {
			currentX--
			directionTurn = 'L'
		} else if currentX > prevX {
			currentY++
			directionTurn = 'R'
		} else {
			return false, 0, 0, direction
		}
	case SOUTH_EAST_PIPE:
		if currentX < prevX {
			currentY++
			directionTurn = 'L'
		} else if currentY < prevY {
			currentX++
			directionTurn = 'R'
		} else {
			return false, 0, 0, direction
		}
	case STARTING_POINT:
		return false, 0, 0, direction
	}
	if directionTurn == 'L' {
		switch prevDirection {
		case Coordinates{0, -1}:
			direction = Coordinates{-1, 0}
		case Coordinates{1, 0}:
			direction = Coordinates{0, -1}
		case Coordinates{0, 1}:
			direction = Coordinates{1, 0}
		case Coordinates{-1, 0}:
			direction = Coordinates{0, 1}
		}
	} else if directionTurn == 'R' {
		switch prevDirection {
		case Coordinates{0, -1}:
			direction = Coordinates{1, 0}
		case Coordinates{1, 0}:
			direction = Coordinates{0, 1}
		case Coordinates{0, 1}:
			direction = Coordinates{-1, 0}
		case Coordinates{-1, 0}:
			direction = Coordinates{0, -1}
		}
	}

	return true, currentX, currentY, direction
}
func traversePath(data [][]byte) map[Coordinates]bool {
	var startX, startY = getStartCoords(data)
	var traversedPath = make(map[Coordinates]bool)
	traversedPath[Coordinates{startX, startY}] = true
	var currentX, currentY = getNextStep(data, startX, startY)
	traversedPath[Coordinates{currentX, currentY}] = true
	var prevX, prevY = startX, startY
	var isValid = false
	for {
		var cx, xy = currentX, currentY
		isValid, currentX, currentY, _ = getNextPosition(data, currentX, currentY, prevX, prevY, Coordinates{0, 0})
		if !isValid {
			break
		}
		traversedPath[Coordinates{currentX, currentY}] = true
		prevX, prevY = cx, xy
	}
	return traversedPath

}

func partOne(traversedPath map[Coordinates]bool) (out int) {
	return len(traversedPath) / 2
}
func getNextDirection(direction Coordinates) Coordinates {
	var nextD = Coordinates{}
	switch direction {
	case Coordinates{1, 0}:
		nextD = Coordinates{0, -1}
	case Coordinates{0, -1}:
		nextD = Coordinates{-1, 0}
	case Coordinates{-1, 0}:
		nextD = Coordinates{0, 1}
	case Coordinates{0, 1}:
		nextD = Coordinates{1, 0}
	}
	return nextD
}
func getInnerArea(traversedPath map[Coordinates]bool, data [][]byte, originalDirection Coordinates) map[Coordinates]bool {
	var startX, startY = getStartCoords(data)
	var innerPoints = make(map[Coordinates]bool)
	var currentX, currentY = getNextStep(data, startX, startY)
	var direction = originalDirection
	var prevX, prevY = startX, startY
	var isValid = false
	for {
		var cx, cy = currentX, currentY
		isValid, currentX, currentY, direction = getNextPosition(data, currentX, currentY, prevX, prevY, direction)
		if !isValid {
			break
		}
		var areaPoint = Coordinates{currentX + direction.X, currentY + direction.Y}
		ok1, _ := traversedPath[areaPoint]
		ok2, _ := innerPoints[areaPoint]
		if !ok1 && !ok2 {
			var isValidArea, visitedPoints = traverseArea(data, traversedPath, innerPoints, areaPoint)
			if !isValidArea {
				var nextD = getNextDirection(originalDirection)
				return getInnerArea(traversedPath, data, nextD)
			}
			for point := range visitedPoints {
				innerPoints[point] = true
			}
		}

		// --- special case, when there is curve and area to check
		areaPoint = Coordinates{cx + direction.X, cy + direction.Y}
		ok1, _ = traversedPath[areaPoint]
		ok2, _ = innerPoints[areaPoint]
		if !ok1 && !ok2 {
			var isValidArea, visitedPoints = traverseArea(data, traversedPath, innerPoints, areaPoint)
			if !isValidArea {
				var nextD = getNextDirection(originalDirection)
				return getInnerArea(traversedPath, data, nextD)
			}
			for point := range visitedPoints {
				innerPoints[point] = true
			}
		}
		// ---
		prevX, prevY = cx, cy
	}
	if len(innerPoints) == 0 {
		return getInnerArea(traversedPath, data, getNextDirection(originalDirection))
	}
	return innerPoints
}
func partTwo(traversedPath map[Coordinates]bool, data [][]byte) (out int) {
	var direction = Coordinates{0, 1}
	var innerPoints = getInnerArea(traversedPath, data, direction)
	return len(innerPoints)
}
func traverseArea(data [][]byte, traversedPath map[Coordinates]bool, innerPoints map[Coordinates]bool, areaPoint Coordinates) (bool, map[Coordinates]bool) {
	var visitedPoints = make(map[Coordinates]bool)
	var directions = []Coordinates{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visitedPoints[areaPoint] = true
	var nextPoints = []Coordinates{areaPoint}
	for len(nextPoints) > 0 {
		var possiblePoints = []Coordinates{}
		for _, point := range nextPoints {
			for _, direction := range directions {
				var nextPoint = Coordinates{point.X + direction.X, point.Y + direction.Y}
				if nextPoint.X < 0 || nextPoint.Y < 0 || nextPoint.X >= len(data[0]) || nextPoint.Y >= len(data) {
					return false, visitedPoints
				}
				if ok, _ := visitedPoints[nextPoint]; ok {
					continue
				}
				if ok, _ := innerPoints[nextPoint]; ok {
					continue
				}
				if ok, _ := traversedPath[nextPoint]; !ok {
					visitedPoints[nextPoint] = true
					possiblePoints = append(possiblePoints, nextPoint)
				}
			}
		}
		nextPoints = possiblePoints
	}
	return true, visitedPoints
}

func main() {
	var data = utils.GetInput("day_10/input.txt")
	var traversedPath = traversePath(data)
	fmt.Println("Part one:", partOne(traversedPath))
	fmt.Println("Part two:", partTwo(traversedPath, data))
}

// 487 -- too low
