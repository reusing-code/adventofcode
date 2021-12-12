package main

import (
	"reflect"
	"testing"
)

func Test_day12_puzzle1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}, want: 10},
		{name: "Example2", input: []string{
			"dc-end",
			"HN-start",
			"start-kj",
			"dc-start",
			"dc-HN",
			"LN-dc",
			"HN-end",
			"kj-sa",
			"kj-HN",
			"kj-dc",
		}, want: 19},
		{name: "Example3", input: []string{
			"fs-end",
			"he-DX",
			"fs-he",
			"start-DX",
			"pj-DX",
			"end-zg",
			"zg-sl",
			"zg-pj",
			"pj-he",
			"RW-he",
			"fs-DX",
			"pj-RW",
			"zg-RW",
			"start-pj",
			"he-WI",
			"zg-he",
			"pj-fs",
			"start-RW",
		}, want: 226},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day12_puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day12_puzzle2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}, want: 36},
		{name: "Example2", input: []string{
			"dc-end",
			"HN-start",
			"start-kj",
			"dc-start",
			"dc-HN",
			"LN-dc",
			"HN-end",
			"kj-sa",
			"kj-HN",
			"kj-dc",
		}, want: 103},
		{name: "Example3", input: []string{
			"fs-end",
			"he-DX",
			"fs-he",
			"start-DX",
			"pj-DX",
			"end-zg",
			"zg-sl",
			"zg-pj",
			"pj-he",
			"RW-he",
			"fs-DX",
			"pj-RW",
			"zg-RW",
			"start-pj",
			"he-WI",
			"zg-he",
			"pj-fs",
			"start-RW",
		}, want: 3509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day12_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
