package main

import (
	"reflect"
	"testing"
)

func Test_day13_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			"",
			"[[1],[2,3,4]]",
			"[[1],4]",
			"",
			"[9]",
			"[[8,7,6]]",
			"",
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			"",
			"[7,7,7,7]",
			"[7,7,7]",
			"",
			"[]",
			"[3]",
			"",
			"[[[]]]",
			"[[]]",
			"",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day13_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day13_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			"",
			"[[1],[2,3,4]]",
			"[[1],4]",
			"",
			"[9]",
			"[[8,7,6]]",
			"",
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			"",
			"[7,7,7,7]",
			"[7,7,7]",
			"",
			"[]",
			"[3]",
			"",
			"[[[]]]",
			"[[]]",
			"",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		}, want: 140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day13_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
