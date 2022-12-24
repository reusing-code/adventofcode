package image

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"

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

func DrawTextImage(field [][]byte, width, height int, filename string) {
	fgColor := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	bgColor := color.RGBA{R: 0x20, G: 0x20, B: 0x20, A: 0xff}

	fontBytes, err := os.ReadFile("/usr/share/fonts/truetype/freefont/FreeMono.ttf")
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	fontSize := float64(12)

	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)

	imageHeight := int(c.PointToFixed(fontSize*1.5)>>6)*height + 100
	imageWidth := int(float64(c.PointToFixed(fontSize)>>6)*float64(width)*0.6 + float64(100))
	rgba := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	textXOffset := 50
	textYOffset := 10 + int(c.PointToFixed(fontSize)>>6)

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range field {
		_, err = c.DrawString(string(s), pt)
		if err != nil {
			panic(err)
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	outputFile, _ := os.Create(filename)
	defer outputFile.Close()
	png.Encode(outputFile, rgba)
}
