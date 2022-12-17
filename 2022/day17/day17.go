package main

import (
	"fmt"
	"image/color"

	"github.com/reusing-code/adventofcode/gohelpers"
	"github.com/reusing-code/adventofcode/gohelpers/image"
)

type shape struct {
	size   gohelpers.Coord
	points []gohelpers.Coord
}

var shapes []*shape = createShapes()

func createShapes() []*shape {
	result := make([]*shape, 5)

	result[0] = &shape{gohelpers.Coord{1, 4}, []gohelpers.Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}}
	result[1] = &shape{gohelpers.Coord{3, 3}, []gohelpers.Coord{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}}
	result[2] = &shape{gohelpers.Coord{3, 3}, []gohelpers.Coord{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}}
	result[3] = &shape{gohelpers.Coord{4, 1}, []gohelpers.Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}}
	result[4] = &shape{gohelpers.Coord{2, 2}, []gohelpers.Coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}}}
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

var (
	imgCounter int    = 0
	draw       bool   = false
	tmpFolder  string = "/tmp"
)

func drawCavern(cavern *[][7]bool, s *shape, sPos gohelpers.Coord) {
	// warning: Switched x/y here!
	colors := make(map[gohelpers.Coord]color.RGBA)
	startX := 0
	if len(*cavern) > 100 {
		startX = len(*cavern) - 100
	}
	for x := startX; x < len(*cavern); x++ {
		xImg := x - startX
		for y := 0; y < 7; y++ {
			if (*cavern)[x][y] {
				colors[gohelpers.Coord{y + 1, xImg}] = color.RGBA{69, 69, 69, 255}
			}
		}
	}
	for x := 0; x < 100; x++ {
		colors[gohelpers.Coord{0, x}] = color.RGBA{255, 255, 255, 255}
		colors[gohelpers.Coord{8, x}] = color.RGBA{255, 255, 255, 255}
	}
	for _, p := range s.points {
		colors[gohelpers.Coord{sPos.Y + p.Y + 1, sPos.X + p.X - startX}] = color.RGBA{255, 194, 25, 255}
	}

	image.DrawImage(colors, gohelpers.Coord{0, 0}, gohelpers.Coord{8, 106}, 4, fmt.Sprintf("%s/img%05d.png", tmpFolder, imgCounter))
	imgCounter++
}

func dropShape(cavern *[][7]bool, s *shape, maxHeight int, jetPattern string, patternIdx *int) int {
	sPos := gohelpers.Coord{maxHeight + 4, 2}
	for {

		isBlocked := func(pos gohelpers.Coord) bool {
			for _, v := range s.points {
				p := gohelpers.Coord{pos.X + v.X, pos.Y + v.Y}
				if p.X < 0 {
					return true
				}
				if p.Y < 0 || p.Y > 6 {
					return true
				}
				if p.X > maxHeight {
					continue
				}
				if (*cavern)[p.X][p.Y] {
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

		if !isBlocked(gohelpers.Coord{sPos.X, sPos.Y + dir}) {
			sPos.Y += dir
		}
		// drop

		if !isBlocked(gohelpers.Coord{sPos.X - 1, sPos.Y}) {
			sPos.X--
		} else {
			if draw {
				drawCavern(cavern, s, sPos)
			}
			currentMax := maxHeight
			for newHeight := currentMax + 1; newHeight <= sPos.X+s.size.X-1; newHeight++ {
				maxHeight = newHeight
				*cavern = append(*cavern, [7]bool{})
			}
			for _, p := range s.points {
				(*cavern)[sPos.X+p.X][sPos.Y+p.Y] = true
			}
			return maxHeight
		}
		if draw {
			drawCavern(cavern, s, sPos)
		}
	}
}

func puzzle1(input []string) int {
	jetPattern := input[0]
	patternIdx := 0
	maxHeight := 0

	cavern := make([][7]bool, 1)
	cavern[0] = [7]bool{true, true, true, true, true, true, true}

	for i := 0; i < 200; i++ {
		s := shapes[i%len(shapes)]

		maxHeight = dropShape(&cavern, s, maxHeight, jetPattern, &patternIdx)
	}
	return 1
}

type cacheEntry struct {
	coords    []gohelpers.Coord
	jetIdx    int
	maxHeight int
	i         int
}

func makeCacheEntry(cavern [][7]bool, jet int, maxHeight int, i int, cacheLines int) *cacheEntry {
	result := &cacheEntry{make([]gohelpers.Coord, 0, cacheLines*7), jet, maxHeight, i}
	for x := len(cavern) - cacheLines; x < len(cavern); x++ {
		for y := 0; y < 7; y++ {
			if cavern[x][y] {
				result.coords = append(result.coords, gohelpers.Coord{x - (len(cavern) - cacheLines), y})
			}
		}
	}
	return result
}

func cmp(a []gohelpers.Coord, b []gohelpers.Coord) bool {
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
