package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/reusing-code/adventofcode/2019/intcode"
)

func nextPerm(p []int64) {
	for i := len(p) - 1; i >= 0; i-- {
		for i == 0 || p[i] < int64(len(p)-i-1) {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int64) []int64 {
	result := append([]int64{}, orig...)
	for i, v := range p {
		result[i], result[int64(i)+v] = result[int64(i)+v], result[i]
	}
	return result
}

func runThrusterSequence(data, p []int64, input int64) int64 {
	chans := make([]chan int64, len(p))
	for i, _ := range chans {
		chans[i] = make(chan int64, 1)
		chans[i] <- p[i]
	}
	amps := make([]intcode.Program, len(p))
	for i, _ := range amps {

		amps[i] = *intcode.NewProgram(strconv.Itoa(i), data, chans[i], chans[(i+1)%len(p)])
	}

	var wg sync.WaitGroup
	for i, _ := range amps {
		wg.Add(1)
		ii := i // do not use loop iterator variabls in goroutines!
		go func() {
			amps[ii].Execute()
			wg.Done()
		}()
	}
	chans[0] <- input
	wg.Wait()

	result := <-chans[0]

	return result
}

func thrusterSequence(data []int64) int64 {
	maxThrust := int64(0)
	orig := []int64{5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	resultChan := make(chan int64, 10)
	for p := make([]int64, 5); p[0] < int64(len(p)); nextPerm(p) {
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
		data := make([]int64, len(split))
		for i, v := range split {
			in, _ := strconv.ParseInt(v, 10, 64)
			data[i] = in
		}

		result := thrusterSequence(data)
		fmt.Printf("Result %v\n", result)
	}
}
