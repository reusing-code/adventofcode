package main

import (
	"reflect"
	"testing"
)

func Test_intcode(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Example1", input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, want: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{name: "Example1", input: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{name: "Example1", input: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{name: "Example1", input: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{name: "Example1", input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intcode(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
