package main

import (
	"fmt"
	"sort"
	"strings"
)

type valve struct {
	name     string
	flowRate int
}

type adjacency map[string][]string

var (
	valves       map[string]valve
	valvesSorted []valve
	adjs         adjacency
)

func parseValves(input []string) (valves map[string]valve, adjs adjacency) {
	valves = make(map[string]valve)
	adjs = make(adjacency)
	for _, in := range input {
		name := ""
		flowRate := 0
		fmt.Sscanf(in, "Valve %s has flow rate=%d;", &name, &flowRate)
		split := strings.Split(in, "to valve")
		if strings.HasPrefix(split[1], "s ") {
			split2 := strings.Split(split[1][2:], ", ")
			adjs[name] = split2
		} else {
			adjs[name] = []string{split[1][1:]}
		}
		valves[name] = valve{name, flowRate}
	}
	return
}

type pathElem struct {
	name string
	dist int
}

type pair struct {
	from string
	to   string
}

var shortestPathCache map[pair]int = make(map[pair]int)

func shortestPath(from string, to string) int {
	if from == to {
		return 0
	}
	if v, prs := shortestPathCache[pair{from, to}]; prs {
		return v
	}
	visited := make(map[string]bool)
	for k := range valves {
		visited[k] = false
	}

	queue := []pathElem{{from, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current.name] {
			continue
		}
		if current.name == to {
			shortestPathCache[pair{from, to}] = current.dist
			return current.dist
		}
		visited[current.name] = true
		for _, adj := range adjs[current.name] {
			queue = append(queue, pathElem{adj, current.dist + 1})
		}
	}
	panic(fmt.Sprintf("No path from %s to %s", from, to))
}

type queueElem struct {
	min              int
	releasedPressure int
	openValves       map[string]bool
	location         string
}

func iterative() int {
	maxResult := 0

	queue := make([]queueElem, 0)
	maxFlow := 0
	openValves := make(map[string]bool)
	for k, v := range valves {
		if v.flowRate == 0 {
			// valves with no flowRate might as well be open
			openValves[k] = true
		} else {
			openValves[k] = false
		}
		maxFlow += v.flowRate
	}

	queue = append(queue, queueElem{1, 0, openValves, "AA"})

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if ((30-current.min+1)*maxFlow)+current.releasedPressure < maxResult {
			continue
		}

		allopen := true
		sumOpen := 0
		for k, v := range current.openValves {
			if v {
				sumOpen += valves[k].flowRate
			} else {
				allopen = false
			}
		}
		newReleasedPressure := current.releasedPressure + sumOpen
		if current.min == 30 {
			if newReleasedPressure > maxResult {
				maxResult = newReleasedPressure
			}
		}
		if allopen {
			res := current.releasedPressure + (30-current.min+1)*sumOpen
			if res > maxResult {
				maxResult = res
			}
		}

		// option 1: open valve
		if !current.openValves[current.location] {
			newOpen := make(map[string]bool)
			for k, v := range current.openValves {
				newOpen[k] = v
			}
			newOpen[current.location] = true
			queue = append(queue, queueElem{current.min + 1, newReleasedPressure, newOpen, current.location})
			continue
		}

		// option 2: move
		for k, v := range current.openValves {
			if v {
				continue
			}
			n := shortestPath(current.location, valves[k].name)
			if n == 0 {
				continue
			}
			name := valves[k].name
			newOpen := make(map[string]bool)
			for k, v := range current.openValves {
				newOpen[k] = v
			}
			queue = append(queue, queueElem{current.min + n, current.releasedPressure + n*sumOpen, newOpen, name})
		}

		// option 3: do nothing
		res := current.releasedPressure + (30-current.min)*sumOpen
		if res > maxResult {
			maxResult = res
		}

	}

	return maxResult
}

func puzzle1(input []string) int {
	valves, adjs = parseValves(input)
	result := iterative()
	return result
}

type worker struct {
	finished    bool
	location    string
	destination string
	travelTime  int
}

type workerOption struct {
	w         worker
	openValve string
}

