package main

import "testing"

func Test_calcModuleFuel(t *testing.T) {

	tests := []struct {
		name string
		mass int
		want int
	}{
		{name: "Example 1", mass: 12, want: 2},
		{name: "Example 2", mass: 14, want: 2},
		{name: "Example 3", mass: 1969, want: 654},
		{name: "Example 4", mass: 100756, want: 33583},
		{name: "Example 5", mass: 6, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcModuleFuel(tt.mass); got != tt.want {
				t.Errorf("calcModuleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcModuleFuelAdvanced(t *testing.T) {

	tests := []struct {
		name string
		mass int
		want int
	}{
		{name: "Example 1", mass: 14, want: 2},
		{name: "Example 2", mass: 1969, want: 966},
		{name: "Example 3", mass: 100756, want: 50346},
		{name: "Example 4", mass: 0, want: 0},
		{name: "Example 4", mass: -5000, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcModuleFuelAdvanced(tt.mass); got != tt.want {
				t.Errorf("calcModuleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
