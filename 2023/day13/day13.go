package main

import (
	"slices"

	h "github.com/reusing-code/adventofcode/gohelpers"
)

func findMirror(field [][]byte, notAllowed int) int {
	sizeX := len(field)
	sizeY := len(field[0])
	possibleMirrorsX := make(map[int]int)
	for i := range field {
		for j := i + 1; j < sizeX; j++ {
			if slices.Compare(field[i], field[j]) == 0 {
				if (j+i)%2 == 1 {
					possibleMirrorsX[(i+j+1)/2] += 1
				}
			}
		}
	}
	for k, v := range possibleMirrorsX {
		if v >= k || v >= (sizeX-k) {
			if k*100 != notAllowed {
				return k * 100
			}
		}
	}
	possibleMirrorsY := make(map[int]int)
	for i := 0; i < sizeY; i++ {
		for j := i + 1; j < sizeY; j++ {
			equal := true
			for x := range field {
				if field[x][i] != field[x][j] {
					equal = false
				}
			}
			if equal {
				if (j+i)%2 == 1 {
					possibleMirrorsY[(i+j+1)/2] += 1
				}
			}
		}
	}
	for k, v := range possibleMirrorsY {
		if v >= k || v >= (sizeY-k) {
			if k != notAllowed {
				return k
			}
		}
	}
	return 0
}

func puzzle1(input []string) int {
	result := 0
	split := h.SplitByEmptyLine(input)
	for _, p := range split {
		field := h.ParseCharField(p)
		result += findMirror(field, 0)
	}
	return result
}

func findMirrorWithSmudge(field [][]byte) int {
	original := findMirror(field, 0)
	for i, row := range field {
		for j, c := range row {
			if c == '#' {
				field[i][j] = '.'
			} else {
				field[i][j] = '#'
			}
			value := findMirror(field, original)
			if value != 0 {
				return value
			}
			field[i][j] = c
		}
	}
	return 0
}

func puzzle2(input []string) int {
	result := 0
	split := h.SplitByEmptyLine(input)
	for _, p := range split {
		field := h.ParseCharField(p)
		result += findMirrorWithSmudge(field)
	}
	return result
}
