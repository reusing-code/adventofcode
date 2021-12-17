package main

import (
	"fmt"
)

func puzzle1(input []string) int {
	var x1, x2, y1, y2 int
	_, err := fmt.Sscanf(input[0], "target area: x=%v..%v, y=%v..%v", &x1, &x2, &y1, &y2)
	if err != nil {
		panic(err)
	}
	globalMaxY := 0
	possibleX := make([]int, 0)

	for x0 := 1; x0 <= x2; x0++ {
		dx := x0
		x := 0
		for x <= x2 {
			if x >= x1 {
				possibleX = append(possibleX, x0)
				break
			}
			if dx == 0 {
				break
			}
			x += dx
			dx--
		}
	}

	absY1 := y1
	if y1 < 0 {
		absY1 *= -1
	}
	absY2 := y2
	if y2 < 0 {
		absY2 *= -1
	}
	absY := absY1
	if absY2 > absY1 {
		absY = absY2
	}

	for _, x0 := range possibleX {
		for y0 := -absY; y0 < absY; y0++ {
			hit := false
			maxY := 0
			step := 0
			x := 0
			y := 0
			dx := x0
			dy := y0
			for x <= x2 && y >= y1 {
				if y > maxY {
					maxY = y
				}
				if x >= x1 && y <= y2 {
					hit = true
					break
				}
				step++
				x += dx
				y += dy
				if dx > 0 {
					dx--
				}
				dy--
			}
			if hit {
				if maxY > globalMaxY {
					globalMaxY = maxY
				}
			}
		}
	}
	return globalMaxY
}

func puzzle2(input []string) int {
	var x1, x2, y1, y2 int
	_, err := fmt.Sscanf(input[0], "target area: x=%v..%v, y=%v..%v", &x1, &x2, &y1, &y2)
	if err != nil {
		panic(err)
	}
	globalMaxY := 0
	count := 0
	possibleX := make([]int, 0)
	for x0 := 1; x0 <= x2; x0++ {
		dx := x0
		x := 0
		for x <= x2 {
			if x >= x1 {
				possibleX = append(possibleX, x0)
				break
			}
			if dx == 0 {
				break
			}
			x += dx
			dx--
		}
	}

	absY1 := y1
	if y1 < 0 {
		absY1 *= -1
	}
	absY2 := y2
	if y2 < 0 {
		absY2 *= -1
	}
	absY := absY1
	if absY2 > absY1 {
		absY = absY2
	}

	for _, x0 := range possibleX {
		for y0 := -absY; y0 < absY; y0++ {
			hit := false
			maxY := 0
			step := 0
			x := 0
			y := 0
			dx := x0
			dy := y0
			for x <= x2 && y >= y1 {
				if y > maxY {
					maxY = y
				}
				if x >= x1 && y <= y2 {
					hit = true
					break
				}
				step++
				x += dx
				y += dy
				if dx > 0 {
					dx--
				}
				dy--
			}
			if hit {
				count++
				if maxY > globalMaxY {
					globalMaxY = maxY
				}
			}
		}
	}
	return count
}
