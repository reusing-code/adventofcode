package main

import "fmt"

type mineral int

const (
	ore mineral = iota
	clay
	obsidian
	geode
)

type cost [4]int

type bluePrint struct {
	number int
	costs  [4]cost
}

type (
	stockPile [4]int
	robots    [4]int
)

type state struct {
	min int
	sp  stockPile
	r   robots
}

func (sp stockPile) add(r robots) stockPile {
	for i := range sp {
		sp[i] += r[i]
	}
	return sp
}

func parseBlueprint(in string) bluePrint {
	bp := bluePrint{}
	fmt.Sscanf(in, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&bp.number, &bp.costs[ore][ore], &bp.costs[clay][ore], &bp.costs[obsidian][ore], &bp.costs[obsidian][clay], &bp.costs[geode][ore], &bp.costs[geode][obsidian])
	return bp
}

func copy(s state) state {
	return state{s.min, stockPile{s.sp[0], s.sp[1], s.sp[2], s.sp[3]}, robots{s.r[0], s.r[1], s.r[2], s.r[3]}}
}

func affordable(sp stockPile, robotCost cost) bool {
	for i := 0; i < 4; i++ {
		if sp[i] < robotCost[i] {
			return false
		}
	}
	return true
}

func subtract(sp *stockPile, robotCost cost) {
	for i := 0; i < 4; i++ {
		(*sp)[i] -= robotCost[i]
	}
}

func minDaysObsidian(required int, robots int, sp int) int {
	for i := 0; i < 24; i++ {
		if required <= sp+i*robots+((i-1)*i)/2 {
			return i
		}
	}
	return -1
}

func (bp bluePrint) simulate(mins int) int {
	start := state{0, stockPile{0, 0, 0, 0}, robots{1, 0, 0, 0}}
	queue := []state{start}
	maxGeode := 0
	var maxCost cost
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if bp.costs[i][j] > maxCost[j] {
				maxCost[j] = bp.costs[i][j]
			}
		}
	}
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		daysLeft := mins - current.min
		if daysLeft == 0 {
			if current.sp[geode] > maxGeode {
				maxGeode = current.sp[geode]
			}
			continue
		}
		// maximum geodes achieveable: current geodes + output of current robots + 1 new geode robot per min
		if current.sp[geode]+daysLeft*current.r[geode]+(((daysLeft-1)*daysLeft)/2) <= maxGeode {
			continue
		}
		mdo := minDaysObsidian(bp.costs[geode][obsidian], current.r[obsidian], current.sp[obsidian])
		if current.r[geode] == 0 && (mdo < 0 || mdo > daysLeft) {
			continue
		}
		// which robot get's built
		af := 0
		for i := 0; i < 4; i++ {
			if daysLeft == 1 {
				// building a robot on last day doesn't do anything
				continue
			}
			c := bp.costs[i]
			if !affordable(current.sp, c) {
				continue
			}
			if i != int(geode) && affordable(current.sp, bp.costs[geode]) {
				// always build geode robot when possible
				continue
			}
			af++
			if daysLeft == 2 && i < int(geode) {
				// building anything less than a geode robot on the second to last day doesn't do anything
				continue
			}

			if daysLeft == 3 && i == int(clay) {
				// clay doesn't help anymore when only 3 day left
				continue
			}
			if i != int(geode) && current.r[i] >= maxCost[i] {
				// we never need more robots of any type than the max cost for that type
				continue
			}
			s := copy(current)
			subtract(&s.sp, c)
			s.sp = s.sp.add(s.r)
			s.r[i]++
			s.min++
			queue = append(queue, s)
		}
		// if we have resources for any robot, not building a robot is not an option
		if af == 4 {
			continue
		}
		if af >= 2 && current.r[clay] == 0 {
			// when we only have ore robots we need to build something
			continue
		}
		// no robot built
		current.sp = current.sp.add(current.r)
		current.min++
		queue = append(queue, current)

	}
	return maxGeode
}

func puzzle1(input []string) int {
	sumQualityLevel := 0
	for _, str := range input {
		bp := parseBlueprint(str)
		geodes := bp.simulate(24)
		fmt.Println(bp.number, geodes)
		sumQualityLevel += bp.number * geodes
	}
	return sumQualityLevel
}

func puzzle2(input []string) int {
	result := 1
	for _, str := range input {
		bp := parseBlueprint(str)
		if bp.number > 3 {
			return result
		}
		geodes := bp.simulate(32)
		fmt.Println(bp.number, geodes)
		result *= geodes
	}
	return result
}
