package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type coord struct {
	x int
	y int
}

func puzzle1(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	field := make(map[coord]bool, 0)
	for _, line := range split[0] {
		c := coord{}
		_, err := fmt.Sscanf(line, "%v,%v", &c.x, &c.y)
		if err != nil {
			panic(err)
		}
		field[c] = true
	}
	var foldDirection rune
	var foldAmount int
	_, err := fmt.Sscanf(split[1][0], "fold along %c=%v", &foldDirection, &foldAmount)
	if err != nil {
		panic(err)
	}

	for k, _ := range field {
		if foldDirection == 'x' {
			if k.x > foldAmount {
				newKey := coord{2*foldAmount - k.x, k.y}
				field[k] = false
				field[newKey] = true
			}
		} else {
			if k.y > foldAmount {
				newKey := coord{k.x, 2*foldAmount - k.y}
				field[k] = false
				field[newKey] = true
			}
		}
	}

	count := 0
	for _, v := range field {
		if v {
			count++
		}
	}

	return count
}

func puzzle2(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	field := make(map[coord]bool, 0)
	for _, line := range split[0] {
		c := coord{}
		_, err := fmt.Sscanf(line, "%v,%v", &c.x, &c.y)
		if err != nil {
			panic(err)
		}
		field[c] = true
	}

	for _, foldStr := range split[1] {

		var foldDirection rune
		var foldAmount int
		_, err := fmt.Sscanf(foldStr, "fold along %c=%v", &foldDirection, &foldAmount)
		if err != nil {
			panic(err)
		}

		for k, _ := range field {
			if foldDirection == 'x' {
				if k.x > foldAmount {
					newKey := coord{2*foldAmount - k.x, k.y}
					field[k] = false
					field[newKey] = true
				}
			} else {
				if k.y > foldAmount {
					newKey := coord{k.x, 2*foldAmount - k.y}
					field[k] = false
					field[newKey] = true
				}
			}
		}
	}

	var max coord

	for k, v := range field {
		if v {
			if k.x > max.x {
				max.x = k.x
			}
			if k.y > max.y {
				max.y = k.y
			}
		}
	}

	drawImage(field, max)

	return 1
}

func drawImage(pixels map[coord]bool, max coord) {
	width := max.x + 1
	height := max.y + 1
	myImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(myImage, image.Rect(0, 0, width, height), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for key, value := range pixels {
		var c color.Color
		if value {
			c = color.White
		} else {
			c = color.Black
		}
		myImage.Set(key.x, key.y, c)
	}

	outputFile, err := os.Create("/tmp/output.png")
	fmt.Println(outputFile.Name())
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, myImage)
}
