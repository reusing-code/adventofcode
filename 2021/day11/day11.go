package main

import "strconv"

type coord struct {
	x int
	y int
}

func puzzle1(input []string) int {
	grid := make([][]int, 0)
	for _, row := range input {
		rowSlice := make([]int, 0)
		for _, elem := range row {
			n, err := strconv.Atoi(string(elem))
			if err != nil {
				panic(err)
			}
			rowSlice = append(rowSlice, n)
		}
		grid = append(grid, rowSlice)
	}

	flashes := 0

	for i := 0; i < 100; i++ {
		flashQueue := make([]coord, 0)
		for x, row := range grid {
			for y := range row {
				grid[x][y]++
				if grid[x][y] > 9 {
					flashQueue = append(flashQueue, coord{x, y})
				}
			}
		}
		for len(flashQueue) > 0 {
			x0, y0 := flashQueue[0].x, flashQueue[0].y
			for x := x0 - 1; x <= x0+1; x++ {
				if x < 0 || x >= len(grid) {
					continue
				}
				for y := y0 - 1; y <= y0+1; y++ {
					if y < 0 || y >= len(grid[x]) {
						continue
					}
					if grid[x][y] == 0 {
						continue
					}
					grid[x][y]++
					if grid[x][y] == 10 {
						flashQueue = append(flashQueue, coord{x, y})
					}
				}
			}
			grid[x0][y0] = 0
			flashes++
			flashQueue = flashQueue[1:]
		}
	}

	return flashes
}

func puzzle2(input []string) int {
	grid := make([][]int, 0)
	for _, row := range input {
		rowSlice := make([]int, 0)
		for _, elem := range row {
			n, err := strconv.Atoi(string(elem))
			if err != nil {
				panic(err)
			}
			rowSlice = append(rowSlice, n)
		}
		grid = append(grid, rowSlice)
	}

	for i := 0; true; i++ {
		flashes := 0
		flashQueue := make([]coord, 0)
		for x, row := range grid {
			for y := range row {
				grid[x][y]++
				if grid[x][y] > 9 {
					flashQueue = append(flashQueue, coord{x, y})
				}
			}
		}
		for len(flashQueue) > 0 {
			x0, y0 := flashQueue[0].x, flashQueue[0].y
			for x := x0 - 1; x <= x0+1; x++ {
				if x < 0 || x >= len(grid) {
					continue
				}
				for y := y0 - 1; y <= y0+1; y++ {
					if y < 0 || y >= len(grid[x]) {
						continue
					}
					if grid[x][y] == 0 {
						continue
					}
					grid[x][y]++
					if grid[x][y] == 10 {
						flashQueue = append(flashQueue, coord{x, y})
					}
				}
			}
			grid[x0][y0] = 0
			flashes++
			flashQueue = flashQueue[1:]
		}
		if flashes == len(grid)*len(grid[0]) {
			return i + 1
		}
	}

	return 0
}
