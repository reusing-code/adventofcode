package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		for i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func runThrusterSequence(data, p []int, input int) int {
	chans := make([]chan int, len(p))
	for i, _ := range chans {
		chans[i] = make(chan int, 1)
		chans[i] <- p[i]
	}
	amps := make([]program, len(p))
	for i, _ := range amps {
		localData := append([]int{}, data...)
		amps[i] = *newProgram(strconv.Itoa(i), localData)
		amps[i].input = chans[i]
		amps[i].output = chans[(i+1)%len(p)]
	}

	var wg sync.WaitGroup
	for i, _ := range amps {
		wg.Add(1)
		ii := i // do not use loop iterator variabls in goroutines!
		go func() {
			amps[ii].execute()
			wg.Done()
		}()
	}
	chans[0] <- input
	wg.Wait()

	result := <- chans[0]

	return result
}

func thrusterSequence(data []int) int {
	maxThrust := 0
	orig := []int{5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	resultChan := make(chan int, 10)
	for p := make([]int, 5); p[0] < len(p); nextPerm(p) {
		wg.Add(1)
		thrusterConfig := getPerm(orig, p)
		go func() {
			resultChan <- runThrusterSequence(data, thrusterConfig, 0)
			wg.Done()
		}()

	}
	go func() {
		for res := range resultChan {
			if res > maxThrust {
				maxThrust = res
			}
		}
	}()
	wg.Wait()

	return maxThrust
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		data := make([]int, len(split))
		for i, v := range split {
			in, _ := strconv.Atoi(v)
			data[i] = in
		}

		result := thrusterSequence(data)
		fmt.Printf("Result %v\n", result)
	}
}
