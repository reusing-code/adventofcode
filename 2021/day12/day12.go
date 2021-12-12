package main

import (
	"strings"
)

func puzzle1(input []string) int {
	neighbors := make(map[string][]string)
	for _, str := range input {
		split := strings.Split(str, "-")
		if _, ok := neighbors[split[0]]; !ok {
			neighbors[split[0]] = make([]string, 0)
		}
		if _, ok := neighbors[split[1]]; !ok {
			neighbors[split[1]] = make([]string, 0)
		}
		neighbors[split[0]] = append(neighbors[split[0]], split[1])
		neighbors[split[1]] = append(neighbors[split[1]], split[0])
	}

	queue := make([][]string, 0)
	queue = append(queue, []string{"start"})
	count := 0
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		currentNode := elem[len(elem)-1]
		for _, neighbor := range neighbors[currentNode] {
			if neighbor == "end" {
				count++
				continue
			}
			if strings.ToLower(neighbor) == neighbor {
				skip := false
				for _, existing := range elem {
					if existing == neighbor {
						skip = true
					}
				}
				if skip {
					continue
				}
			}
			newelem := make([]string, 0, len(elem)+1)
			newelem = append(newelem, elem...)
			newelem = append(newelem, neighbor)
			queue = append(queue, newelem)
		}
	}
	return count
}

func hasDouble(in []string) bool {
	for i, str := range in {
		if strings.ToLower(str) == str {
			for j := i + 1; j < len(in); j++ {
				if in[j] == str {
					return true
				}
			}
		}
	}
	return false
}

func puzzle2(input []string) int {
	neighbors := make(map[string][]string)
	for _, str := range input {
		split := strings.Split(str, "-")
		if _, ok := neighbors[split[0]]; !ok {
			neighbors[split[0]] = make([]string, 0)
		}
		if _, ok := neighbors[split[1]]; !ok {
			neighbors[split[1]] = make([]string, 0)
		}
		neighbors[split[0]] = append(neighbors[split[0]], split[1])
		neighbors[split[1]] = append(neighbors[split[1]], split[0])
	}

	queue := make([][]string, 0)
	queue = append(queue, []string{"start"})
	count := 0
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		currentNode := elem[len(elem)-1]
		for _, neighbor := range neighbors[currentNode] {
			if neighbor == "end" {
				count++
				continue
			}
			if neighbor == "start" {
				continue
			}
			if strings.ToLower(neighbor) == neighbor {
				skip := false
				if hasDouble(elem) {
					for _, existing := range elem {
						if existing == neighbor {
							skip = true
						}
					}
				}
				if skip {
					continue
				}
			}
			newelem := make([]string, 0, len(elem)+1)
			newelem = append(newelem, elem...)
			newelem = append(newelem, neighbor)
			queue = append(queue, newelem)
		}
	}
	return count
}
