package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	model         color.Model
}

func (i Image) ColorModel() color.Model {
	return i.model
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	r := byte(x ^ y - (x * y / 2))
	g := byte((x+y)/2 - (x*y)/4)
	b := byte(x*y - (x + y/2))
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{100, 100, color.RGBAModel}
	pic.ShowImage(m)
}
