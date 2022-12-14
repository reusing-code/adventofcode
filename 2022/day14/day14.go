package main

import (
	"fmt"
	"sort"
	"strings"
)

type coord struct {
	x int
	y int
}

func set(columns map[int][]int, val coord) {
	if _, prs := columns[val.x]; !prs {
		columns[val.x] = []int{val.y}
	} else {
		columns[val.x] = append(columns[val.x], val.y)
	}
}

func sortColumns(columns map[int][]int) {
	for _, slice := range columns {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	}
}

func blocked(columns map[int][]int, c coord) bool {
	if _, prs := columns[c.x]; !prs {
		return false
	}
	for _, v := range columns[c.x] {
		if v > c.y {
			return false
		}
		if v == c.y {
			return true
		}
	}
	return false
}

func pourSand(columns map[int][]int, from coord) bool {
	current := from
	if _, prs := columns[current.x]; !prs {
		return false
	}
	for _, v := range columns[current.x] {
		if v > current.y {
			current.y = v
		}
		if v == current.y {
			if (!blocked(columns, coord{current.x - 1, current.y})) {
				return pourSand(columns, coord{current.x - 1, current.y})
			}
			if (!blocked(columns, coord{current.x + 1, current.y})) {
				return pourSand(columns, coord{current.x + 1, current.y})
			}
			columns[current.x] = append(columns[current.x], current.y-1)
			sort.Slice(columns[current.x], func(i, j int) bool {
				return columns[current.x][i] < columns[current.x][j]
			})
			return true
		}
	}
	return false
}

func pourSand2(columns map[int][]int, from coord, floor int) bool {
	current := from
	if _, prs := columns[current.x]; !prs {
		columns[current.x] = []int{floor - 1, floor}
		return true
	}
	for _, v := range columns[current.x] {
		if v > current.y {
			current.y = v
		}
		if v == current.y {
			if (!blocked(columns, coord{current.x - 1, current.y})) {
				return pourSand2(columns, coord{current.x - 1, current.y}, floor)
			}
			if (!blocked(columns, coord{current.x + 1, current.y})) {
				return pourSand2(columns, coord{current.x + 1, current.y}, floor)
			}
			columns[current.x] = append(columns[current.x], current.y-1)
			sort.Slice(columns[current.x], func(i, j int) bool {
				return columns[current.x][i] < columns[current.x][j]
			})
			if current.x == 500 && current.y-1 == 0 {
				return false
			}
			return true
		}
	}
	return false
}

func puzzle1(input []string) int {
	columns := make(map[int][]int)
	for _, inStr := range input {
		splitStr := strings.Split(inStr, " -> ")
		current := coord{0, 0}
		fmt.Sscanf(splitStr[0], "%d,%d", &current.x, &current.y)
		set(columns, current)
		for i := 1; i < len(splitStr); i++ {
			next := coord{0, 0}
			fmt.Sscanf(splitStr[i], "%d,%d", &next.x, &next.y)
			if current.x == next.x {
				if current.y > next.y {
					for j := current.y - 1; j >= next.y; j-- {
						current.y = j
						set(columns, current)
					}
				} else {
					for j := current.y + 1; j <= next.y; j++ {
						current.y = j
						set(columns, current)
					}
				}
			} else {
				if current.x > next.x {
					for j := current.x - 1; j >= next.x; j-- {
						current.x = j
						set(columns, current)
					}
				} else {
					for j := current.x + 1; j <= next.x; j++ {
						current.x = j
						set(columns, current)
					}
				}
			}
		}
	}
	sortColumns(columns)

	pourFrom := coord{500, 0}
	i := 0
	for {
		if pourSand(columns, pourFrom) {
			i++
		} else {
			break
		}
	}
	return i
}

func puzzle2(input []string) int {
	columns := make(map[int][]int)
	maxY := 0
	for _, inStr := range input {
		splitStr := strings.Split(inStr, " -> ")
		current := coord{0, 0}
		fmt.Sscanf(splitStr[0], "%d,%d", &current.x, &current.y)
		if current.y > maxY {
			maxY = current.y
		}
		set(columns, current)
		for i := 1; i < len(splitStr); i++ {
			next := coord{0, 0}
			fmt.Sscanf(splitStr[i], "%d,%d", &next.x, &next.y)
			if next.y > maxY {
				maxY = next.y
			}
			if current.x == next.x {
				if current.y > next.y {
					for j := current.y - 1; j >= next.y; j-- {
						current.y = j
						set(columns, current)
					}
				} else {
					for j := current.y + 1; j <= next.y; j++ {
						current.y = j
						set(columns, current)
					}
				}
			} else {
				if current.x > next.x {
					for j := current.x - 1; j >= next.x; j-- {
						current.x = j
						set(columns, current)
					}
				} else {
					for j := current.x + 1; j <= next.x; j++ {
						current.x = j
						set(columns, current)
					}
				}
			}
		}
	}
	sortColumns(columns)
	for i := range columns {
		columns[i] = append(columns[i], maxY+2)
	}

	pourFrom := coord{500, 0}
	i := 0
	for {
		if pourSand2(columns, pourFrom, maxY+2) {
			i++
		} else {
			break
		}
	}
	return i + 1
}
