package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

const WIDTH = 25
const HEIGHT = 6

type layer struct {
	data  []int
	count []int
}

func parseImage(data []int) []layer {
	pixelPerLayer := WIDTH * HEIGHT
	layerCount := len(data) / (pixelPerLayer)
	layers := make([]layer, layerCount)

	for i := 0; i < layerCount; i++ {
		layers[i] = layer{data: data[i*pixelPerLayer : (i+1)*pixelPerLayer], count: make([]int, 10)}
		for _, v := range layers[i].data {
			layers[i].count[v]++
		}
	}
	return layers
}

func extractPassword(layers []layer) int {
	lowest0 := math.MaxInt32
	result := 0
	for _, layer := range layers {
		if layer.count[0] < lowest0 {
			lowest0 = layer.count[0]
			result = layer.count[1] * layer.count[2]
		}
	}
	return result

}

func calcPassword(data []int) int {
	layers := parseImage(data)
	return extractPassword(layers)
}

func writeImage(data []int) {
	layers := parseImage(data)
	pixelPerLayer := WIDTH * HEIGHT
	imageData := make([]int, pixelPerLayer)
	for i, _ := range imageData {
		for _, layer := range layers {
			if layer.data[i] != 2 {
				imageData[i] = layer.data[i]
				break
			}
		}
	}

	myImage := image.NewGray(image.Rect(0, 0, WIDTH, HEIGHT))
	for i, _ := range imageData {
		for j := 0; j < WIDTH; j++ {
			myImage.Pix[i] = uint8(imageData[i] * 255)
		}
	}

	outputFile, _ := os.Create("output.png")
	defer outputFile.Close()
	png.Encode(outputFile, myImage)
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
	split := strings.Split(scanner.Text(), "")
	data := make([]int, len(split))
	for i, v := range split {
		in, _ := strconv.Atoi(v)
		data[i] = in
	}

	result := calcPassword(data)
	fmt.Printf("Result %v\n", result)
	writeImage(data)
}
