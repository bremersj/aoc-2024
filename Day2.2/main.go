package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	levels := make([][]int, len(lines))

	// Process each line
	for i, line := range lines {
		numbers := strings.Fields(line)
		levels[i] = make([]int, len(numbers))

		for j, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil
			}
			levels[i][j] = num
		}
	}

	return levels
}

const (
	Up int = iota
	Down
	Equal
)

func getDirection(a, b int) int {
	if a < b {
		return Up
	} else if a > b {
		return Down
	} else {
		return Equal
	}
}

func isSameDirection(a, b int, state int) bool {
	if a < b && state == Up {
		return true
	} else if a > b && state == Down {
		return true
	} else {
		return false
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isSafeDistance(a, b int) bool {
	return abs(a-b) <= 3
}

func isSafeLevel(level []int) bool {
	state := getDirection(level[0], level[1])

	if state == Equal {
		return false
	}

	for i := 0; i < len(level); i++ {
		num := level[i]

		// if we get to the last number, it is safe
		if i == len(level)-1 {
			return true
		}

		if !(isSameDirection(num, level[i+1], state) && isSafeDistance(num, level[i+1])) {
			return false
		}
	}

	return false
}

func countSafeLevels(levels [][]int) int {
	count := 0

	for _, level := range levels {
		if isSafeLevel(level) {
			count++
			continue
		} else {
			for i := 0; i < len(level); i++ {
				// Create new level without element at index i
				newLevel := []int{}
				newLevel = append(newLevel, level[:i]...)
				newLevel = append(newLevel, level[i+1:]...)

				fmt.Printf("New level: %v\n", newLevel)

				if isSafeLevel(newLevel) {
					count++
					break
				}
			}
		}
	}
	return count
}

func main() {
	testFlag := flag.Bool("t", false, "Use test input file")
	testFlagLong := flag.Bool("test", false, "Use test input file")
	flag.Parse()

	inputFile := "input.txt"
	if *testFlag || *testFlagLong {
		inputFile = "test_input.txt"
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	levels := parseInput(string(data))

	safeCount := countSafeLevels(levels)
	fmt.Printf("Safe count: %d\n", safeCount)
}
