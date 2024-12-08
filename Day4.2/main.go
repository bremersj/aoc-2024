package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	if rows == 0 {
		return nil
	}
	cells := make([][]rune, rows)

	for i, line := range lines {
		cells[i] = []rune(line)
	}

	return cells
}

func isValidCross(window [3][3]rune) bool {
	// if middle is an A, then make other checks
	if window[1][1] == 'A' {
		// now check if 0,0 is M and 2,2 is A or 0,0 is M and 2,2 is A
		if ((window[0][0] == 'M' && window[2][2] == 'S') || (window[0][0] == 'S' && window[2][2] == 'M')) &&
			((window[0][2] == 'M' && window[2][0] == 'S') || (window[0][2] == 'S' && window[2][0] == 'M')) {
			return true
		}
	}
	return false
}

// move 3x3 window over canvas and check from cross mas match
func countCrosses(canvas [][]rune) int {
	total := 0
	for i := 0; i < len(canvas)-2; i++ {
		for j := 0; j < len(canvas[0])-2; j++ {
			window := [3][3]rune{
				{canvas[i][j], canvas[i][j+1], canvas[i][j+2]},
				{canvas[i+1][j], canvas[i+1][j+1], canvas[i+1][j+2]},
				{canvas[i+2][j], canvas[i+2][j+1], canvas[i+2][j+2]},
			}
			if isValidCross(window) {
				total++
			}
		}
	}
	return total
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

	canvas := parseInput(string(data))

	count := countCrosses(canvas)
	fmt.Printf("Count: %d\n", count)
}
