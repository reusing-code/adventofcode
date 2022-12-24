package main

import (
	"fmt"
	"strconv"

	"github.com/reusing-code/adventofcode/gohelpers"
)

var (
	directions []gohelpers.Coord = []gohelpers.Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirChar    []byte            = []byte{'>', 'v', '<', '^'}
)

type position struct {
	pos gohelpers.Coord
	dir int
}

type move struct {
	val       int
	clockwise bool
}

func printField(field [][]byte) {
	for _, line := range field {
		for _, v := range line {
			fmt.Print(string(v))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func parseMoves(in string) []move {
	number := 0
	result := make([]move, 0)
	for _, c := range in {
		if c == 'R' || c == 'L' {
			if number != 0 {
				result = append(result, move{number, false})
			}
			dir := c == 'R'
			result = append(result, move{0, dir})
			number = 0
		} else {
			n, _ := strconv.Atoi(string(c))
			number = number*10 + n
		}
	}
	if number != 0 {
		result = append(result, move{number, false})
	}
	return result
}

func puzzle1(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	field := gohelpers.ParseCharField(split[0])
	pos := position{gohelpers.Coord{0, 0}, 0}
	for i, v := range field[0] {
		if v == '.' {
			pos.pos.Y = i
			break
		}
	}
	moves := parseMoves(split[1][0])
	for _, move := range moves {
		if move.val == 0 {
			if move.clockwise {
				pos.dir = (pos.dir + 1) % len(directions)
			} else {
				pos.dir = (pos.dir - 1) % len(directions)
				if pos.dir < 0 {
					pos.dir += len(directions)
				}
			}
		} else {
			next := pos.pos
			for i := 0; i < move.val; i++ {
				next = next.Add(directions[pos.dir])
				if directions[pos.dir].X != 0 {
					if next.X < 0 {
						next.X = len(field) - 1
					}
					if next.X >= len(field) {
						next.X = 0
					}
					if next.Y >= len(field[next.X]) {
						i--
						continue
					}
				} else {
					if next.Y < 0 {
						next.Y = len(field[next.X]) - 1
					}
					if next.Y >= len(field[next.X]) {
						next.Y = 0
					}
				}
				if field[next.X][next.Y] == '.' {
					pos.pos = next
					continue
				}
				if field[next.X][next.Y] == '#' {
					break
				}
				if field[next.X][next.Y] == ' ' {
					i--
					continue
				}
			}
		}
	}
	return 1000*(pos.pos.X+1) + 4*(pos.pos.Y+1) + pos.dir
}

func puzzle2(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	field := gohelpers.ParseCharField(split[0])
	pos := position{gohelpers.Coord{0, 0}, 0}
	for i, v := range field[0] {
		if v == '.' {
			pos.pos.Y = i
			break
		}
	}
	moves := parseMoves(split[1][0])
	for _, move := range moves {
		if move.val == 0 {
			if move.clockwise {
				pos.dir = (pos.dir + 1) % len(directions)
			} else {
				pos.dir = (pos.dir - 1) % len(directions)
				if pos.dir < 0 {
					pos.dir += len(directions)
				}
			}
		} else {
			next := pos.pos
			for i := 0; i < move.val; i++ {
				next = next.Add(directions[pos.dir])
				var dir int
				next, dir = mapping(next, pos.dir)
				if field[next.X][next.Y] == '.' {
					pos.pos = next
					pos.dir = dir
					continue
				}
				if field[next.X][next.Y] == '#' {
					break
				}
				panic(fmt.Sprintf("%v %v %v", pos.pos, next, directions[pos.dir]))
			}
		}
	}
	return 1000*(pos.pos.X+1) + 4*(pos.pos.Y+1) + pos.dir
}

func mapping(next gohelpers.Coord, dir int) (gohelpers.Coord, int) {
	rotated := gohelpers.Coord{199 - next.X, 149 - next.Y}
	dirRotated := (dir + 2) % 4
	result, dirResult := mapping2(rotated, dirRotated)
	return gohelpers.Coord{199 - result.X, 149 - result.Y}, (dirResult + 2) % 4
}

// mapping is rotated 180° :(((
func mapping2(next gohelpers.Coord, dir int) (gohelpers.Coord, int) {
	// 1 -> 8
	if next.X == -1 && dir == 3 {
		return gohelpers.Coord{199, next.Y - 50}, 3
	}
	// 2 -> 7
	if next.Y == 150 && dir == 0 && next.X < 50 {
		return gohelpers.Coord{199, next.X + 50}, 3
	}
	// 3 -> 6
	if next.Y == 150 && dir == 0 && next.X >= 50 && next.X < 100 {
		return gohelpers.Coord{(50 - (next.X - 50)) + 149, 99}, 2
	}

	// 4 -> 5
	if next.X == 100 && dir == 1 && next.Y >= 100 {
		return gohelpers.Coord{next.Y, 99}, 2
	}

	// 5 -> 4
	if next.Y == 100 && dir == 0 && next.X >= 100 && next.X < 150 {
		return gohelpers.Coord{99, next.X}, 3
	}

	// 6 -> 3
	if next.Y == 100 && dir == 0 && next.X >= 150 {
		return gohelpers.Coord{(50 - (next.X - 150)) + 49, 149}, 2
	}

	// 7 -> 2
	if next.X == 200 && dir == 1 && next.Y >= 50 && next.Y < 150 {
		return gohelpers.Coord{next.Y - 50, 149}, 2
	}
	// 8 -> 1
	if next.X == 200 && dir == 1 && next.Y < 50 {
		return gohelpers.Coord{0, next.Y + 100}, 1
	}

	// 9 -> 12
	if next.Y == -1 && dir == 2 && next.X >= 150 {
		return gohelpers.Coord{(50 - (next.X - 150)) + 49, 50}, 0
	}

	// 10 -> 11
	if next.X == 149 && dir == 3 && next.Y < 50 {
		return gohelpers.Coord{next.Y + 100, 50}, 0
	}

	// 11 -> 10
	if next.Y == 49 && dir == 2 && next.X >= 100 && next.X < 150 {
		return gohelpers.Coord{150, next.X - 100}, 1
	}

	// 12 -> 9
	if next.Y == 49 && dir == 2 && next.X >= 50 && next.X < 100 {
		return gohelpers.Coord{(50 - (next.X - 50)) + 149, 0}, 0
	}

	// 13 -> 14
	if next.X == 49 && dir == 3 && next.Y >= 50 && next.Y < 100 {
		return gohelpers.Coord{next.Y - 50, 100}, 0
	}

	// 14 -> 13
	if next.Y == 99 && dir == 2 && next.X < 50 {
		return gohelpers.Coord{50, next.X + 50}, 1
	}

	return next, dir
}
