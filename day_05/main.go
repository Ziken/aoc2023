package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"math"
	"regexp"
	"strconv"
)

type RequirementRange struct {
	Source      int
	Destination int
	Length      int
	Multiplier  int
}

type SeedRequirement struct {
	FromDest     string
	ToDest       string
	requirements []RequirementRange
}

func (r *RequirementRange) NextID(seed int) int {
	// Return -1 if not found in range
	if seed >= r.Source && seed < (r.Source+r.Length) {
		return r.Destination + seed - r.Source
	}
	return -1
}

func (s *SeedRequirement) getNextReqID(seed int) int {
	for _, req := range s.requirements {
		var nextID = req.NextID(seed)
		if nextID != -1 {
			return nextID
		}
	}
	return seed
}

func getSeeds(data []byte) (out []int) {
	var reg = regexp.MustCompile(`\d+`)

	for _, rawNum := range reg.FindAll(data, -1) {
		var n, _ = strconv.Atoi(string(rawNum))
		out = append(out, n)
	}
	return
}
func getMap(data [][]byte, startIndex int) (fromDest string, toDest string) {
	var mapRegexp = regexp.MustCompile(`(\w+)-to-(\w+).*`)
	var matched = mapRegexp.FindSubmatch(data[startIndex])
	fromDest, toDest = string(matched[1]), string(matched[2])

	return
}
func getRangeRequirements(data [][]byte, startIndex int) RequirementRange {
	var rangeRegexp = regexp.MustCompile(`(\d+)`)
	var matched = rangeRegexp.FindAllSubmatch(data[startIndex], -1)
	// clean up this indexes
	var destination, source, length = matched[0][0], matched[1][0], matched[2][0]
	var destInt, _ = strconv.Atoi(string(destination))
	var sourceInt, _ = strconv.Atoi(string(source))
	var lengthInt, _ = strconv.Atoi(string(length))

	return RequirementRange{Source: sourceInt, Destination: destInt, Length: lengthInt}

}
func getSeedRequirements(data [][]byte) (out []SeedRequirement) {
	var startIndex = 2

	for startIndex < len(data) {
		var fromDest, toDest = getMap(data, startIndex)
		var seedReq = SeedRequirement{FromDest: fromDest, ToDest: toDest}
		startIndex++
		for startIndex < len(data) && len(data[startIndex]) > 0 {
			seedReq.requirements = append(seedReq.requirements, getRangeRequirements(data, startIndex))
			startIndex++
		}
		out = append(out, seedReq)
		startIndex++
	}
	return
}

func partOne(seeds []int, seedRequirements []SeedRequirement) int {
	var lowestLocation = math.Inf(1)

	for _, seed := range seeds {
		for _, seedReq := range seedRequirements {
			seed = seedReq.getNextReqID(seed)
		}
		if float64(seed) < lowestLocation {
			lowestLocation = float64(seed)
		}
	}
	return int(lowestLocation)
}

func partTwo(seeds []int, seedRequirements []SeedRequirement) int {
	var lowestLocation = math.Inf(1)
	var seedIndex = 0
	for seedIndex < len(seeds) {
		var startSeed, seedLength = seeds[seedIndex], seeds[seedIndex+1]
		fmt.Println("Start seed:", startSeed, "Seed length:", seedLength)
		for s1 := startSeed; s1 < (startSeed + seedLength); s1++ {
			var result = s1
			for _, seedReq := range seedRequirements {
				result = seedReq.getNextReqID(result)
			}

			if float64(result) < lowestLocation {
				lowestLocation = float64(result)
			}
		}

		seedIndex += 2
	}
	return int(lowestLocation)
}

func main() {
	var data = utils.GetInput("day_05/input.txt")
	var seeds = getSeeds(data[0])
	var seedRequirements = getSeedRequirements(data)
	fmt.Println("Part one:", partOne(seeds, seedRequirements))
	fmt.Println("Part Two:", partTwo(seeds, seedRequirements))
}
