package main

import (
	"bytes"
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"math"
	"strconv"
)

type Lens struct {
	Value string
	Order int
}

func hashString(s []byte) (current int) {
	for _, char := range s {
		current += int(char)
		current *= 17
		current %= 256
	}

	return
}

func partOne(splittedData [][]byte) (out int) {
	for _, part := range splittedData {
		out = out + hashString(part)
	}
	return
}
func partTwo(splittedData [][]byte) (out int) {
	var boxes = make([]map[string]Lens, 257)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = make(map[string]Lens)
	}

	for i, part := range splittedData {
		if p := bytes.Split(part, []byte("=")); len(p) == 2 {
			var h = hashString(p[0])
			var key = string(p[0])
			if _, ok := boxes[h][key]; !ok {
				boxes[h][key] = Lens{Value: string(p[1]), Order: i}
			} else {
				boxes[h][key] = Lens{Value: string(p[1]), Order: boxes[h][key].Order}
			}
		} else {
			var p = bytes.Split(part, []byte("-"))
			var h = hashString(p[0])
			delete(boxes[h], string(p[0]))
		}

	}

	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		var order = 1
		for len(box) > 0 {
			var smallestOrder = math.MaxInt64
			var smallestKey string
			for key, lensObj := range box {
				if lensObj.Order < smallestOrder {
					smallestOrder = lensObj.Order
					smallestKey = key
				}
			}
			var focalLength, _ = strconv.Atoi(box[smallestKey].Value)
			out += (i + 1) * order * focalLength
			order++
			delete(box, smallestKey)
		}
	}
	return
}
func main() {
	var data = utils.GetInput("day_15/input.txt")
	var splittedData = bytes.Split(data[0], []byte(","))
	fmt.Println("Part One", partOne(splittedData))
	fmt.Println("Part Two", partTwo(splittedData))
}
