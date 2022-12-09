package main

import (
	"fmt"
	"sort"
)

type nodeType int

const (
	nodeType_dir = iota
	nodeType_file
)

type fsNode interface {
	size() int
	children() map[string]fsNode
	parent() fsNode
	nodeType() nodeType
}

type dir struct {
	mSize     int
	mChildren map[string]fsNode
	mParent   fsNode
}

func (d *dir) size() int {
	if d.mSize >= 0 {
		return d.mSize
	}
	d.mSize = 0
	for _, child := range d.mChildren {
		d.mSize += child.size()
	}
	return d.mSize
}

func (d *dir) children() map[string]fsNode {
	return d.mChildren
}

func (d *dir) parent() fsNode {
	return d.mParent
}

func (d *dir) nodeType() nodeType {
	return nodeType_dir
}

type file struct {
	mSize   int
	mParent fsNode
}

func (f *file) size() int {
	return f.mSize
}

func (f *file) children() map[string]fsNode {
	return make(map[string]fsNode)
}

func (f *file) parent() fsNode {
	return f.mParent
}

func (f *file) nodeType() nodeType {
	return nodeType_file
}

func visit(root fsNode, visitor func(fsNode)) {
	visitor(root)
	for _, child := range root.children() {
		visit(child, visitor)
	}
}

func parseFS(input []string) fsNode {
	root := &dir{-1, make(map[string]fsNode), nil}
	var currentNode fsNode = nil
	inLs := false
	for _, in := range input {
		if in[0] == '$' {
			inLs = false
			// command
			if in[2:4] == "cd" {
				if in[5] == '/' {
					currentNode = root
				} else if len(in) >= 7 && in[5:7] == ".." {
					if currentNode.parent() == nil {
						panic("invalid input")
					}
					currentNode = currentNode.parent()
				} else {
					dirName := in[5:]
					if _, prs := currentNode.children()[dirName]; !prs || currentNode.children()[dirName].nodeType() != nodeType_dir {
						panic("invalid input")
					}
					currentNode = currentNode.children()[dirName]
				}
			} else if in[2:4] == "ls" {
				if currentNode.nodeType() != nodeType_dir {
					panic("invalid input")
				}
				inLs = true
			} else {
				panic("invalid input")
			}
		} else {
			// output
			if in[:3] == "dir" {
				if !inLs {
					panic("invalid input")
				}
				dirname := ""
				fmt.Sscanf(in, "dir %s", &dirname)
				if _, prs := currentNode.children()[dirname]; !prs {
					currentNode.children()[dirname] = &dir{-1, make(map[string]fsNode), currentNode}
				}
			} else {
				size := 0
				filename := ""
				fmt.Sscanf(in, "%d %s", &size, &filename)
				if _, prs := currentNode.children()[filename]; !prs {
					currentNode.children()[filename] = &file{size, currentNode}
				}
			}
		}
	}
	return root
}

func puzzle1(input []string) int {
	root := parseFS(input)
	const dirMaxSize = 100000
	sum := 0
	visit(root, func(node fsNode) {
		if node.nodeType() == nodeType_dir && node.size() <= dirMaxSize {
			sum += node.size()
		}
	})
	return sum
}

type dirSize struct {
	d fsNode
	s int
}

func puzzle2(input []string) int {
	root := parseFS(input)
	required := 30000000 - (70000000 - root.size())
	dirs := make([]dirSize, 0)
	visit(root, func(node fsNode) {
		if node.nodeType() == nodeType_dir {
			dirs = append(dirs, dirSize{node, node.size()})
		}
	})
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].s < dirs[j].s
	})
	for _, dir := range dirs {
		if dir.s >= required {
			return dir.s
		}
	}

	return 0
}
