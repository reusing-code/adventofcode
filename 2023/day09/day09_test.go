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
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}, want: 114},
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
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day09_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
