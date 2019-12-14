package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const trillion = int64(1000000000000)

type reagent struct {
	name     string
	quantity int64
}

type reaction struct {
	out reagent
	in  []reagent
}

type nanofactory struct {
	want      map[string]int64
	have      map[string]int64
	reactions map[string]reaction
}

func newReagent(in string) *reagent {
	result := &reagent{}
	elems := strings.Split(strings.TrimSpace(in), " ")
	result.name = elems[1]
	quantiy, _ := strconv.Atoi(elems[0])
	result.quantity = int64(quantiy)
	return result
}

func newNanofactory(in []string) *nanofactory {
	result := &nanofactory{}
	result.want = make(map[string]int64)
	result.have = make(map[string]int64)
	result.reactions = make(map[string]reaction)
	result.want["FUEL"] = 1
	result.have["ORE"] = math.MaxInt64

	for _, str := range in {
		react := reaction{}
		react.in = make([]reagent, 0)
		split := strings.Split(strings.TrimSpace(str), "=>")
		react.out = *newReagent(split[1])
		split2 := strings.Split(split[0], ",")
		for _, inReag := range split2 {
			react.in = append(react.in, *newReagent(inReag))
		}
		result.reactions[react.out.name] = react
	}
	return result
}

func (nf *nanofactory) iteration() bool {
	for wantKey, wantVal := range nf.want {
		enough := false
		if haveVal, ok := nf.have[wantKey]; ok {
			if haveVal >= wantVal {
				enough = true
			} else {
				wantVal -= haveVal
			}
		}
		if enough {
			continue
		}

		react := nf.reactions[wantKey]
		reactionTimes := int64(math.Ceil(float64(wantVal) / float64(react.out.quantity)))
		nf.have[wantKey] += reactionTimes * react.out.quantity
		for _, inReag := range react.in {
			nf.want[inReag.name] += reactionTimes * inReag.quantity
		}
		return false
	}
	return true
}

func calcOreRequirement(reactions []string) int64 {
	nf := newNanofactory(reactions)
	for !nf.iteration() {
	}
	return nf.want["ORE"]
}

func trillionOre(input []string) int64 {
	orePerFuel := calcOreRequirement(input)
	nf := newNanofactory(input)

	fuel := int64(trillion / orePerFuel)
	nf.want["FUEL"] = fuel
	for !nf.iteration() {
	}
	newOrePerFule := (nf.want["ORE"] / fuel) + 1

	for i := int64(trillion / newOrePerFule); ; i++ {
		nf.want["FUEL"] = i
		for !nf.iteration() {
		}
		if nf.want["ORE"] > trillion {
			return i - 1
		}
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Printf("Result %v\n", calcOreRequirement(input))

	fmt.Printf("Result 2 %v\n", trillionOre(input))

}
