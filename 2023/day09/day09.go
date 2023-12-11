package main

import (
	"github.com/reusing-code/adventofcode/gohelpers"
)

func calculateDifferences(input string) [][]int {
	values := make([][]int, 0)
	values = append(values, gohelpers.ParseIntVecSingleLineSpace(input))
	allZeroes := false
	for !allZeroes {
		allZeroes = true
		newValues := make([]int, 0)
		current := values[len(values)-1]
		for i := 0; i < len(current)-1; i++ {
			diff := current[i+1] - current[i]
			if diff != 0 {
				allZeroes = false
			}
			newValues = append(newValues, diff)
		}
		values = append(values, newValues)
	}
	return values
}

func puzzle1(input []string) int {
	result := 0
	for _, line := range input {
		values := calculateDifferences(line)

		for i := len(values) - 2; i >= 0; i-- {
			v := values[i][len(values[i])-1] + values[i+1][len(values[i+1])-1]
			values[i] = append(values[i], v)
		}
		result += values[0][len(values[0])-1]
	}
	return result
}

func puzzle2(input []string) int {
	result := 0
	for _, line := range input {
		values := calculateDifferences(line)

		for i := len(values) - 2; i >= 0; i-- {
			v := values[i][0] - values[i+1][0]
			values[i] = append([]int{v}, values[i]...)
		}
		result += values[0][0]
	}
	return result
}
