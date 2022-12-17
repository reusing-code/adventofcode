package main

import "fmt"

type coord struct {
	x int
	y int
}

type shape struct {
	size   coord
	points []coord
}

var shapes []*shape = createShapes()

func createShapes() []*shape {
	result := make([]*shape, 5)

	result[0] = &shape{coord{1, 4}, []coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}}
	result[1] = &shape{coord{3, 3}, []coord{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}}
	result[2] = &shape{coord{3, 3}, []coord{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}}
	result[3] = &shape{coord{4, 1}, []coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}}
	result[4] = &shape{coord{2, 2}, []coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}}}
	return result
}

func printCavern(cavern [][7]bool) {
	for i_ := range cavern {
		i := len(cavern) - i_ - 1
		fmt.Print("|")
		for _, v := range cavern[i] {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println("")
	}
	fmt.Println("")
}

func dropShape(cavern *[][7]bool, s *shape, maxHeight int, jetPattern string, patternIdx *int) int {
	sPos := coord{maxHeight + 4, 2}
	for {

		isBlocked := func(pos coord) bool {
			for _, v := range s.points {
				p := coord{pos.x + v.x, pos.y + v.y}
				if p.x < 0 {
					return true
				}
				if p.y < 0 || p.y > 6 {
					return true
				}
				if p.x > maxHeight {
					continue
				}
				if (*cavern)[p.x][p.y] {
					return true
				}
			}
			return false
		}
		// jet push
		dir := 1
		if jetPattern[*patternIdx] == '<' {
			dir = -1
		}
		*patternIdx = (*patternIdx + 1) % len(jetPattern)

		if !isBlocked(coord{sPos.x, sPos.y + dir}) {
			sPos.y += dir
		}
		// drop

		if !isBlocked(coord{sPos.x - 1, sPos.y}) {
			sPos.x--
		} else {
			currentMax := maxHeight
			for newHeight := currentMax + 1; newHeight <= sPos.x+s.size.x-1; newHeight++ {
				maxHeight = newHeight
				*cavern = append(*cavern, [7]bool{})
			}
			for _, p := range s.points {
				(*cavern)[sPos.x+p.x][sPos.y+p.y] = true
			}
			return maxHeight
		}
	}
}

func puzzle1(input []string) int {
	jetPattern := input[0]
	patternIdx := 0
	maxHeight := 0

	cavern := make([][7]bool, 1)
	cavern[0] = [7]bool{true, true, true, true, true, true, true}

	for i := 0; i < 2022; i++ {
		s := shapes[i%len(shapes)]

		maxHeight = dropShape(&cavern, s, maxHeight, jetPattern, &patternIdx)
	}
	return maxHeight
}

type cacheEntry struct {
	coords    []coord
	jetIdx    int
	maxHeight int
	i         int
}

func makeCacheEntry(cavern [][7]bool, jet int, maxHeight int, i int, cacheLines int) *cacheEntry {
	result := &cacheEntry{make([]coord, 0, cacheLines*7), jet, maxHeight, i}
	for x := len(cavern) - cacheLines; x < len(cavern); x++ {
		for y := 0; y < 7; y++ {
			if cavern[x][y] {
				result.coords = append(result.coords, coord{x - (len(cavern) - cacheLines), y})
			}
		}
	}
	return result
}

func cmp(a []coord, b []coord) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func puzzle2(input []string) int {
	jetPattern := input[0]
	patternIdx := 0
	maxHeight := 0

	cavern := make([][7]bool, 1)
	cavern[0] = [7]bool{true, true, true, true, true, true, true}

	cache := make([]*cacheEntry, 0)
	cacheLines := 20

	period := 0
	heightPerPeriod := 0

	for i := 0; i < 1000000000000; i++ {
		s := shapes[i%len(shapes)]
		maxHeight = dropShape(&cavern, s, maxHeight, jetPattern, &patternIdx)
		if period > 0 {
			if (1000000000000-(i+1))%period == 0 {
				return maxHeight + ((1000000000000-(i+1))/period)*heightPerPeriod
			}
		} else {
			if maxHeight > cacheLines && i%5 == 0 {
				newEntry := makeCacheEntry(cavern, patternIdx, maxHeight, i, cacheLines)

				for k := len(cache) - 1; k > 0; k-- {
					if cache[k].jetIdx != newEntry.jetIdx {
						continue
					}
					if cmp(cache[k].coords, newEntry.coords) {
						period = i - cache[k].i
						heightPerPeriod = maxHeight - cache[k].maxHeight
					}
				}

				cache = append(cache, newEntry)
			}
		}
	}
	return maxHeight
}
