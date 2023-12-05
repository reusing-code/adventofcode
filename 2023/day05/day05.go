package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type rang struct {
	startDestination int
	startSource      int
	length           int
}

func parseRanges(input []string) []*rang {
	result := make([]*rang, 0)
	for _, r := range input[1:] {
		rang := &rang{}
		fmt.Sscanf(r, "%d %d %d", &rang.startDestination, &rang.startSource, &rang.length)
		result = append(result, rang)
	}
	slices.SortFunc(result, func(a *rang, b *rang) int { return cmp.Compare(a.startSource, b.startSource) })
	return result
}

func doStep(in int, step []*rang) int {
	for _, r := range step {
		if in < r.startSource {
			return in
		}
		if in < r.startSource+r.length {
			return r.startDestination + in - r.startSource
		}
	}
	return in
}

func puzzle1(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	split0 := strings.Fields(split[0][0])
	seeds := make([]int, 0)
	for _, v := range split0[1:] {
		i, _ := strconv.Atoi(v)
		seeds = append(seeds, i)
	}

	steps := make([][]*rang, 0)
	for _, v := range split[1:] {
		steps = append(steps, parseRanges(v))
	}
	result := math.MaxInt
	for _, seed := range seeds {
		current := seed
		for _, step := range steps {
			current = doStep(current, step)
		}
		if current < result {
			result = current
		}
	}
	return result
}

type valueRange struct {
	from int
	to   int
}

func puzzle2(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	split0 := strings.Fields(split[0][0])
	seeds := split0[1:]
	seedRanges := make([]*valueRange, 0)
	for i := 0; i < len(seeds)/2; i++ {
		from, _ := strconv.Atoi(seeds[i*2])
		length, _ := strconv.Atoi(seeds[i*2+1])
		seedRanges = append(seedRanges, &valueRange{from, from + length - 1})
	}
	slices.SortFunc(seedRanges, func(a *valueRange, b *valueRange) int { return cmp.Compare(a.from, b.from) })
	for i := 1; i < len(seedRanges); i++ {
		if seedRanges[i-1].to >= seedRanges[i].from {
			panic("Overlapping ranges")
		}
	}

	steps := make([][]*rang, 0)
	for _, v := range split[1:] {
		steps = append(steps, parseRanges(v))
	}
	result := math.MaxInt

	for _, step := range steps {
		newRanges := make([]*valueRange, 0)
		for i := 0; i < len(seedRanges); i++ {
			finished := false
			for _, stepRange := range step {
				if seedRanges[i].from < stepRange.startSource {
					newRanges = append(newRanges, &valueRange{seedRanges[i].from, min(seedRanges[i].to, stepRange.startSource)})
					if seedRanges[i].to >= stepRange.startSource {
						seedRanges[i] = &valueRange{stepRange.startSource, seedRanges[i].to}
						i--
					}
					finished = true
					break
				}
				if seedRanges[i].from < stepRange.startSource+stepRange.length {
					diff := stepRange.startDestination - stepRange.startSource
					newRanges = append(newRanges, &valueRange{seedRanges[i].from + diff, min(seedRanges[i].to, stepRange.startSource+stepRange.length-1) + diff})
					if seedRanges[i].to >= stepRange.startSource+stepRange.length {
						seedRanges[i] = &valueRange{stepRange.startSource + stepRange.length, seedRanges[i].to}
						i--
					}
					finished = true
					break
				}
			}
			if !finished {
				newRanges = append(newRanges, seedRanges[i])
			}
		}
		seedRanges = newRanges
		slices.SortFunc(seedRanges, func(a *valueRange, b *valueRange) int { return cmp.Compare(a.from, b.from) })

	}

	for _, v := range seedRanges {
		if v.from < result {
			result = v.from
		}
	}
	return result
}
