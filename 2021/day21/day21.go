package main

import (
	"strconv"
)

type random struct {
	current int
	max     int
	count   int
}

func (r *random) getNext() (res int) {
	res = r.current
	r.count++
	r.current++
	if r.current > r.max {
		r.current = 1
	}
	return
}

func puzzle1(input []string) int {
	r := &random{1, 100, 0}
	p1, err := strconv.Atoi(input[0][len(input[0])-1:])
	if err != nil {
		panic(err)
	}
	p1--
	p2, err := strconv.Atoi(input[1][len(input[0])-1:])
	if err != nil {
		panic(err)
	}
	p2--
	p1points := 0
	p2points := 0

	for {
		p1 += r.getNext()
		p1 += r.getNext()
		p1 += r.getNext()
		p1 = p1 % 10
		p1points += p1 + 1
		if p1points >= 1000 {
			return p2points * r.count
		}
		p2 += r.getNext()
		p2 += r.getNext()
		p2 += r.getNext()
		p2 = p2 % 10
		p2points += p2 + 1
		if p2points >= 1000 {
			return p1points * r.count
		}
	}
	return 0
}

type state struct {
	p1            int
	p2            int
	p1points      int
	p2points      int
	possibilities int
}

var pos = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

func doStep(s state, p1next bool) (int, int) {
	if s.p1points >= 21 {
		return s.possibilities, 0
	}
	if s.p2points >= 21 {
		return 0, s.possibilities
	}
	p1win := 0
	p2win := 0

	if p1next {
		for key, val := range pos {
			p1 := (s.p1 + key) % 10
			p1points := s.p1points + p1 + 1
			possibilities := s.possibilities * val
			p1tmp, p2tmp := doStep(state{p1, s.p2, p1points, s.p2points, possibilities}, false)
			p1win += p1tmp
			p2win += p2tmp
		}
	} else {
		for key, val := range pos {
			p2 := (s.p2 + key) % 10
			p2points := s.p2points + p2 + 1
			possibilities := s.possibilities * val
			p1tmp, p2tmp := doStep(state{s.p1, p2, s.p1points, p2points, possibilities}, true)
			p1win += p1tmp
			p2win += p2tmp
		}
	}
	return p1win, p2win
}

func puzzle2(input []string) int {
	p1, err := strconv.Atoi(input[0][len(input[0])-1:])
	if err != nil {
		panic(err)
	}
	p1--
	p2, err := strconv.Atoi(input[1][len(input[0])-1:])
	if err != nil {
		panic(err)
	}
	p2--
	p1win, p2win := doStep(state{p1, p2, 0, 0, 1}, true)
	if p1win > p2win {
		return p1win
	}
	return p2win
}
