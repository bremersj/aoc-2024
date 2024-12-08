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

func countSafeLevels(levels [][]int) int {
	count := 0

	for _, level := range levels {
		state := getDirection(level[0], level[1])

		if state == Equal {
			continue
		}

		for i, num := range level {
			// if we get to the last number, then it is safe
			if i == len(level)-1 {
				count++
				continue
			}

			if !isSameDirection(num, level[i+1], state) {
				break
			}

			if !isSafeDistance(num, level[i+1]) {
				break
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
