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

func runBOOST(data []int64, mode int) {
	var wg sync.WaitGroup
	input := make(chan int64, 1)
	output := make(chan int64, 1)

	program := intcode.NewProgram("BOOST", data, input, output)
	go func() {
		program.Execute()
		close(output)

	}()
	wg.Add(1)
	go func() {
		for res := range output {
			fmt.Printf("%v\n", res)
		}
		wg.Done()
	}()
	input <- int64(mode)
	wg.Wait()
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
	runBOOST(data, 1)
	runBOOST(data, 2)

}
