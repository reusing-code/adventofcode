package main

import "testing"

func Test_getClosestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{name: "Example1", input: []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, want: 30},
		{name: "Example2", input: []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, want: 610},
		{name: "Example3", input: []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, want: 410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClosestIntersection(tt.input); got != tt.want {
				t.Errorf("getClosestIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
