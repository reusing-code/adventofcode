package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

func comp(i *pair, j *pair) bool {
	mi := manhatten(i)
	mj := manhatten(j)
	if mi == mj {
		return i.x < j.x
	}
	return mi < mj
}

func manhatten(x *pair) int {
	return abs(x.x) + abs(x.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func getClosestIntersection(input []string) int {
	if len(input) != 2 {
		panic("Input has wrong dimension!")
	}
	res1 := getVisitedFields(input[0])
	res2 := getVisitedFields(input[1])

	i, j := 0, 0
	for {
		if res1[i] == res2[j] {
			return manhatten(&res1[i])
		}
		if comp(&res1[i], &res2[j]) {
			i++
		} else {
			j++
		}
	}

}

func getVisitedFields(input string) []pair {
	result := make([]pair, 0)
	steps := strings.Split(input, ",")
	current := pair{0, 0}
	for _, val := range steps {
		stepStr := string(val[1:])
		steps, _ := strconv.Atoi(stepStr)
		var mod func(*pair)
		switch val[0] {
		case 'R':
			mod = func(a *pair) {
				a.x += 1
			}
		case 'L':
			mod = func(a *pair) {
				a.x -= 1
			}
		case 'U':
			mod = func(a *pair) {
				a.y += 1
			}
		case 'D':
			mod = func(a *pair) {
				a.y -= 1
			}
		}
		for i := 0; i < steps; i++ {
			mod(&current)
			result = append(result, current)
		}
	}
	sort.Slice(result, func(i int, j int) bool {
		return comp(&result[i], &result[j])
	})
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

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	distance := getClosestIntersection(input)
	fmt.Printf("Result %v\n", distance)

}
