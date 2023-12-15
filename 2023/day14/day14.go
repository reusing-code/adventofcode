package main

import (
	h "github.com/reusing-code/adventofcode/gohelpers"
)

func tiltNorth(field [][]byte) {
	height := len(field)
	width := len(field[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] == 'O' {
				for ii := i - 1; ii >= 0; ii-- {
					if field[ii][j] == '.' {
						field[ii+1][j] = '.'
						field[ii][j] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltWest(field [][]byte) {
	height := len(field)
	width := len(field[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] == 'O' {
				for jj := j - 1; jj >= 0; jj-- {
					if field[i][jj] == '.' {
						field[i][jj+1] = '.'
						field[i][jj] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltSouth(field [][]byte) {
	height := len(field)
	width := len(field[0])
	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			if field[i][j] == 'O' {
				for ii := i + 1; ii < height; ii++ {
					if field[ii][j] == '.' {
						field[ii-1][j] = '.'
						field[ii][j] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltEast(field [][]byte) {
	height := len(field)
	width := len(field[0])
	for i := 0; i < height; i++ {
		for j := width - 1; j >= 0; j-- {
			if field[i][j] == 'O' {
				for jj := j + 1; jj < width; jj++ {
					if field[i][jj] == '.' {
						field[i][jj-1] = '.'
						field[i][jj] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func cycle(field [][]byte) {
	tiltNorth(field)
	tiltWest(field)
	tiltSouth(field)
	tiltEast(field)
}

func calcWeight(field [][]byte) int {
	height := len(field)
	width := len(field[0])
	result := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] == 'O' {
				result += height - i
			}
		}
	}
	return result
}

func puzzle1(input []string) int {
	field := h.ParseCharField(input)
	tiltNorth(field)
	return calcWeight(field)
}

func hashKey(field [][]byte) string {
	result := make([]byte, 0, len(field)*(len(field[0])+1))
	for _, v := range field {
		result = append(result, v...)
		result = append(result, 'X')
	}
	return string(result)
}

func puzzle2(input []string) int {
	field := h.ParseCharField(input)
	weightAfterCycle := make([]int, 0)
	weightAfterCycle = append(weightAfterCycle, 0)
	lookup := make(map[string]int)
	for i := 1; ; i++ {
		cycle(field)
		k := hashKey(field)
		if v, ok := lookup[k]; ok {
			return weightAfterCycle[((1000000000-v)%(i-v))+v]
		}
		weightAfterCycle = append(weightAfterCycle, calcWeight(field))
		lookup[k] = i
	}
}
