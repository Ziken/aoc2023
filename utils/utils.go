package utils

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetInput(inputPath string) (in [][]byte) {
	file, errFile := os.Open(inputPath)
	check(errFile)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		row := scanner.Text()
		in = append(in, []byte(row))
	}

	return
}

func ReverseString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)
	// Get the length of the slice
	n := len(runes)

	// Swap the runes from the ends towards the center
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}
func ReverseBytes(s []byte) []byte {
	// Convert string to a slice of runes
	// Get the length of the slice
	n := len(s)

	// Swap the runes from the ends towards the center
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}

	// Convert the slice of runes back to a string
	return s
}
