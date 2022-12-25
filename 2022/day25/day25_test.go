package main

import (
	"reflect"
	"testing"
)

func Test_day25_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: "Example1", input: []string{
			"1=-0-2",
			"12111",
			"2=0=",
			"21",
			"2=01",
			"111",
			"20012",
			"112",
			"1=-1=",
			"1-12",
			"12",
			"1=",
			"122",
		}, want: "2=-1=0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day25_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day25_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"1"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day25_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_snafu2decimal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "1", input: "1=-0-2", want: 1747},
		{name: "2", input: "12111", want: 906},
		{name: "3", input: "2=0=", want: 198},
		{name: "4", input: "21", want: 11},
		{name: "5", input: "2=01", want: 201},
		{name: "6", input: "111", want: 31},
		{name: "7", input: "20012", want: 1257},
		{name: "8", input: "112", want: 32},
		{name: "9", input: "1=-1=", want: 353},
		{name: "10", input: "1-12", want: 107},
		{name: "11", input: "12", want: 7},
		{name: "12", input: "1=", want: 3},
		{name: "13", input: "122", want: 37},
		{name: "14", input: "1", want: 1},
		{name: "15", input: "2", want: 2},
		{name: "16", input: "1=", want: 3},
		{name: "17", input: "1-", want: 4},
		{name: "18", input: "10", want: 5},
		{name: "19", input: "11", want: 6},
		{name: "20", input: "12", want: 7},
		{name: "21", input: "2=", want: 8},
		{name: "22", input: "2-", want: 9},
		{name: "23", input: "20", want: 10},
		{name: "24", input: "1=0", want: 15},
		{name: "24", input: "1-0", want: 20},
		{name: "24", input: "1=11-2", want: 2022},
		{name: "24", input: "1-0---0", want: 12345},
		{name: "24", input: "1121-1110-1=0", want: 314159265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snafu2decimal(tt.input); got != tt.want {
				t.Errorf("snafu2decimal() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimal2snafu(tt.want); got != tt.input {
				t.Errorf("decimal2snafu() = %v, want %v", got, tt.input)
			}
		})
	}
}
