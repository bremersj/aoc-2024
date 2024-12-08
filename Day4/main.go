package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func generateColumnStrings(input []string) []string {
	columns := make([]string, len(input[0]))
	for _, line := range input {
		for j, char := range line {
			columns[j] += string(char)
		}
	}
	return columns
}

func generateRightDiagonalStrings(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	rows := len(input)
	cols := len(input[0])
	diagonals := make([]string, rows+cols-1)

	for i, line := range input {
		for j, char := range line {
			diagonals[i+j] += string(char)
		}
	}
	return diagonals
}

func generateLeftDiagonalStrings(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	rows := len(input)
	cols := len(input[0])
	// For left diagonals, index will range from -(cols-1) to (rows-1)
	// So total number of diagonals is still (rows + cols - 1)
	diagonals := make([]string, rows+cols-1)

	for i, line := range input {
		for j, char := range line {
			// We need to offset the index since i-j can be negative
			// Add (cols-1) to shift all indices to be non-negative
			diagIndex := i - j + (cols - 1)
			diagonals[diagIndex] += string(char)
		}
	}
	return diagonals
}

func countPatternInString(s string, pattern string) int {
	count := 0
	// Check for overlapping matches by sliding the window
	for i := 0; i <= len(s)-len(pattern); i++ {
		if s[i:i+len(pattern)] == pattern {
			count++
		}
	}
	return count
}

func countMatches(input []string) int {
	count := 0
	patterns := []string{"XMAS", "SAMX"}

	for _, line := range input {
		for _, pattern := range patterns {
			count += countPatternInString(line, pattern)
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

	count := 0
	count += countMatches(parseInput(string(data)))
	fmt.Println("count: ", count)

	columns := generateColumnStrings(parseInput(string(data)))
	fmt.Println("Columns: ", columns)
	count += countMatches(columns)
	fmt.Println("count: ", count)

	RDiagonals := generateRightDiagonalStrings(parseInput(string(data)))
	fmt.Println("RDiagonals: ", RDiagonals)
	count += countMatches(RDiagonals)
	fmt.Println("count: ", count)

	LDiagonals := generateLeftDiagonalStrings(parseInput(string(data)))
	fmt.Println("LDiagonals: ", LDiagonals)
	count += countMatches(LDiagonals)

	fmt.Println("Final Count: ", count)
}
