package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	w, h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	if x % 4 == 0 {
		return color.RGBA{255, 0, 255, 255}
	} else {
		return color.RGBA{0, 255, 255, 255}
	}
}


func main() {
	m := Image{50, 50}
	pic.ShowImage(m)
}
