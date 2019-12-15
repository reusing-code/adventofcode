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
	fmt.Printf("min: %v max: %v width %v height: %v\n", min, max, width, height)
	myImage := image.NewGray(image.Rect(0, 0, width, height))
	for key, value := range colors {
		x := key.x - min.x
		y := height - 1 - (key.y - min.y)
		c := color.Black
		if value == 1 {
			c = color.White
		}
		myImage.Set(x, y, c)
	}

	outputFile, _ := os.Create("output.png")
	defer outputFile.Close()
	png.Encode(outputFile, myImage)

}

func runPaintRobot(data []int64) {
	input := make(chan int64, 2)
	output := make(chan int64, 1)
	program := intcode.NewProgram("Painter", data, input, output)

	min := coord{0, 0}
	max := coord{0, 0}
	colors := make(map[coord]int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		program.Execute()
		wg.Done()
	}()

	go func() {
		pos := coord{0, 0}
		direction := 0
		colors[pos] = 1
		for {
			oldColor := 0
			if val, ok := colors[pos]; ok {
				oldColor = val
			}
			input <- int64(oldColor)
			newColor := int(<-output)
			turn := <-output
			if newColor != oldColor {
				colors[pos] = newColor
				updateMinMax(&pos, &min, &max)
			}
			if turn == 0 {
				direction = (direction + 3) % 4
			} else {
				direction = (direction + 1) % 4
			}
			switch direction {
			case 0:
				pos.y++
			case 1:
				pos.x++
			case 2:
				pos.y--
			case 3:
				pos.x--
			}

		}

	}()

	wg.Wait()

	drawImage(colors, min, max)

	fmt.Printf("Panels painted atleast once: %v\n", len(colors))
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

	runPaintRobot(data)
}
