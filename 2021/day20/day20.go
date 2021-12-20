package main

import (
	"github.com/reusing-code/adventofcode/gohelpers"
)

type coord struct {
	x int
	y int
}

type field struct {
	data     map[coord]bool
	min      coord
	max      coord
	infinity bool
}

func (f *field) set(x, y int, v bool) {
	c := coord{x, y}
	if v {
		f.data[c] = true
		if x < f.min.x {
			f.min.x = x
		}
		if x > f.max.x {
			f.max.x = x
		}
		if y < f.min.y {
			f.min.y = y
		}
		if y > f.max.y {
			f.max.y = y
		}
	} else {
		delete(f.data, c)
	}
}

func (f *field) get(x, y int) bool {
	if x < f.min.x || x > f.max.x {
		return f.infinity
	}
	if y < f.min.y || y > f.max.y {
		return f.infinity
	}

	if v, ok := f.data[coord{x, y}]; ok {
		return v
	} else {
		return false
	}
}

func (f *field) calcVal(x, y int) int {
	idx := 0
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			idx *= 2
			if f.get(i, j) {
				idx += 1
			}
		}
	}
	return idx
}

func puzzle1(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	decoder := make([]bool, 512)
	for i, c := range split[0][0] {
		if c == '#' {
			decoder[i] = true
		}
	}
	f := &field{}
	f.data = make(map[coord]bool)
	for i, str := range split[1] {
		for j, c := range str {
			if c == '#' {
				f.set(j, i, true)
			}
		}
	}

	f.infinity = false

	for i := 0; i < 2; i++ {
		fNew := &field{}
		fNew.data = make(map[coord]bool)
		fromX := f.min.x - 1
		toX := f.max.x + 1
		fromY := f.min.y - 1
		toY := f.max.y + 1
		for y := fromY; y <= toY; y++ {
			for x := fromX; x <= toX; x++ {
				fNew.set(x, y, decoder[f.calcVal(x, y)])
			}
		}
		oldInfinity := f.infinity
		f = fNew
		if oldInfinity {
			f.infinity = decoder[511]
		} else {
			f.infinity = decoder[0]
		}
	}
	return len(f.data)
}

func puzzle2(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	decoder := make([]bool, 512)
	for i, c := range split[0][0] {
		if c == '#' {
			decoder[i] = true
		}
	}
	f := &field{}
	f.data = make(map[coord]bool)
	for i, str := range split[1] {
		for j, c := range str {
			if c == '#' {
				f.set(j, i, true)
			}
		}
	}

	f.infinity = false

	for i := 0; i < 50; i++ {
		fNew := &field{}
		fNew.data = make(map[coord]bool)
		fromX := f.min.x - 1
		toX := f.max.x + 1
		fromY := f.min.y - 1
		toY := f.max.y + 1
		for y := fromY; y <= toY; y++ {
			for x := fromX; x <= toX; x++ {
				fNew.set(x, y, decoder[f.calcVal(x, y)])
			}
		}
		oldInfinity := f.infinity
		f = fNew
		if oldInfinity {
			f.infinity = decoder[511]
		} else {
			f.infinity = decoder[0]
		}
	}
	return len(f.data)
}
