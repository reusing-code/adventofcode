package main

import (
	"reflect"
	"testing"
)

func Test_day01_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}, want: 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day01_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day01_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}, want: 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day01_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
