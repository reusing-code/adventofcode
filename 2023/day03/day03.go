package main

import (
	"unicode"
)

func hasSymbol(input []string, x, y int) bool {
	if x < 0 || x >= len(input) {
		return false
	}
	if y < 0 || y >= len(input[x]) {
		return false
	}
	c := input[x][y]
	if c == '.' {
		return false
	}
	if unicode.IsNumber(rune(c)) {
		return false
	}
	if unicode.IsLetter(rune(c)) {
		return false
	}
	return true
}

func hasSymbolAnyNeighbour(input []string, x, y int) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if hasSymbol(input, x+dx, y+dy) {
				return true
			}
		}
	}
	return false
}

func puzzle1(input []string) int {
	result := 0
	for x, line := range input {
		currentNumber := 0
		symbol := false
		for y, c := range line {
			if unicode.IsNumber(c) {
				currentNumber *= 10
				currentNumber += (int)(c - '0')
				symbol = symbol || hasSymbolAnyNeighbour(input, x, y)
			} else {
				if symbol {
					result += currentNumber
				}
				currentNumber = 0
				symbol = false
			}
		}
		if symbol {
			result += currentNumber
		}
	}
	return result
}

type number struct {
	value int
	fromX int
	toX   int
	fromY int
	toY   int
	valid bool
}

func puzzle2(input []string) int {
	result := 0
	numbers := make([]number, 0)
	for x, line := range input {
		currentNumber := &number{}
		for y, c := range line {
			if unicode.IsNumber(c) {
				if !currentNumber.valid {
					currentNumber.valid = true
					currentNumber.fromX = x - 1
					currentNumber.toX = x + 1
					currentNumber.fromY = y - 1
				}
				currentNumber.value *= 10
				currentNumber.value += (int)(c - '0')
				currentNumber.toY = y + 1
			} else {
				if currentNumber.valid {
					numbers = append(numbers, *currentNumber)
					currentNumber = &number{}
				}
			}
		}
		if currentNumber.valid {
			numbers = append(numbers, *currentNumber)
		}
	}

	for x, line := range input {
		for y, c := range line {
			if c == '*' {
				gearValue := 0
				for _, n := range numbers {
					if x >= n.fromX && x <= n.toX && y >= n.fromY && y <= n.toY {
						if gearValue == 0 {
							gearValue += n.value
						} else {
							result += gearValue * n.value
							break
						}
					}
				}
			}
		}
	}
	return result
}
