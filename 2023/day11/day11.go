package main

import (
	"slices"

	h "github.com/reusing-code/adventofcode/gohelpers"
)

func puzzle(input []string, expansionFactor int) int {
	field := h.ParseCharField(input)
	galaxies := make([]h.Coord, 0)
	emptyRowsMap := make(map[int]bool)
	emptyColumnsMap := make(map[int]bool)
	for i := 0; i < len(field); i++ {
		emptyRowsMap[i] = true
	}
	for i := 0; i < len(field[0]); i++ {
		emptyColumnsMap[i] = true
	}
	for i, l := range field {
		for j, c := range l {
			if c == '#' {
				delete(emptyRowsMap, i)
				delete(emptyColumnsMap, j)
				galaxies = append(galaxies, h.Coord{i, j})
			}
		}
	}
	emptyRows := make([]int, 0)
	for k := range emptyRowsMap {
		emptyRows = append(emptyRows, k)
	}
	slices.Sort(emptyRows)
	emptyColumns := make([]int, 0)
	for k := range emptyColumnsMap {
		emptyColumns = append(emptyColumns, k)
	}
	slices.Sort(emptyColumns)

	for i := range galaxies {
		g := &galaxies[i]
		dx, _ := slices.BinarySearch(emptyRows, g.X)
		g.X += dx * (expansionFactor - 1)
		dy, _ := slices.BinarySearch(emptyColumns, g.Y)
		g.Y += dy * (expansionFactor - 1)
	}
	result := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += galaxies[i].Manhatten(galaxies[j])
		}
	}

	return result
}

func puzzle1(input []string) int {
	return puzzle(input, 2)
}

func puzzle2(input []string) int {
	return puzzle(input, 1000000)
}
