package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	coords []int
}

type system struct {
	moons      []vec
	velocities []vec
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int64) int64 {
	result := a * b / gcd(a, b)
	return result
}

func (s system) calcEnergy() int {
	totalEnergy := 0
	for i := range s.moons {
		pot := 0
		for _, p := range s.moons[i].coords {
			pot += abs(p)
		}
		kin := 0
		for _, p := range s.velocities[i].coords {
			kin += abs(p)
		}
		totalEnergy += pot * kin
	}
	return totalEnergy
}

func (s system) calcStep() {
	for i := range s.moons {
		for j := i + 1; j < len(s.moons); j++ {
			for k := 0; k < 3; k++ {
				if s.moons[i].coords[k] > s.moons[j].coords[k] {
					s.velocities[i].coords[k]--
					s.velocities[j].coords[k]++
				} else if s.moons[i].coords[k] < s.moons[j].coords[k] {
					s.velocities[i].coords[k]++
					s.velocities[j].coords[k]--
				}
			}
		}
	}

	for i := range s.moons {
		for k := 0; k < 3; k++ {
			s.moons[i].coords[k] += s.velocities[i].coords[k]
		}
	}
}

func createSystem(input []string) system {
	sys := system{make([]vec, 0), make([]vec, 0)}
	for _, str := range input {
		pos := vec{coords: make([]int, 3)}
		split := strings.Split(strings.Trim(strings.TrimSpace(str), "<>"), ",")
		for i, val := range split {
			split2 := strings.Split(strings.TrimSpace(val), "=")
			number, _ := strconv.Atoi(strings.TrimSpace(split2[1]))
			pos.coords[i] = number
		}
		sys.moons = append(sys.moons, pos)
		sys.velocities = append(sys.velocities, vec{coords: make([]int, 3)})
	}
	return sys
}

func simulateSystem(input []string, steps int) int {
	sys := createSystem(input)

	for i := 0; i < steps; i++ {
		sys.calcStep()
	}
	return sys.calcEnergy()
}

func howManySteps(in []string) int64 {
	sys := createSystem(in)
	periods := make([]int64, 3)
	func() {
		for i := int64(1); ; i++ {
			sys.calcStep()
			for k := range periods {
				periodFound := true
				for m := range sys.moons {
					if sys.velocities[m].coords[k] != 0 {
						periodFound = false
					}
				}
				if periodFound && periods[k] == 0 {
					periods[k] = i
					if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 {
						return
					}
				}
			}
		}
	}()
	result := lcm(periods[0], periods[1])
	result = lcm(result, periods[2])
	return result * 2
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
	fmt.Printf("Result %v\n", simulateSystem(input, 1000))
	fmt.Printf("Result2 %v\n", howManySteps(input))

}
