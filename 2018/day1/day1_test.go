package main

import (
	"reflect"
	"testing"
)

func Test_day1_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"+1", "-2", "+3", "+1"}, want: 3},
		{name: "Example2", input: []string{"+1", "+1", "+1"}, want: 3},
		{name: "Example3", input: []string{"+1", "+1", "-2"}, want: 0},
		{name: "Example4", input: []string{"-1", "-2", "-3"}, want: -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day1_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day1_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"+1", "-1"}, want: 0},
		{name: "Example2", input: []string{"+3", "+3", "+4", "-2", "-4"}, want: 10},
		{name: "Example3", input: []string{"-6", "+3", "+8", "+5", "-6"}, want: 5},
		{name: "Example4", input: []string{"+7", "+7", "-2", "-7", "-4"}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day1_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