func workerOptions(w worker, openValves map[string]bool) []workerOption {
	// option 1: already decided to not do anything anymore
	if w.finished {
		return []workerOption{{w, ""}}
	}
	// option 2: worker is traveling
	if w.travelTime > 0 {
		w.travelTime--
		if w.travelTime == 0 {
			w.location = w.destination
			w.destination = ""
		}
		return []workerOption{{w, ""}}
	}
	// option 3: open valve. We always open valves if we can, since we only travel for this exact purpose
	if !openValves[w.location] {
		return []workerOption{{w, w.location}}
	}
	options := make([]workerOption, 0)

	// option 4: travel
	count := 0
	for _, v := range valvesSorted {
		if openValves[v.name] {
			continue
		}
		count++
		n := shortestPath(w.location, v.name)
		newW := w
		if n > 1 {
			newW.destination = v.name
			newW.travelTime = n - 1
		} else {
			newW.location = v.name
			newW.travelTime = 0
		}
		options = append(options, workerOption{newW, ""})
	}

	// option 5: do nothing
	if count < 2 {
		newW := w
		newW.finished = true
		options = append(options, workerOption{newW, ""})
	}

	return options
}

func workerOptionsNoTravel(w worker, openValves map[string]bool) []workerOption {
	// option 1: already decided to not do anything anymore
	if w.finished {
		return []workerOption{{w, ""}}
	}
	// option 2: worker is traveling
	if w.travelTime > 0 {
		w.travelTime--
		if w.travelTime == 0 {
			w.location = w.destination
			w.destination = ""
		}
		return []workerOption{{w, ""}}
	}
	// option 3: open valve. We always open valves if we can, since we only travel for this exact purpose
	if !openValves[w.location] {
		return []workerOption{{w, w.location}}
	}
	// option 5: do nothing
	newW := w
	newW.finished = true
	return []workerOption{{newW, ""}}
}

type queueElem2 struct {
	min              int
	releasedPressure int
	worker1          worker
	worker2          worker
	openValves       map[string]bool
}

func iterative2() int {
	maxResult := 0

	queue := make([]queueElem2, 0)
	maxFlow := 0
	startValves := make(map[string]bool)
	for k, v := range valves {
		if v.flowRate == 0 {
			// valves with no flowRate might as well be open
			startValves[k] = true
		} else {
			startValves[k] = false
		}
		maxFlow += v.flowRate
	}

	queue = append(queue, queueElem2{5, 0, worker{false, "AA", "", 0}, worker{false, "AA", "", 0}, startValves})
	iterations := 0
	for len(queue) > 0 {
		iterations++
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if ((30-current.min+1)*maxFlow)+current.releasedPressure < maxResult {
			continue
		}

		allopen := true
		sumOpen := 0
		for k, v := range current.openValves {
			if v {
				sumOpen += valves[k].flowRate
			} else {
				allopen = false
			}
		}
		newReleasedPressure := current.releasedPressure + sumOpen
		if current.min == 30 {
			if newReleasedPressure > maxResult {
				maxResult = newReleasedPressure
			}
			continue
		}
		if allopen || (current.worker1.finished && current.worker2.finished) {
			res := current.releasedPressure + (30-current.min+1)*sumOpen
			if res > maxResult {
				maxResult = res
			}
			continue
		}

		worker1Options := workerOptions(current.worker1, current.openValves)
		worker2Options := workerOptions(current.worker2, current.openValves)

		for _, w1opt := range worker1Options {
			for _, w2opt := range worker2Options {
				if w1opt.w.destination != "" && w1opt.w.destination == w2opt.w.destination {
					continue
				}
				if (w1opt.w.destination == w2opt.openValve && w2opt.openValve != "") || (w2opt.w.destination == w1opt.openValve && w1opt.openValve != "") {
					continue
				}
				newOpenValves := make(map[string]bool)
				for k, v := range current.openValves {
					newOpenValves[k] = v
				}

				if w1opt.openValve != "" {
					newOpenValves[w1opt.openValve] = true
				}
				if w2opt.openValve != "" {
					newOpenValves[w2opt.openValve] = true
				}
				queue = append(queue, queueElem2{current.min + 1, newReleasedPressure, w1opt.w, w2opt.w, newOpenValves})

			}
		}

	}

	fmt.Println(iterations)

	return maxResult
}

func puzzle2(input []string) int {
	valves, adjs = parseValves(input)
	valvesSorted = make([]valve, 0, len(valves))
	for _, v := range valves {
		valvesSorted = append(valvesSorted, v)
	}
	sort.Slice(valvesSorted, func(i, j int) bool {
		return valvesSorted[i].flowRate < valvesSorted[j].flowRate
	})
	result := iterative2()
	return result
}
