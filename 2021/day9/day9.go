package main

import (
	"strconv"
	"sort"
)

func puzzle1(input []string) int {
	field := make([][]int, len(input))
	for i, str := range input {
		field[i] = make([]int, len(str))
		for j, s := range str {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				panic(err)
			}
			field[i][j] = n
		}
	}
	result := 0

	for i, row := range field {
		for j, val := range row {
			if i > 0 && field[i-1][j] <= val {
				continue
			}
			if j > 0 && field[i][j-1] <= val {
				continue
			}
			if i < len(field) - 1 && field[i+1][j] <= val {
				continue
			}
			if j < len(row) - 1 && field[i][j+1] <= val {
				continue
			}
			result += val + 1
		}
	}
	return result
}

type spot struct {
	value int
	nextX int
	nextY int
	count int
}

func puzzle2(input []string) int {
	field := make([][]spot, len(input))
	for i, str := range input {
		field[i] = make([]spot, len(str))
		for j, s := range str {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				panic(err)
			}
			field[i][j].value = n
			field[i][j].nextX = -1
			field[i][j].nextY = -1
			field[i][j].count = 0
		}
	}

	for i, row := range field {
		for j, val := range row {
			if val.value == 9 {
				continue
			}
			if i > 0 && field[i-1][j].value < val.value {
				field[i][j].nextX = i - 1
				field[i][j].nextY = j
				continue
			}
			if j > 0 && field[i][j-1].value < val.value {
				field[i][j].nextX = i
				field[i][j].nextY = j - 1
				continue
			}
			if i < len(field) - 1 && field[i+1][j].value < val.value {
				field[i][j].nextX = i + 1
				field[i][j].nextY = j
				continue
			}
			if j < len(row) - 1 && field[i][j+1].value < val.value {
				field[i][j].nextX = i
				field[i][j].nextY = j + 1
				continue
			}
		}
	}

	for i, row := range field {
		for j, val := range row {
			if val.value == 9 {
				continue
			}
			i1 := i
			j1 := j
			for {
				if field[i1][j1].nextX == -1 || field[i1][j1].nextY == -1 {
					field[i1][j1].count++
					break
				}
				i1, j1 = field[i1][j1].nextX, field[i1][j1].nextY
			}
		}
	}

	basins := make([]int, 0)

	for i, row := range field {
		for j, val := range row {
			if val.value == 9 {
				continue
			}
			if field[i][j].nextX == -1 || field[i][j].nextY == -1 {
				basins = append(basins, field[i][j].count)
			}
		}
	}
	sort.Slice(basins, func (i, j int) bool {
		return basins[i] > basins[j]
	})
	return basins[0] * basins[1] * basins[2]
}
