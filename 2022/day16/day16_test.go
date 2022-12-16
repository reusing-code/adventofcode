package main

import (
	"reflect"
	"testing"
)

func Test_day16_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
			"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
			"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
			"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
			"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
			"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
			"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
			"Valve HH has flow rate=22; tunnel leads to valve GG",
			"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
			"Valve JJ has flow rate=21; tunnel leads to valve II",
		}, want: 1651},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day16_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day16_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
			"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
			"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
			"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
			"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
			"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
			"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
			"Valve HH has flow rate=22; tunnel leads to valve GG",
			"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
			"Valve JJ has flow rate=21; tunnel leads to valve II",
		}, want: 1707},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day16_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
