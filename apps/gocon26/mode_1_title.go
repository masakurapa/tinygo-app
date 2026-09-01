package main

import (
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	// x = 65 で完全に "k" が消える
	titleMaxX int16 = 65
	// x = -160 で完全に "i" が消える
	titleMinX int16 = -160
)

var (
	titleX int16 = titleMaxX
)

func displayTitle(disp displayer, lab *label) {
	scaled := lab.Scale(5)
	scaled.FillScreen(white)

	tinyfont.WriteLine(scaled, &freesans.Regular24pt7b, titleX, 40, "kaonavi", black)
	disp.DrawRGBBitmap(0, 0, scaled.buf, scaled.w, scaled.h)

	time.Sleep(16 * time.Millisecond)

	titleX--
	if titleX < titleMinX {
		titleX = titleMaxX
	}
}
