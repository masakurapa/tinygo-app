package main

import (
	"image/color"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

type displayer interface {
	DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error
}

func main() {
	disp := initDisplay()
	lab := newLabel(320, 240)

	for {
		displayTitle(disp, lab)
	}
}
