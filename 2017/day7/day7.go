package day7

import (
	"adventofcode"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	weight   int
	children []*Node
	parent   *Node
}

type Mapping struct {
	parent string
	child  string
}

func createNode(name string, weight int) *Node {
	n := Node{name: name, weight: weight, parent: nil, children: make([]*Node, 0)}
	return &n
}

func findRoot(fileName string) *Node {
	nodes := make(map[string]*Node)
	mappings := make([]Mapping, 0)
	lines := adventofcode.ReadFile(fileName)
	re := regexp.MustCompile(`([[:alpha:]]+) \((.*)\)[ -> ]*(.*)`)
	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		name := res[1]
		weight, _ := strconv.Atoi(res[2])
		if len(res[3]) > 0 {
			children := strings.Split(res[3], ", ")
			for _, child := range children {
				child = strings.TrimSpace(child)
				mappings = append(mappings, Mapping{name, child})
			}
		}
		node := createNode(name, weight)
		nodes[name] = node
	}

	for _, mapping := range mappings {
		parent := nodes[mapping.parent]
		child := nodes[mapping.child]

		parent.children = append(parent.children, child)
		child.parent = parent
	}

	for _, node := range nodes {
		if node.parent == nil {
			return node
		}
	}

	return createNode("a", 0)
}


func getCorrectedWeight(fileName string) int {
	root := findRoot(fileName)

	return getCorrectedWeightSubtree(root)
}

func getCorrectedWeightSubtree(node *Node) int {
	if len(node.children) == 0 {
		return 0
	}
	weights := make(map[int]int)
	for _, child := range node.children {
		weights[getCompleteWeight(child)] += 1
	}
	if len(weights) == 1 {
		return 0
	}

	for weight, count := range weights {
		if count > 1 {
			continue
		}

	}
}

func getCompleteWeight(node *Node) int {
	weight := node.weight
	for _, child := range node.children {
		weight += getCompleteWeight(child)
	}
	return weight
}