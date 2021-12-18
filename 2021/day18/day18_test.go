package main

import (
	"reflect"
	"testing"
)

func Test_day18_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Magnitude1", input: []string{
			"[[1,2],[[3,4],5]]",
		}, want: 143},
		{name: "Magnitude2", input: []string{
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		}, want: 3488},
		{name: "Example1", input: []string{
			"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			"[[[5,[2,8]],4],[5,[[9,9],0]]]",
			"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
			"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
			"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
			"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
			"[[[[5,4],[7,7]],8],[[8,3],8]]",
			"[[9,3],[[9,9],[6,[4,9]]]]",
			"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
		}, want: 4140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day18_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day18_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			"[[[5,[2,8]],4],[5,[[9,9],0]]]",
			"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
			"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
			"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
			"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
			"[[[[5,4],[7,7]],8],[[8,3],8]]",
			"[[9,3],[[9,9],[6,[4,9]]]]",
			"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
		}, want: 3993},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day18_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
