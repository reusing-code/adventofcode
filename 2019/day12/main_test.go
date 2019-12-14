package main

import (
	"strings"
	"testing"
)

func Test_calcEnergy(t *testing.T) {
	type args struct {
		input []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{strings.Split(`<x=-1, y=0, z=2>
			<x=2, y=-10, z=-7>
			<x=4, y=-8, z=8>
			<x=3, y=5, z=-1>`, "\n"), 10}, 179},
		{"Example2", args{strings.Split(`<x=-8, y=-10, z=0>
			<x=5, y=5, z=10>
			<x=2, y=-7, z=3>
			<x=9, y=-8, z=-3>`, "\n"), 100}, 1940},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateSystem(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("simulateSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_howManySteps(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example1", args{strings.Split(`<x=-1, y=0, z=2>
			<x=2, y=-10, z=-7>
			<x=4, y=-8, z=8>
			<x=3, y=5, z=-1>`, "\n")}, 2772},
		{"Example2", args{strings.Split(`<x=-8, y=-10, z=0>
			<x=5, y=5, z=10>
			<x=2, y=-7, z=3>
			<x=9, y=-8, z=-3>`, "\n")}, 4686774924},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManySteps(tt.args.in); got != tt.want {
				t.Errorf("howManySteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
