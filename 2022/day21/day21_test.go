package main

import (
	"reflect"
	"testing"
)

func Test_day21_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"root: pppw + sjmn",
			"dbpl: 5",
			"cczh: sllz + lgvd",
			"zczc: 2",
			"ptdq: humn - dvpt",
			"dvpt: 3",
			"lfqf: 4",
			"humn: 5",
			"ljgn: 2",
			"sjmn: drzm * dbpl",
			"sllz: 4",
			"pppw: cczh / lfqf",
			"lgvd: ljgn * ptdq",
			"drzm: hmdt - zczc",
			"hmdt: 32",
		}, want: 152},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day21_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day21_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"root: pppw + sjmn",
			"dbpl: 5",
			"cczh: sllz + lgvd",
			"zczc: 2",
			"ptdq: humn - dvpt",
			"dvpt: 3",
			"lfqf: 4",
			"humn: 5",
			"ljgn: 2",
			"sjmn: drzm * dbpl",
			"sllz: 4",
			"pppw: cczh / lfqf",
			"lgvd: ljgn * ptdq",
			"drzm: hmdt - zczc",
			"hmdt: 32",
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day21_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
