package main

import (
	"fmt"
)

func puzzle1(input []string) int {
	var floor [1000][1000]int
	for _, str := range input {
		var (
			x1 int
			y1 int
			x2 int
			y2 int
		)
		n, err := fmt.Sscanf(str, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			panic(n)
		}

		if x1 == x2 {
			if y1 > y2 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			for y := y1; y <= y2; y++ {
				floor[x1][y]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}
			for x := x1; x <= x2; x++ {
				floor[x][y1]++
			}
		}
	}
	result := 0

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if floor[x][y] > 1 {
				result++
			}
		}
	}

	return result
}

func puzzle2(input []string) int {
	var floor [1000][1000]int
	for _, str := range input {
		var (
			x1 int
			y1 int
			x2 int
			y2 int
		)
		n, err := fmt.Sscanf(str, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			panic(n)
		}

		if x1 == x2 {
			if y1 > y2 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			for y := y1; y <= y2; y++ {
				floor[x1][y]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}
			for x := x1; x <= x2; x++ {
				floor[x][y1]++
			}
		} else {
			if x1 > x2 {
				tmp := x1
				x1 = x2
				x2 = tmp
				tmp = y1
				y1 = y2
				y2 = tmp
			}
			y := y1
			for x := x1; x <= x2; x++ {
				floor[x][y]++
				if (y1 < y2) {
					y++
				} else {
					y--
				}
			}
		}
	}

	result := 0

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if floor[x][y] > 1 {
				result++
			}
		}
	}

	return result
}
