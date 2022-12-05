package main

import (
	"reflect"
	"testing"
)

func Test_day05_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "Example1", input: []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}, want: "CMZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day05_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day05_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "Example1", input: []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}, want: "MCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day05_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
