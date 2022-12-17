package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func DrawImage(colors map[gohelpers.Coord]color.RGBA, min, max gohelpers.Coord, scale int, filename string) {
	width := max.X - min.X + 1
	widthScaled := width * scale
	height := max.Y - min.Y + 1
	heightScaled := height * scale
	myImage := image.NewRGBA(image.Rect(0, 0, widthScaled, heightScaled))
	for key, value := range colors {
		x := key.X - min.X
		y := height - 1 - (key.Y - min.Y)
		for xs := 0; xs < scale; xs++ {
			for ys := 0; ys < scale; ys++ {
				myImage.Set(x*scale+xs, y*scale+ys, value)
			}
		}
	}

	outputFile, _ := os.Create(filename)
	defer outputFile.Close()
	png.Encode(outputFile, myImage)
}
