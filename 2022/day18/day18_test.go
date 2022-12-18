package main

import (
	"reflect"
	"testing"
)

func Test_day18_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1,1,1",
			"2,1,1",
		}, want: 10},
		{name: "Example2", input: []string{
			"2,2,2",
			"1,2,2",
			"3,2,2",
			"2,1,2",
			"2,3,2",
			"2,2,1",
			"2,2,3",
			"2,2,4",
			"2,2,6",
			"1,2,5",
			"3,2,5",
			"2,1,5",
			"2,3,5",
		}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day18_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day18_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example2", input: []string{
			"2,2,2",
			"1,2,2",
			"3,2,2",
			"2,1,2",
			"2,3,2",
			"2,2,1",
			"2,2,3",
			"2,2,4",
			"2,2,6",
			"1,2,5",
			"3,2,5",
			"2,1,5",
			"2,3,5",
		}, want: 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day18_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
