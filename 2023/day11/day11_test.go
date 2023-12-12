package main

import (
	"reflect"
	"testing"
)

func Test_day11_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, want: 374},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day11_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day11_puzzle2(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		factor int
		want   int
	}{
		{name: "Example1", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, factor: 10, want: 1030},
		{name: "Example1", input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}, factor: 100, want: 8410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle(tt.input, tt.factor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day11_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
