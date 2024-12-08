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

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}

	for _, num := range arr[:len(arr)/2] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	for _, num := range arr[len(arr)/2+1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func sortData(data Data) Data {
	sortedCol1 := quickSort(data.col1)
	sortedCol2 := quickSort(data.col2)

	return Data{sortedCol1, sortedCol2}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func byElementDistance(data Data) int {
	total := 0

	for i, num := range data.col1 {
		total += abs(num - data.col2[i])
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

	sortedData := sortData(parseInput(string(data)))
	distance := byElementDistance(sortedData)

	fmt.Printf("Distance: %d\n", distance)
}
