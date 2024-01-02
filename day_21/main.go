package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
)

type Point struct {
	X, Y int
}

func findStart(data [][]byte) Point {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == 'S' {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{}
}

func getNextPossiblePoints(data [][]byte, current Point) []Point {
	var directions = []Point{{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}}
	var possiblePoints []Point
	for _, dir := range directions {
		var next = Point{X: current.X + dir.X, Y: current.Y + dir.Y}
		if next.X < 0 || next.Y < 0 || next.X >= len(data[0]) || next.Y >= len(data) {
			continue
		}
		possiblePoints = append(possiblePoints, next)
	}
	return possiblePoints
}

func partOne(data [][]byte) (out int) {
	var start = findStart(data)
	var visited = make(map[Point]bool)
	var queue = []Point{start}
	var steps = 0
	var leftUpCorner = Point{X: len(data[0]) - 1, Y: len(data) - 1}
	var rightDownCorner = Point{X: 0, Y: 0}
	for ; len(queue) > 0 && steps < 1+64; steps++ {
		var nextQueue []Point
		for _, current := range queue {
			queue = queue[1:]
			if visited[current] {
				continue
			}
			visited[current] = true
			if current.X < leftUpCorner.X {
				leftUpCorner.X = current.X
			}
			if current.Y < leftUpCorner.Y {
				leftUpCorner.Y = current.Y
			}
			if current.X > rightDownCorner.X {
				rightDownCorner.X = current.X
			}
			if current.Y > rightDownCorner.Y {
				rightDownCorner.Y = current.Y
			}
			for _, next := range getNextPossiblePoints(data, current) {
				if data[next.Y][next.X] == '#' {
					continue
				}
				nextQueue = append(nextQueue, next)
			}
		}
		queue = nextQueue
		nextQueue = []Point{}
	}
	fmt.Println(leftUpCorner, rightDownCorner)

	var checkStarts = []Point{start, {X: start.X - 1, Y: start.Y}, {X: start.X, Y: start.Y - 1}}
	var bestGardenPlots = make(map[Point]bool)

	for _, checkStart := range checkStarts {
		var gardenPlots = make(map[Point]bool)
		var startX = leftUpCorner.X
		var startY = leftUpCorner.Y
		if (checkStart.X-leftUpCorner.X+1)%2 != 0 {
			startX -= 1
		}
		if (checkStart.Y-leftUpCorner.Y+1)%2 != 0 {
			startY -= 1
		}
		for y := startY; y <= rightDownCorner.Y; y++ {
			var changeX = (y - startY) % 2
			for x := startX + changeX; x <= rightDownCorner.X; x += 2 {
				var p = Point{X: x, Y: y}
				if _, ok := visited[p]; ok {
					gardenPlots[Point{X: x, Y: y}] = true
				}
			}
		}
		if len(gardenPlots) > len(bestGardenPlots) {
			fmt.Println("Found better", startX, startY)
			bestGardenPlots = gardenPlots
		}
	}

	for y, row := range data {
		for x, cell := range row {
			if bestGardenPlots[Point{X: x, Y: y}] {
				fmt.Print("G")
			} else if visited[Point{X: x, Y: y}] {
				fmt.Print("O")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}

	return len(bestGardenPlots)
}

func main() {
	var data = utils.GetInput("day_21/input.txt")
	fmt.Println("Part one", partOne(data))
}
