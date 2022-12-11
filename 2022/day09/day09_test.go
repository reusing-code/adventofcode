package main

import (
	"reflect"
	"testing"
)

func Test_day09_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day09_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day09_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		}, want: 1},
		{name: "Example1", input: []string{
			"R 5",
			"U 8",
			"L 8",
			"D 3",
			"R 17",
			"D 10",
			"L 25",
			"U 20",
		}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day09_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
