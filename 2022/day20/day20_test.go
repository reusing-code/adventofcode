package main

import (
	"reflect"
	"testing"
)

func Test_day20_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1",
			"2",
			"-3",
			"3",
			"-2",
			"0",
			"4",
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day20_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day20_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1",
			"2",
			"-3",
			"3",
			"-2",
			"0",
			"4",
		}, want: 1623178306},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day20_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
