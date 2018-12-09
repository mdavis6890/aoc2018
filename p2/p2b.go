package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func bucketLetters(s string) map[rune]int {
	var letters map[rune]int
	letters = make(map[rune]int)

	for _, i := range s {
		letters[i] = letters[i] + 1
		// fmt.Println("%c", i)
	}
	return letters
}

func letterCount(l map[rune]int, c int) bool {
	for _, i := range l {
		if i == c {
			return true
		}
	}
	return false
}

func match(s1 string, s2 string) (bool, string) {
	misses := 0
	isMatch := true
	misMatchIdx := 0
	// convert strings to rune slices
	r1 := []rune(s1)
	r2 := []rune(s2)

	for i, l := range r1 {
		if l != r2[i] {
			misses++
			misMatchIdx = i
		}
		if misses > 1 {
			isMatch = false
			break
		}
	}
	commonString := s1[:misMatchIdx] + s2[misMatchIdx+1:]
	return isMatch, commonString
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func testMatches(s string, l []string) (bool, string) {
	for _, i := range l {
		isMatch, commonString := match(s, i)
		if isMatch {
			return true, commonString
		}
	}
	return false, ""
}

func main() {

	input, err := readLines("p2/p2-input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var commonString string
	var foundMatch bool

	for i, line := range input {
		foundMatch, commonString = testMatches(line, input[i+1:])
		if foundMatch {
			//s1 = line
			break
		}
	}
	fmt.Println(commonString)
}
