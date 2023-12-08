package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
)

type NodeCycle struct {
	offset           int
	cycleLength      int
	currentStepCount int
}

func parseInput(data [][]byte) (map[string][]string, []int) {
	var directionRegexp = regexp.MustCompile(`\w{3}`)
	var network = make(map[string][]string)
	var directions = make([]int, 0)
	for _, d := range data[0] {
		if d == 'L' {
			directions = append(directions, 0)
		} else {
			directions = append(directions, 1)
		}
	}

	for _, row := range data[2:] {
		var matches = directionRegexp.FindAll(row, -1)
		var from, left, right = string(matches[0]), string(matches[1]), string(matches[2])
		//network[from] = append(network[from], left, right)
		network[from] = []string{left, right}
	}

	return network, directions
}

func partOne(network map[string][]string, directions []int) (stepsCount int) {
	var currentNode = "AAA"
	for {
		for _, direction := range directions {
			currentNode = network[currentNode][direction]
			stepsCount++
			if currentNode == "ZZZ" {
				return
			}
		}
	}
}

func findCycle(network map[string][]string, directions []int, startNode string) (offset, cycle int) {
	var currentNode = startNode
	var foundEndNode = false
	for {
		for _, direction := range directions {
			currentNode = network[currentNode][direction]
			if foundEndNode && currentNode[2] == 'Z' {
				return offset, cycle
			}

			if currentNode[2] == 'Z' {
				foundEndNode = true
			}
			if foundEndNode {
				cycle++
			} else {
				offset++
			}
		}
	}
}

func partTwo(network map[string][]string, directions []int) int {
	var startingNodes = []string{}

	var startingNodeCycle = make([]NodeCycle, 0)
	for key := range network {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}

	for _, startNode := range startingNodes {
		var offset, cycleLength = findCycle(network, directions, startNode)
		startingNodeCycle = append(startingNodeCycle, NodeCycle{offset, cycleLength, offset})
	}

	for {
		var maxStepCount = 0
		for i := range startingNodeCycle {
			if startingNodeCycle[i].currentStepCount > maxStepCount {
				maxStepCount = startingNodeCycle[i].currentStepCount
			}
		}

		for i := range startingNodeCycle {
			for startingNodeCycle[i].currentStepCount < maxStepCount {
				startingNodeCycle[i].currentStepCount += startingNodeCycle[i].cycleLength
			}
		}
		var areAllEqual = true
		for i := range startingNodeCycle[1:] {
			if startingNodeCycle[i].currentStepCount != startingNodeCycle[i+1].currentStepCount {
				areAllEqual = false
				break
			}
		}
		if areAllEqual {
			break
		} else {
			// increment all by cycle length to optimize computations (for finding maxStepCount)
			for i := range startingNodeCycle {
				startingNodeCycle[i].currentStepCount += startingNodeCycle[i].cycleLength
			}
		}
	}

	return startingNodeCycle[0].currentStepCount + 1
}
func main() {
	var data = utils.GetInput("day_08/input.txt")
	var network, directions = parseInput(data)
	fmt.Println("Part one:", partOne(network, directions))
	fmt.Println("Part Two:", partTwo(network, directions))
}
