package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	col1 []int
	col2 []int
}

func parseInput(input string) Data {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)

	col1 := make([]int, rows)
	col2 := make([]int, rows)
	for i, line := range lines {
		values := strings.Split(strings.TrimSpace(line), "   ")
		col1[i], _ = strconv.Atoi(values[0])
		col2[i], _ = strconv.Atoi(values[1])
	}

	return Data{col1, col2}
}

func frequencyCount(col []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range col {
		counts[num]++
	}
	return counts
}

func similarityScore(col []int, countMap map[int]int) int {
	var score = 0
	for _, num := range col {
		multiple := countMap[num]
		score += multiple * num
	}

	return score
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

	countMap := frequencyCount(parseInput(string(data)).col2)
	score := similarityScore(parseInput(string(data)).col1, countMap)

	fmt.Printf("Similarity score: %d\n", score)
}
