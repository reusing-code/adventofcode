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
		{name: "Example1", input: []string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 7},
		{name: "Example2", input: []string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 5},
		{name: "Example3", input: []string{"nppdvjthqldpwncqszvftbrmjlhg"}, want: 6},
		{name: "Example4", input: []string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 10},
		{name: "Example5", input: []string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 11},
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
		{name: "Example1", input: []string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 19},
		{name: "Example2", input: []string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 23},
		{name: "Example3", input: []string{"nppdvjthqldpwncqszvftbrmjlhg"}, want: 23},
		{name: "Example4", input: []string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 29},
		{name: "Example5", input: []string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day06_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
