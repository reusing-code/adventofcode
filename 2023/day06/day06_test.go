package main

import (
	"reflect"
	"testing"
)

func Test_day06_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		}, want: 288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day06_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day06_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		}, want: 71503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day06_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
