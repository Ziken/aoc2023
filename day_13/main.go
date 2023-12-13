package main

import (
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
)

type Mirror struct {
	Surface [][]byte
}

const (
	Invalid    = -1
	Vertical   = 0
	Horizontal = 1
)

func parseInput(data [][]byte) (mirrors []Mirror) {
	var surface = [][]byte{}
	for _, row := range data {
		if len(row) == 0 {
			mirrors = append(mirrors, Mirror{Surface: surface})
			surface = [][]byte{}
		} else {
			surface = append(surface, row)
		}
	}
	mirrors = append(mirrors, Mirror{Surface: surface})

	return
}

func getVerticalMirrorReflection(mirror Mirror, ignoreMirrorIndex int) int {
	for reflectionStartIndex := 1; reflectionStartIndex < len(mirror.Surface[0]); reflectionStartIndex++ {
		var leftI, rightI = reflectionStartIndex - 1, reflectionStartIndex
		var isMatch = true

		for i := 0; i < len(mirror.Surface) && isMatch; i++ {
			var li, ri = leftI, rightI
			for li >= 0 && ri < len(mirror.Surface[0]) {
				if mirror.Surface[i][li] != mirror.Surface[i][ri] {
					isMatch = false
					break
				}
				li--
				ri++
			}
		}

		if isMatch && ignoreMirrorIndex != (reflectionStartIndex-1) {
			return reflectionStartIndex - 1
		}
	}
	return -1
}

func getHorizontalMirrorReflection(mirror Mirror, ignoreMirrorIndex int) int {
	for reflectionStartIndex := 1; reflectionStartIndex < len(mirror.Surface); reflectionStartIndex++ {
		var topI, bottomI = reflectionStartIndex - 1, reflectionStartIndex
		var isMatch = true

		for j := 0; j < len(mirror.Surface[0]) && isMatch; j++ {
			var ti, bi = topI, bottomI
			for ti >= 0 && bi < len(mirror.Surface) {
				if mirror.Surface[ti][j] != mirror.Surface[bi][j] {
					isMatch = false
					break
				}
				ti--
				bi++
			}
		}

		if isMatch && ignoreMirrorIndex != (reflectionStartIndex-1) {
			return reflectionStartIndex - 1
		}
	}
	return -1
}

func partOne(mirrors []Mirror) (out int) {
	for _, mirror := range mirrors {
		mirrorIndex, direction := getReflection(mirror, -1, Invalid)
		out += calculateResult(mirrorIndex, direction)
	}
	return
}

func getReverseMirrorPart(part byte) byte {
	if part == '.' {
		return '#'
	}
	return '.'
}
func calculateResult(mirrorIndex, direction int) (out int) {
	if direction == Vertical {
		out += mirrorIndex + 1
	} else {
		out += (mirrorIndex + 1) * 100
	}
	return
}
func getReflection(mirror Mirror, ignoreIndex int, ignoreDirection int) (index, direction int) {
	var verticalReflection, horizontalReflection int
	if ignoreDirection == Horizontal {
		verticalReflection = getVerticalMirrorReflection(mirror, -1)
		horizontalReflection = getHorizontalMirrorReflection(mirror, ignoreIndex)
	} else {
		verticalReflection = getVerticalMirrorReflection(mirror, ignoreIndex)
		horizontalReflection = getHorizontalMirrorReflection(mirror, -1)
	}

	if ignoreDirection == Vertical {
		if horizontalReflection != -1 {
			return horizontalReflection, Horizontal
		}
		if verticalReflection != -1 {
			return verticalReflection, Vertical
		}
	} else {
		if verticalReflection != -1 {
			return verticalReflection, Vertical
		}
		if horizontalReflection != -1 {
			return horizontalReflection, Horizontal
		}
	}

	return -1, Invalid
}

func partTwo(mirrors []Mirror) (out int) {
	for _, mirror := range mirrors {
		var originalMirrorIndex, originalDirection = getReflection(mirror, -1, Invalid)
		var foundNewDirection = false
		for i := 0; i < len(mirror.Surface) && !foundNewDirection; i++ {
			for j := 0; j < len(mirror.Surface[0]); j++ {
				mirror.Surface[i][j] = getReverseMirrorPart(mirror.Surface[i][j])
				var newMirrorIndex, newDirection = getReflection(mirror, originalMirrorIndex, originalDirection)
				if newDirection != Invalid && (newMirrorIndex != originalMirrorIndex || newDirection != originalDirection) {
					out += calculateResult(newMirrorIndex, newDirection)
					foundNewDirection = true
					// revert change
					mirror.Surface[i][j] = getReverseMirrorPart(mirror.Surface[i][j])
					break
				}
				// revert change
				mirror.Surface[i][j] = getReverseMirrorPart(mirror.Surface[i][j])
			}
		}
	}
	return
}
func main() {
	var data = utils.GetInput("day_13/input.txt")
	var mirrors = parseInput(data)
	fmt.Println("Part one:", partOne(mirrors))
	fmt.Println("Part two:", partTwo(mirrors))
}
