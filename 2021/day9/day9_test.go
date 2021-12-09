package main

import (
	"reflect"
	"testing"
)

func Test_day9_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678"}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day9_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day9_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678"}, want: 1134},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day9_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
