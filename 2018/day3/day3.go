package main

import (
	"fmt"
	"strings"
)

type rect struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func puzzle1(input []string) int {
	var fabric [1000][1000]int
	for _, str := range input {
		var (
			r rect
		)
		reader := strings.NewReader(str)
		n, err := fmt.Fscanf(reader, "#%d @ %d,%d: %dx%d", &r.id, &r.x, &r.y, &r.w, &r.h)
		if err != nil {
			panic(err)
		}
		if n != 5 {
			panic(n)
		}
		for x := r.x; x < r.x + r.w; x++ {
			for y := r.y; y < r.y + r.h; y++ {
				fabric[x][y]++
			}
		}
	}
	result := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				result++
			}
		}
	}
	return result
}

func puzzle2(input []string) int {
	var fabric [1000][1000]int
	regions := make(map[int]bool)

	for _, str := range input {
		var (
			r rect
		)
		reader := strings.NewReader(str)
		n, err := fmt.Fscanf(reader, "#%d @ %d,%d: %dx%d", &r.id, &r.x, &r.y, &r.w, &r.h)
		if err != nil {
			panic(err)
		}
		if n != 5 {
			panic(n)
		}
		regions[r.id] = true
		for x := r.x; x < r.x + r.w; x++ {
			for y := r.y; y < r.y + r.h; y++ {
				if fabric[x][y] != 0 {
					regions[r.id] = false
					regions[fabric[x][y]] = false
				}
				fabric[x][y] = r.id
			}
		}
	}
	for k, v := range regions {
		if v {
			return k
		}
	}
	return 0
}
