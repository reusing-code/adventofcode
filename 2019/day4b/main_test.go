package main

import "testing"

func Test_validCombination(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want bool
	}{
		{"Example1", 122345, true},
		{"Example2", 322345, false},
		{"Example3", 122343, false},
		{"Example4", 111123, false},
		{"Example5", 135679, false},
		{"Example6", 135677, true},
		{"Example7", 111111, false},
		{"Example8", 223450, false},
		{"Example9", 123789, false},
		{"Example10", 111122, true},
		{"Example11", 112233, true},
		{"Example12", 889999, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := validCombination(tt.in); got != tt.want {
				t.Errorf("validCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
