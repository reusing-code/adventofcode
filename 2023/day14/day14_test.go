package main

import (
	"reflect"
	"testing"
)

func Test_day14_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#....",
		}, want: 136},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day14_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day14_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"O....#....",
			"O.OO#....#",
			".....##...",
			"OO.#O....O",
			".O.....O#.",
			"O.#..O.#.#",
			"..O..#O..O",
			".......O..",
			"#....###..",
			"#OO..#....",
		}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day14_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
