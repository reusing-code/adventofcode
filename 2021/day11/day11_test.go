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
			"5483143223",
			"2745854711",
			"5264556173",
			"6141336146",
			"6357385478",
			"4167524645",
			"2176841721",
			"6882881134",
			"4846848554",
			"5283751526",
		}, want: 1656},
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
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"5483143223",
			"2745854711",
			"5264556173",
			"6141336146",
			"6357385478",
			"4167524645",
			"2176841721",
			"6882881134",
			"4846848554",
			"5283751526",
		}, want: 195},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day11_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
