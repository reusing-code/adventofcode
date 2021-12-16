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
		{name: "Example1", input: []string{"8A004A801A8002F478"}, want: 16},
		{name: "Example2", input: []string{"620080001611562C8802118E34"}, want: 12},
		{name: "Example3", input: []string{"C0015000016115A2E0802F182340"}, want: 23},
		{name: "Example4", input: []string{"A0016C880162017C3686B18A3D4780"}, want: 31},
		{name: "Literal", input: []string{"D2FE28"}, want: 6},
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
		{name: "Sum", input: []string{"C200B40A82"}, want: 3},
		{name: "Product", input: []string{"04005AC33890"}, want: 54},
		{name: "Minimum", input: []string{"880086C3E88112"}, want: 7},
		{name: "Maximum", input: []string{"CE00C43D881120"}, want: 9},
		{name: "Less", input: []string{"D8005AC2A8F0"}, want: 1},
		{name: "Not greater", input: []string{"F600BC2D8F"}, want: 0},
		{name: "Not equal", input: []string{"9C005AC2F8F0"}, want: 0},
		{name: "Complex", input: []string{"9C0141080250320F1802104A08"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := puzzle2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test_day16_puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
