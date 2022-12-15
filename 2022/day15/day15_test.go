package main

import (
	"reflect"
	"testing"
)

func Test_day15_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
			"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
			"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
			"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
			"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
			"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
			"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
			"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
			"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
			"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
			"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
			"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
		}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input, 10); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day15_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day15_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
			"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
			"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
			"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
			"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
			"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
			"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
			"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
			"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
			"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
			"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
			"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
		}, want: 56000011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input, 20); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day15_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
