package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

func main() {
	fabric := [][]int{}

	input, err := readLines("p3/p3-input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	r, _ := regexp.Compile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for _, line := range input {
		groups := r.FindStringSubmatch(line)
		fmt.Printf("Matches: %s\n", groups)
		claimId, _ := strconv.ParseInt(groups[1], 10, 64)
		xOffset, _ := strconv.ParseInt(groups[2], 10, 64)
		yOffset, _ := strconv.ParseInt(groups[3], 10, 64)
		width, _ := strconv.ParseInt(groups[4], 10, 64)
		height, _ := strconv.ParseInt(groups[5], 10, 64)
		fmt.Printf("Claim ID: %d\n", claimId)
		fmt.Printf("xOffset: %d\n", xOffset)
		fmt.Printf("yOffset: %d\n", yOffset)
		fmt.Printf("width: %d\n", width)
		fmt.Printf("height: %d\n", height)

		fabric[xOffset][yOffset]++
	}
	fmt.Println(fabric)
}
