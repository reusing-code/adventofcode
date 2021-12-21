package main

import (
	"reflect"
	"testing"
)

func Test_day21_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Player 1 starting position: 4",
			"Player 2 starting position: 8",
		}, want: 739785},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day21_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day21_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Player 1 starting position: 4",
			"Player 2 starting position: 8",
		}, want: 444356092776315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day21_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
