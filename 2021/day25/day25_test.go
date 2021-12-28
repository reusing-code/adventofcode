package main

import (
	"reflect"
	"testing"
)

func Test_day25_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		// {name: "Example0", input: []string{
		// 	"...>...",
		// 	".......",
		// 	"......>",
		// 	"v.....>",
		// 	"......>",
		// 	".......",
		// 	"..vvv..",
		// }, want: 58},
		{name: "Example1", input: []string{
			"v...>>.vv>",
			".vv>>.vv..",
			">>.>v>...v",
			">>v>>.>.v.",
			"v>v.vv.v..",
			">.>>..v...",
			".vv..>.>v.",
			"v.v..>>v.v",
			"....v..v.>",
		}, want: 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day25_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day25_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"1"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day25_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
