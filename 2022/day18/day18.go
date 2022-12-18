package main

import (
	"fmt"
	"sort"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type coord3d struct {
	x int
	y int
	z int
}

func (c coord3d) add(other coord3d) coord3d {
	return coord3d{c.x + other.x, c.y + other.y, c.z + other.z}
}

var neighbors = []coord3d{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

func puzzle1(input []string) int {
	droplets := make(map[coord3d]bool)
	for _, str := range input {
		c := coord3d{}
		fmt.Sscanf(str, "%d,%d,%d", &c.x, &c.y, &c.z)
		droplets[c] = true
	}

	result := 0
	for key := range droplets {
		for _, c := range neighbors {
			if _, prs := droplets[key.add(c)]; !prs {
				result++
			}
		}
	}
	return result
}

const (
	outside = -1
	inside  = 1
	unknown = 0
	drop    = 2
)

func findPath(from coord3d, to coord3d, cube *[][][]int) bool {
	queue := []coord3d{from}
	visited := make(map[coord3d]bool)
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if (*cube)[current.x][current.y][current.z] != unknown {
			continue
		}
		if _, prs := visited[current]; prs {
			continue
		}
		visited[current] = true
		next := make([]coord3d, 6)
		for i, n := range neighbors {
			next[i] = current.add(n)
		}
		sort.Slice(next, func(i, j int) bool {
			distA := gohelpers.Abs(next[i].x-to.x) + gohelpers.Abs(next[i].y-to.y) + gohelpers.Abs(next[i].z-to.z)
			distB := gohelpers.Abs(next[j].x-to.x) + gohelpers.Abs(next[j].y-to.y) + gohelpers.Abs(next[j].z-to.z)
			return distA < distB
		})
		for _, n := range next {
			if n.x < 0 || n.y < 0 || n.z < 0 || n.x >= len(*cube) || n.y >= len((*cube)[0]) || n.z >= len((*cube)[0][0]) {
				continue
			}
			if n == to {
				for v := range visited {
					(*cube)[v.x][v.y][v.z] = outside
				}
				return true
			}
			if (*cube)[n.x][n.y][n.z] == drop {
				continue
			}
			if (*cube)[n.x][n.y][n.z] == inside {
				for v := range visited {
					(*cube)[v.x][v.y][v.z] = inside
				}
				return false
			}
			if (*cube)[n.x][n.y][n.z] == outside {
				for v := range visited {
					(*cube)[v.x][v.y][v.z] = outside
				}
				return true
			}
			queue = append(queue, n)
		}

	}
	for v := range visited {
		(*cube)[v.x][v.y][v.z] = inside
	}
	return false
}

func puzzle2(input []string) int {
	droplets := make(map[coord3d]bool)
	max := coord3d{0, 0, 0}
	for _, str := range input {
		c := coord3d{}
		fmt.Sscanf(str, "%d,%d,%d", &c.x, &c.y, &c.z)
		droplets[c] = true
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.z > max.z {
			max.z = c.z
		}
	}
	max.x++
	max.y++
	max.z++
	max.x++
	max.y++
	max.z++

	dropletsCube := make([][][]int, max.x)
	for x := range dropletsCube {
		dropletsCube[x] = make([][]int, max.y)
		for y := range dropletsCube[x] {
			dropletsCube[x][y] = make([]int, max.z)
		}
	}

	for k := range droplets {
		dropletsCube[k.x][k.y][k.z] = drop
	}

	zero := coord3d{0, 0, 0}
	dropletsCube[0][0][0] = outside

	result := 0
	for key := range droplets {
		for _, c := range neighbors {
			p := key.add(c)
			if dropletsCube[p.x][p.y][p.z] == outside {
				result++
				continue
			}
			if dropletsCube[p.x][p.y][p.z] == inside {
				continue
			}
			if dropletsCube[p.x][p.y][p.z] == drop {
				continue
			}

			found := findPath(p, zero, &dropletsCube)
			if found {
				result++
			}
		}
	}
	return result
}
