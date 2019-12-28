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

func getSpace(data []int64) [][]rune {
	input := make(chan int64)
	output := make(chan int64, 1)
	program := intcode.NewProgram("ASCII", data, input, output)

	space := make([][]rune, 1)
	space[0] = make([]rune, 0)

	close(input)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		program.Execute()
		wg.Done()
		close(output)
	}()
	go func() {
		for val := range output {
			if val == 10 {
				space = append(space, make([]rune, 0))
			} else {
				space[len(space)-1] = append(space[len(space)-1], rune(val))
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return space
}

func runASCII(data []int64) {
	space := getSpace(data)
	alignment := 0
	for i := 1; i < len(space)-1; i++ {
		for j := 1; j < len(space[i])-1; j++ {
			h := rune(35)
			if space[i][j] == h && space[i-1][j] == h && space[i][j-1] == h && space[i+1][j] == h && space[i][j+1] == h {
				alignment += i * j
			}
		}
	}
	fmt.Printf("Alignment: %v\n", alignment)

	input := make(chan int64)
	output := make(chan int64, 1)
	data[0] = 2
	program := intcode.NewProgram("ASCII", data, input, output)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		program.Execute()
		wg.Done()
		close(output)
	}()
	go func() {
		for val := range output {
			fmt.Println(val)
		}
		wg.Done()
	}()

}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	split := strings.Split(scanner.Text(), ",")
	data := make([]int64, len(split))
	for i, v := range split {
		in, _ := strconv.ParseInt(v, 10, 64)
		data[i] = in
	}
	data[0] = 2
	runASCII(data)

}
