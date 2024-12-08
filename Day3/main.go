package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sumMulExpressions(input string) int {
	// Compile the regular expression
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re, _ := regexp.Compile(pattern)

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)

	// Get the operands from each match
	pattern2 := `\d{1,3}`
	re2, _ := regexp.Compile(pattern2)

	// Print the matches
	total := 0
	for _, match := range matches {
		mults := re2.FindAllString(match, -1)
		//multiply the operands
		x, _ := strconv.Atoi(mults[0])
		y, _ := strconv.Atoi(mults[1])
		total += x * y
	}

	return total
}

func removeTextBetweenPatterns(input string) string {
	// Escape parentheses in the pattern and use non-greedy matching
	pattern := `(?s)don't\(\).*?do\(\)`
	re := regexp.MustCompile(pattern)

	// Replace all matches with empty string
	return re.ReplaceAllString(input, "")
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

	cleanInput := removeTextBetweenPatterns(string(data))
	sum := sumMulExpressions(string(cleanInput))
	fmt.Printf("Count: %d\n", sum)
}
