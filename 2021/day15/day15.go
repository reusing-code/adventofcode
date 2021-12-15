package main

import (
	"sort"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type coord struct {
	x int
	y int
}

func puzzle1(input []string) int {
	field := gohelpers.ParseIntField(input)
	distances := make([][]int, len(input))
	for i, _ := range distances {
		distances[i] = make([]int, len(input[i]))
	}

	type queueEntry struct {
		val int
		c   coord
	}

	priorityQueue := make([]queueEntry, 0)
	priorityQueue = append(priorityQueue, queueEntry{0, coord{0, 0}})

	for len(priorityQueue) > 0 {
		current := priorityQueue[0]
		priorityQueue = priorityQueue[1:]
		if distances[current.c.x][current.c.y] > 0 {
			continue
		}
		distances[current.c.x][current.c.y] = current.val
		if current.c.x == len(field)-1 && current.c.y == len(field[current.c.x])-1 {
			return current.val
		}

		check := func(x int, y int) {
			if x < 0 || x >= len(field) {
				return
			}
			if y < 0 || y >= len(field[0]) {
				return
			}
			if x == 0 && y == 0 {
				return
			}
			if distances[x][y] > 0 {
				return
			}
			priorityQueue = append(priorityQueue, queueEntry{current.val + field[x][y], coord{x, y}})
		}

		check(current.c.x-1, current.c.y)
		check(current.c.x+1, current.c.y)
		check(current.c.x, current.c.y-1)
		check(current.c.x, current.c.y+1)
		sort.Slice(priorityQueue, func(i int, j int) bool {
			return priorityQueue[i].val < priorityQueue[j].val
		})
	}

	return 1
}

func puzzle2(input []string) int {
	field := gohelpers.ParseIntField(input)
	distances := make([][]int, len(input)*5)
	for i, _ := range distances {
		distances[i] = make([]int, len(input[0])*5)
	}

	type queueEntry struct {
		val int
		c   coord
	}

	priorityQueue := make([]queueEntry, 0)
	priorityQueue = append(priorityQueue, queueEntry{0, coord{0, 0}})

	for len(priorityQueue) > 0 {
		current := priorityQueue[0]
		priorityQueue = priorityQueue[1:]
		if distances[current.c.x][current.c.y] > 0 {
			continue
		}
		distances[current.c.x][current.c.y] = current.val
		if current.c.x == len(field)*5-1 && current.c.y == len(field[0])*5-1 {
			return current.val
		}

		check := func(x int, y int) {
			if x < 0 || x >= len(field)*5 {
				return
			}
			if y < 0 || y >= len(field[0])*5 {
				return
			}
			if x == 0 && y == 0 {
				return
			}
			if distances[x][y] > 0 {
				return
			}
			x1 := x % len(field)
			y1 := y % len(field[0])
			add := int(x / len(field))
			add += int(y / len(field[0]))
			dist := field[x1][y1] + add
			if dist > 9 {
				dist = dist - 9
			}
			priorityQueue = append(priorityQueue, queueEntry{current.val + dist, coord{x, y}})
		}

		check(current.c.x-1, current.c.y)
		check(current.c.x+1, current.c.y)
		check(current.c.x, current.c.y-1)
		check(current.c.x, current.c.y+1)
		sort.Slice(priorityQueue, func(i int, j int) bool {
			return priorityQueue[i].val < priorityQueue[j].val
		})
	}

	return 1
}
