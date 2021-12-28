package main

import "fmt"

type Value int64

const (
	Empty = iota
	East
	South
)

type Field struct {
	data [][]Value
}

func (f *Field) getSize() (int, int) {
	return len(f.data), len(f.data[0])
}

func (f *Field) get(x, y int) Value {
	xs, ys := f.getSize()
	return f.data[x%xs][y%ys]
}

func (f *Field) set(x, y int, v Value) {
	xs, ys := f.getSize()
	f.data[x%xs][y%ys] = v
}

func (f *Field) copyEmpty() *Field {
	result := &Field{make([][]Value, len(f.data))}
	for i := range result.data {
		result.data[i] = make([]Value, len(f.data[0]))
	}
	return result
}

func (f *Field) compare(other *Field) bool {
	for i, slice := range f.data {
		for j, val := range slice {
			if val != other.data[i][j] {
				return false
			}
		}
	}
	return true
}

func (f *Field) print() {
	fmt.Println("")
	for _, slice := range f.data {
		for _, val := range slice {
			if val == South {
				fmt.Print("v")
			}
			if val == East {
				fmt.Print(">")
			}
			if val == Empty {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func puzzle1(input []string) int {
	f := &Field{make([][]Value, len(input))}
	for i, str := range input {
		f.data[i] = make([]Value, len(input[0]))
		for j, c := range str {
			if c == '>' {
				f.data[i][j] = East
			}
			if c == 'v' {
				f.data[i][j] = South
			}
		}
	}

	for i := 0; ; i++ {
		fNew := f.copyEmpty()

		// EAST
		for x, slice := range f.data {
			for y, val := range slice {
				if val == East {
					if f.get(x, y+1) == Empty {
						fNew.set(x, y+1, East)
					} else {
						fNew.set(x, y, East)
					}
				}
			}
		}

		// SOUTH
		for x, slice := range f.data {
			for y, val := range slice {
				if val == South {
					if f.get(x+1, y) != South && fNew.get(x+1, y) == Empty {
						fNew.set(x+1, y, South)
					} else {
						fNew.set(x, y, South)
					}
				}
			}
		}

		if f.compare(fNew) {
			return i + 1
		}
		f = fNew
	}
}

func puzzle2(input []string) int {
	return 1
}
