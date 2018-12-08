package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func bucketLetters (s string) map[rune]int {
	var letters map[rune]int
	letters = make(map[rune]int)

	for _, i := range s {
		letters[i] = letters[i] + 1
		// fmt.Println("%c", i)
	}
	return letters
}

func letterCount (l map[rune]int, c int) bool {
	for _, i := range l {
		if i == c {
			return true
		}
	}
	return false
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

func main () {

	input, err := readLines("p2-input.txt")
	if err != nil {
    	log.Fatalf("readLines: %s", err)
    }
	count2 := 0
	count3 := 0

	for _, line := range input {

		b := bucketLetters(line)
		m2 := letterCount(b, 2)
		m3 := letterCount(b, 3)
		fmt.Println(line)
		fmt.Printf("2: %t\n", m2)
		fmt.Printf("3: %t\n", m3)
		if m2 {count2++}
		if m3 {count3++}
	}
	fmt.Printf("Count 2: %d\n", count2)
	fmt.Printf("Count 3: %d\n", count3)
	fmt.Printf("CheckSum: %d\n", count2 * count3)
}