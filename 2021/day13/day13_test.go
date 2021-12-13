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
			"6,10",
			"0,14",
			"9,10",
			"0,3",
			"10,4",
			"4,11",
			"6,0",
			"6,12",
			"4,1",
			"0,13",
			"10,12",
			"3,4",
			"3,0",
			"8,4",
			"1,10",
			"2,14",
			"8,10",
			"9,0",
			"",
			"fold along y=7",
			"fold along x=5",
		}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day13_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}
