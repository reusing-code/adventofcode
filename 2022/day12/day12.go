package main

import (
	"container/heap"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func manhatten(start, goal gohelpers.Coord) int {
	return gohelpers.Abs(start.X-goal.X) + gohelpers.Abs(start.Y-goal.Y)
}

func constructPath(cameFrom map[gohelpers.Coord]gohelpers.Coord, goal gohelpers.Coord) []gohelpers.Coord {
	current := goal
	result := []gohelpers.Coord{goal}
	for {
		if _, prs := cameFrom[current]; !prs {
			return result
		}
		current = cameFrom[current]
		result = append(result, current)
	}
}

func astar(start, goal gohelpers.Coord, field [][]byte) []gohelpers.Coord {
	height := len(field)
	width := len(field[0])

	pq := make(gohelpers.PriorityQueue, 1)
	pq[0] = &gohelpers.Item{Value: start, Priority: 0, Index: 0}
	heap.Init(&pq)

	cameFrom := make(map[gohelpers.Coord]gohelpers.Coord)
	gScore := make(map[gohelpers.Coord]int)
	gScore[start] = 0
	fScore := make(map[gohelpers.Coord]int)
	fScore[start] = manhatten(start, goal)

	for {
		if len(pq) == 0 {
			return nil
		}
		current := heap.Pop(&pq).(*gohelpers.Item)
		if current.Value == goal {
			return constructPath(cameFrom, current.Value)
		}
		for i := 0; i < 4; i++ {
			next := gohelpers.Coord{X: current.Value.X, Y: current.Value.Y}
			switch i {
			case 0:
				// left
				next.Y -= 1
			case 1:
				// right
				next.Y += 1
			case 2:
				// down
				next.X -= 1
			case 3:
				// up
				next.X += 1
			}
			if next.X < 0 || next.Y < 0 || next.X >= height || next.Y >= width {
				continue
			}
			if field[next.X][next.Y] > field[current.Value.X][current.Value.Y]+1 {
				continue
			}
			score := gScore[current.Value] + 1
			_, prs := gScore[next]
			if !prs || score < gScore[next] {
				cameFrom[next] = current.Value
				gScore[next] = score
				fScore[next] = gScore[current.Value] + manhatten(next, goal)
				found := false
				for i := range pq {
					if pq[i].Value == next {
						found = true
						pq[i].Priority = fScore[next]
						heap.Fix(&pq, pq[i].Index)
					}
				}
				if !found {
					heap.Push(&pq, &gohelpers.Item{Value: next, Priority: fScore[next]})
				}
			}
		}
	}
}

func puzzle1(input []string) int {
	charField := gohelpers.ParseCharField(input)

	start := gohelpers.Coord{0, 0}
	goal := gohelpers.Coord{0, 0}
	for x := range charField {
		for y := range charField[x] {
			if charField[x][y] == 'S' {
				start = gohelpers.Coord{x, y}
				charField[x][y] = 'a'
			}
			if charField[x][y] == 'E' {
				goal = gohelpers.Coord{x, y}
				charField[x][y] = 'z'
			}
		}
	}
	result := astar(start, goal, charField)
	return len(result) - 1
}

func puzzle2(input []string) int {
	charField := gohelpers.ParseCharField(input)

	possibleStart := make([]gohelpers.Coord, 0)
	goal := gohelpers.Coord{0, 0}
	for x := range charField {
		for y := range charField[x] {
			if charField[x][y] == 'S' {
				charField[x][y] = 'a'
			}
			if charField[x][y] == 'E' {
				goal = gohelpers.Coord{x, y}
				charField[x][y] = 'z'
			}
			if charField[x][y] == 'a' {
				possibleStart = append(possibleStart, gohelpers.Coord{x, y})
			}
		}
	}
	min := len(charField) * len(charField[0])
	for _, start := range possibleStart {
		res := astar(start, goal, charField)
		if res == nil {
			continue
		}
		if len(res)-1 < min {
			min = len(res) - 1
		}

	}
	return min
}
