package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type planet struct {
	name   string
	parent *planet
	orbits int
}

type orbit struct {
	planets map[string]*planet
}

func (p *planet) calcOrbits() int {
	if p.orbits < 0 {
		if p.parent == nil {
			p.orbits = 0
		} else {
			p.orbits = p.parent.calcOrbits() + 1
		}
	}
	return p.orbits
}

func (o *orbit) get(key string) *planet {
	if _, ok := o.planets[key]; !ok {
		newP := planet{name: key, parent: nil, orbits: -1}
		o.planets[key] = &newP
	}

	return o.planets[key]
}

func (o *orbit) add(in string) {
	splitted := strings.Split(in, ")")
	parent := o.get(splitted[0])
	child := o.get(splitted[1])
	child.parent = parent
}

func (o *orbit) calcOrbits() int {
	result := 0
	for _, val := range o.planets {
		result += val.calcOrbits()
	}
	return result
}

func (o *orbit) calcSanta() int {
	result := 0
	yp := o.get("YOU").parent
	sp := o.get("SAN").parent
	for ; yp.name != sp.name; result++ {
		if yp.calcOrbits() > sp.calcOrbits() {
			yp = yp.parent
		} else {
			sp = sp.parent
		}
	}
	return result
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	o := orbit{planets: make(map[string]*planet)}
	for scanner.Scan() {
		o.add(scanner.Text())
	}
	fmt.Printf("Result: %v\n", o.calcOrbits())

	fmt.Printf("Result2: %v\n", o.calcSanta())
}
