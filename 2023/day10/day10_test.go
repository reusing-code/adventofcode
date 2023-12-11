package main

import (
	"reflect"
	"testing"
)

func Test_day10_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			".....",
			".S-7.",
			".|.|.",
			".L-J.",
			".....",
		}, want: 4},
		{name: "Example1", input: []string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}, want: 4},
		{name: "Example1", input: []string{
			"..F7.",
			".FJ|.",
			"SJ.L7",
			"|F--J",
			"LJ...",
		}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day10_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day10_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		}, want: 4},
		{name: "Example1", input: []string{
			"..........",
			".S------7.",
			".|F----7|.",
			".||....||.",
			".||....||.",
			".|L-7F-J|.",
			".|..||..|.",
			".L--JL--J.",
			"..........",
		}, want: 4},

		{name: "Example1", input: []string{
			".F----7F7F7F7F-7....",
			".|F--7||||||||FJ....",
			".||.FJ||||||||L7....",
			"FJL7L7LJLJ||LJ.L-7..",
			"L--J.L7...LJS7F-7L7.",
			"....F-J..F7FJ|L7L7L7",
			"....L7.F7||L7|.L7L7|",
			".....|FJLJ|FJ|F7|.LJ",
			"....FJL-7.||.||||...",
			"....L---J.LJ.LJLJ...",
		}, want: 8},
		{name: "Example1", input: []string{
			"FF7FSF7F7F7F7F7F---7",
			"L|LJ||||||||||||F--J",
			"FL-7LJLJ||||||LJL-77",
			"F--JF--7||LJLJ7F7FJ-",
			"L---JF-JLJ.||-FJLJJ7",
			"|F|F-JF---7F7-L7L|7|",
			"|FFJF7L7F-JF7|JL---7",
			"7-L-JL7||F7|L7F-7F7|",
			"L.L7LFJ|||||FJL7||LJ",
			"L7JLJL-JLJLJL--JLJ.L",
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day10_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
