package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type station struct {
	x, y int
}

func isBlockingLineOfSight(s1, s2, block station) bool {
	if s1.x < s2.x {
		if block.x < s1.x || block.x > s2.x {
			return false
		}
	} else {
		if block.x < s2.x || block.x > s1.x {
			return false
		}
	}
	if s1.y < s2.y {
		if block.y < s1.y || block.y > s2.y {
			return false
		}
	} else {
		if block.y < s2.y || block.y > s1.y {
			return false
		}
	}

	x1 := s2.x - s1.x
	y1 := s2.y - s1.y
	x2 := block.x - s1.x
	y2 := block.y - s1.y
	if x2 == 0 {
		return x1 == 0
	}
	if x1*y2%x2 != 0 {
		return false
	}
	return x1*y2/x2 == y1
}

func sortStations(sut station, all []station) {
	sort.Slice(all, func(i int, j int) bool {
		xiSqr := all[i].x - sut.x
		xiSqr = xiSqr * xiSqr
		yiSqr := all[i].y - sut.y
		yiSqr = yiSqr * yiSqr
		xjSqr := all[j].x - sut.x
		xjSqr = xjSqr * xjSqr
		yjSqr := all[j].y - sut.y
		yjSqr = yjSqr * yjSqr
		return xiSqr+yiSqr < xjSqr+yjSqr
	})
}

func sortByQuadrant(sut station, seen []int, all []station) {
	sort.Slice(seen, func(i int, j int) bool {
		xi := all[seen[i]].x - sut.x
		yi := -all[seen[i]].y + sut.y
		xj := all[seen[j]].x - sut.x
		yj := -all[seen[j]].y + sut.y
		ai := math.Atan2(float64(xi), float64(yi))
		if ai < 0 {
			ai += math.Pi * 2
		}
		aj := math.Atan2(float64(xj), float64(yj))
		if aj < 0 {
			aj += math.Pi * 2
		}
		return ai < aj
	})
}

func calcVisible(sut station, all []station) (int, []int) {
	sortStations(sut, all)
	seen := make([]int, 0)
	for i := 1; i < len(all); i++ {
		add := true
		for _, blocking := range seen {
			if isBlockingLineOfSight(sut, all[i], all[blocking]) {
				add = false
				break
			}
		}
		if add {
			seen = append(seen, i)
		}
	}
	return len(seen), seen
}

func calculateMostVisible(input []string) int {
	stations := make([]station, 0)

	for y := 0; y < len(input); y++ {
		line := strings.Split(input[y], "")
		for x := 0; x < len(line); x++ {
			if line[x] == "#" {
				stations = append(stations, station{x, y})
			}
		}
	}

	maxVisible := 0
	var maxStation station
	for _, sut := range stations {
		visible, _ := calcVisible(sut, append(make([]station, 0), stations...))
		if visible > maxVisible {
			maxVisible = visible
			maxStation = sut
		}
	}
	if len(stations) > 200 {
		destroyStations(maxStation, stations)
	}
	return maxVisible
}

func destroyStations(sut station, all []station) {
	destroyed := 0
	for {
		count, seen := calcVisible(sut, all)
		if destroyed+count < 200 {
			destroyed += count
			destMap := make(map[int]bool)
			for _, v := range seen {
				destMap[v] = true
			}
			remaining := make([]station, 0, len(all)-count)
			for i, v := range all {
				if _, ok := destMap[i]; !ok {
					remaining = append(remaining, v)
				}
			}
			all = remaining
			continue
		} else {
			sortByQuadrant(sut, seen, all)
			fmt.Printf("\n200th Station: %v\n", all[seen[200-destroyed-1]])
			break
		}
	}
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	distance := calculateMostVisible(input)
	fmt.Printf("Result %v\n", distance)

}
