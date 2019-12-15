package main

import "testing"

import "strings"

const ex1 string = `.#..#
.....
#####
....#
...##`

const ex2 string = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

const ex3 string = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

const ex4 string = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

const ex5 string = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

func Test_calculateMostVisible(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{strings.Split(ex1, "\n")}, 8},
		{"Example 2", args{strings.Split(ex2, "\n")}, 33},
		{"Example 3", args{strings.Split(ex3, "\n")}, 35},
		{"Example 4", args{strings.Split(ex4, "\n")}, 41},
		{"Example 5", args{strings.Split(ex5, "\n")}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMostVisible(tt.args.input); got != tt.want {
				t.Errorf("calculateMostVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}
