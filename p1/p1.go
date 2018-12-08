package main

import (
	"bufio"
	"fmt"
	"os"
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

func inSlice(s []int64, val int64) bool {
	println("Checking if val in slice")

	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	freq := []int64{0}
	found := false
	lines, err := readLines("p1/p1-input.txt")
	if err != nil {
		fmt.Print(err)
	}
	for found == false {
		println("Starting Over...")
		for _, s := range lines {
			currFreq := freq[len(freq)-1]
			//fmt.Println(i, s)
			value, err := strconv.ParseInt(s, 0, 0)
			if err != nil {
				panic(err)
			}
			resFreq := value + currFreq
			if inSlice(freq, resFreq) {
				found = true
			}
			freq = append(freq, resFreq)
			fmt.Println("Res Freq: ", freq[len(freq)-1])
			if found == true {
				break
			}
		}
	}
}
