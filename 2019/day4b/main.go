package main

import (
	"fmt"
)

func validCombination(in int) bool {

	doubleDigit := false
	pairFound := false
	inDoubles := false
	lastDigit := 11
	for i := in; i > 0; i = i / 10 {
		digit := i % 10
		if digit == lastDigit {
			if inDoubles {
				doubleDigit = false
			} else {
				doubleDigit = true
			}
			inDoubles = true
		} else if digit > lastDigit {
			return false
		} else {
			if doubleDigit {
				pairFound = true
			}
			inDoubles = false
			doubleDigit = false
		}
		lastDigit = digit
	}
	if doubleDigit {
		pairFound = true
	}

	return pairFound
}

func countCombinations(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if validCombination(i) {
			count++
		}
	}
	return count
}

func main() {
	result := countCombinations(357253, 892942)
	fmt.Printf("Result: %v\n", result)
}
