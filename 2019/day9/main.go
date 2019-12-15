package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func runBOOST(data []int64) {
	var wg sync.WaitGroup
	
	program := newProgram("BOOST", data)
	go func() {
		program.execute()
		close(program.output)
		
	}()
	go func() {
		wg.Add(1)
		for res := range program.output {
			fmt.Printf("%v\n", res)
		}
		wg.Done()
	}()
	program.input <- 2
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

	runBOOST(data)

}
