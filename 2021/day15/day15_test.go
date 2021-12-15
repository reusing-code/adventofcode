package main

import (
	"reflect"
	"testing"
)

func Test_day15_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1163751742",
			"1381373672",
			"2136511328",
			"3694931569",
			"7463417111",
			"1319128137",
			"1359912421",
			"3125421639",
			"1293138521",
			"2311944581",
		}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day15_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day15_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1163751742",
			"1381373672",
			"2136511328",
			"3694931569",
			"7463417111",
			"1319128137",
			"1359912421",
			"3125421639",
			"1293138521",
			"2311944581",
		}, want: 315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day15_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
