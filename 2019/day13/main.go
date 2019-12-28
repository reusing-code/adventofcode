package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/reusing-code/adventofcode/2019/intcode"
)

type coord struct {
	x, y int
}

func updateMinMax(pos, min, max *coord) {
	if pos.x < min.x {
		min.x = pos.x
	}
	if pos.x > max.x {
		max.x = pos.x
	}
	if pos.y < min.y {
		min.y = pos.y
	}
	if pos.y > max.y {
		max.y = pos.y
	}
}

func drawImage(colors map[coord]int, min, max coord) {
	width := max.x - min.x + 1
	height := max.y - min.y + 1
	myImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for key, value := range colors {
		x := key.x - min.x
		y := key.y - min.y
		var c color.Color
		switch value {
		case 0:
			c = color.Black
		case 1:
			c = color.RGBA{128, 128, 128, 255}
		case 2:
			c = color.RGBA{255, 0, 0, 255}
		case 3:
			c = color.RGBA{139, 69, 19, 255}
		case 4:
			c = color.White
		}
		myImage.Set(x, y, c)
	}

	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, myImage)

}

func playArcade(data []int64) {
	input := make(chan int64)
	output := make(chan int64, 1)
	program := intcode.NewProgram("Breakout", data, input, output)

	min := coord{0, 0}
	max := coord{0, 0}
	colors := make(map[coord]int)
	score := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		program.Execute()
		wg.Done()
	}()

	go func() {
		pos := coord{0, 0}
		for i := 0; ; i++ {
			pos.x = int(<-output)
			pos.y = int(<-output)
			newColor := int(<-output)
			if pos.x == -1 && pos.y == 0 {
				score = newColor
				continue
			}

			colors[pos] = newColor
			updateMinMax(&pos, &min, &max)
		}

	}()

	wg.Wait()
	time.Sleep(time.Second)
	drawImage(colors, min, max)
	count := 0
	for _, val := range colors {
		if val == 2 {
			count++
		}
	}
	fmt.Printf("Blocks: %v, Score: %v\n", count, score)
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
	playArcade(data)
}
