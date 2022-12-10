package main

import (
	"github.com/reusing-code/adventofcode/gohelpers"
)

func puzzle1(input []string) int {
	intField := gohelpers.ParseIntField(input)
	boolField := make([][]bool, len(intField))
	for i := range boolField {
		boolField[i] = make([]bool, len(intField[i]))
	}

	for i := range intField {
		maxPos := -1
		maxNeg := -1
		for j := range intField[0] {
			valPos := intField[i][j]
			if valPos > maxPos {
				maxPos = valPos
				boolField[i][j] = true
			}
			valNeg := intField[i][len(intField[0])-1-j]
			if valNeg > maxNeg {
				maxNeg = valNeg
				boolField[i][len(intField[0])-1-j] = true
			}
		}
	}

	for j := range intField[0] {
		maxPos := -1
		maxNeg := -1
		for i := range intField {
			valPos := intField[i][j]
			if valPos > maxPos {
				maxPos = valPos
				boolField[i][j] = true
			}
			valNeg := intField[len(intField)-1-i][j]
			if valNeg > maxNeg {
				maxNeg = valNeg
				boolField[len(intField)-1-i][j] = true
			}
		}
	}

	result := 0
	for i := range boolField {
		for j := range boolField[i] {
			if boolField[i][j] {
				result += 1
			}
		}
	}

	return result
}

func puzzle2(input []string) int {
	intField := gohelpers.ParseIntField(input)
	viewField := make([][]int, len(intField))
	for i := range viewField {
		viewField[i] = make([]int, len(intField[i]))
		for j := range viewField[i] {
			viewField[i][j] = 1
		}
	}

	for i := range intField {
		for j := range intField[0] {
			val := intField[i][j]
			left := 0
			for jl := j - 1; jl >= 0; jl-- {
				left++
				if intField[i][jl] >= val {
					break
				}
			}
			right := 0
			for jr := j + 1; jr < len(intField[i]); jr++ {
				right++
				if intField[i][jr] >= val {
					break
				}
			}
			up := 0
			for iu := i - 1; iu >= 0; iu-- {
				up++
				if intField[iu][j] >= val {
					break
				}
			}
			down := 0
			for id := i + 1; id < len(intField); id++ {
				down++
				if intField[id][j] >= val {
					break
				}
			}
			viewField[i][j] = left * right * up * down

		}
	}

	result := 0
	for i := range viewField {
		for j := range viewField[i] {
			if viewField[i][j] > result {
				result = viewField[i][j]
			}
		}
	}

	return result
}
