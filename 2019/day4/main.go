package main

import (
	"fmt"
)

func validCombination(in int) bool {

	doubleDigit := false
	lastDigit := 11
	for i := in; i > 0; i = i / 10 {
		digit := i % 10
		if digit == lastDigit {
			doubleDigit = true
		} else if digit > lastDigit {
			return false
		}
		lastDigit = digit
	}

	return doubleDigit
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
