package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m Image) At(x, y int) color.Color {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return color.RGBA{} // Return transparent for out of bounds
	}
	// Create a color based on the pixel position
	r := uint8(x % 256)
	g := uint8(y % 256)
	b := uint8((x + y) % 256)
	return color.RGBA{R: r, G: g, B: b, A: 255} // Fully opaque
}

func main() {
	m := Image{1024, 1024}
	pic.ShowImage(m)
}
