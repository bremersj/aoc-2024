package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculation struct {
	total  int
	values []int
}

func parseInput(input string) []Calculation {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	if rows == 0 {
		return nil
	}
	calculations := make([]Calculation, rows)

	for i, line := range lines {
		elements := strings.Split(strings.TrimSpace(line), ":")
		total, err := strconv.Atoi(elements[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input.txt: %v\n", err)
			os.Exit(1)
		}
		values := []int{}
		for _, numStr := range strings.Split(strings.TrimSpace(elements[1]), " ") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing input.txt: %v\n", err)
				os.Exit(1)
			}
			values = append(values, num)
		}
		calculations[i] = Calculation{total, values}
	}

	return calculations
}

func hasValidValues(c Calculation) bool {
	// if only one value and equal to total, return true
	if len(c.values) == 1 {
		return c.values[0] == c.total
	}

	// check if total is divisible by last value
	if c.total%c.values[len(c.values)-1] == 0 {
		// divide total by last value and recurse
		if hasValidValues(Calculation{c.total / c.values[len(c.values)-1], c.values[:len(c.values)-1]}) {
			return true
		}
	}

	// substract last value from total and recurse
	if hasValidValues(Calculation{c.total - c.values[len(c.values)-1], c.values[:len(c.values)-1]}) {
		return true
	}

	return false
}

func sumValidValues(calculations []Calculation) int {
	sum := 0
	for _, c := range calculations {
		if hasValidValues(c) {
			sum += c.total
			fmt.Println(sum)
		}
	}

	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v", err)
		os.Exit(1)
	}

	calculations := parseInput(string(data))

	result := sumValidValues(calculations)
	fmt.Printf("Result: %d\n", result)
}
