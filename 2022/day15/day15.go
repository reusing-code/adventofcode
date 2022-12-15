package main

import (
	"fmt"
	"sort"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type coord struct {
	x int
	y int
}

type sensor struct {
	at     coord
	beacon coord
}

func manhattenDistance(a, b coord) int {
	return gohelpers.Abs(a.x-b.x) + gohelpers.Abs(a.y-b.y)
}

func distanceToYline(a coord, y int) int {
	return gohelpers.Abs(a.y - y)
}

func parseSensors(input []string) []*sensor {
	sensors := make([]*sensor, 0, len(input))
	for _, in := range input {
		s := &sensor{}
		fmt.Sscanf(in, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.at.x, &s.at.y, &s.beacon.x, &s.beacon.y)
		sensors = append(sensors, s)
	}
	return sensors
}

func puzzle1(input []string, y int) int {
	sensors := parseSensors(input)

	blocked := make(map[int]bool)
	for _, s := range sensors {
		md := manhattenDistance(s.at, s.beacon)
		dy := distanceToYline(s.at, y)
		if dy > md {
			continue
		}
		diff := md - dy
		for x := s.at.x - diff; x <= s.at.x+diff; x++ {
			blocked[x] = true
		}
	}

	// remove beacons already present
	for _, s := range sensors {
		if s.at.y == y {
			delete(blocked, s.at.x)
		}
		if s.beacon.y == y {
			delete(blocked, s.beacon.x)
		}
	}

	return len(blocked)
}

type interval struct {
	from int
	to   int
}

func puzzle2(input []string, maxCoord int) int {
	sensors := parseSensors(input)
	blocked := make([][]interval, maxCoord+1)
	for i := range blocked {
		blocked[i] = make([]interval, 0)
	}

	for _, s := range sensors {
		md := manhattenDistance(s.at, s.beacon)
		for y := s.at.y - md; y <= s.at.y+md; y++ {
			if y < 0 || y > maxCoord {
				continue
			}
			newmd := md - gohelpers.Abs(s.at.y-y)
			from := s.at.x - newmd
			to := s.at.x + newmd
			from = gohelpers.Clamp(from, 0, maxCoord)
			to = gohelpers.Clamp(to, 0, maxCoord)
			blocked[y] = append(blocked[y], interval{from, to})

		}
	}

	for i, block := range blocked {
		sort.Slice(block, func(a, b int) bool {
			return block[a].from < block[b].from
		})

		current := block[0]
		start := block[0].from
		end := block[0].to
		for j := 1; j < len(block); j++ {
			if block[j].from > current.to {
				return (current.to+1)*4000000 + i
			}
			if block[j].to <= current.to {
				continue
			}
			end = block[j].to
			current = block[j]
		}
		if start != 0 {
			return 0 + i
		}
		if end != maxCoord {
			return maxCoord*4000000 + i
		}

	}
	return 1
}
