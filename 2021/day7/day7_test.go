package main

import (
	"reflect"
	"testing"
)

func Test_day7_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"16,1,2,0,4,2,7,1,2,14"}, want: 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day7_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day7_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"16,1,2,0,4,2,7,1,2,14"}, want: 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day7_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
