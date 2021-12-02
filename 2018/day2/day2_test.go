package main

import (
	"reflect"
	"testing"
)

func Test_day2_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day2_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day2_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "Example1", input: []string{"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"fguij",
			"axcye",
			"wvxyz"}, want: "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day2_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
