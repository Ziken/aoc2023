package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
)

func parseData(data [][]byte) (times, distances []int) {
	var numRegexp = regexp.MustCompile(`\d+`)
	var rawTimes = numRegexp.FindAll(data[0], -1)
	var rawDistances = numRegexp.FindAll(data[1], -1)

	for i := range rawTimes {
		var t, _ = strconv.Atoi(string(rawTimes[i]))
		var d, _ = strconv.Atoi(string(rawDistances[i]))
		times = append(times, t)
		distances = append(distances, d)
	}
	return
}

func countWinningConditions(time, distance int) (winConditions int) {
	for speed := 1; speed < time; speed++ {
		if (time-speed)*speed > distance {
			winConditions += 1
		}
	}
	return
}

func partOne(times, distances []int) (out int) {
	out = 1

	for i := range times {
		var time = times[i]
		var distance = distances[i]
		out *= countWinningConditions(time, distance)
	}
	return
}
func partTwo(times, distances []int) (out int) {
	var totalRawTime = ""
	var totalRawDistance = ""
	for i := range times {
		totalRawTime += strconv.Itoa(times[i])
		totalRawDistance += strconv.Itoa(distances[i])
	}
	time, _ := strconv.Atoi(totalRawTime)
	distance, _ := strconv.Atoi(totalRawDistance)

	return countWinningConditions(time, distance)
}

func main() {
	var data = utils.GetInput("day_06/input.txt")
	var times, distances = parseData(data)
	fmt.Println("Part one:", partOne(times, distances))
	fmt.Println("Part Two:", partTwo(times, distances))
}
