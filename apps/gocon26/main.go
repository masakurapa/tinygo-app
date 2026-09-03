package main

import (
	"image/color"
	"time"
)

const (
	modeTitle = iota
	modeQR
	modeGame
)

const (
	frame = 16
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}

	currentMode = modeTitle
)

type displayer interface {
	DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error
	FillScreen(c color.Color)
}

func main() {
	disp := initDisplay()
	lab := newLabel(320, 240)

	for {
		switch currentMode {
		case modeTitle:
			displayTitle(disp, lab)
		case modeQR:
			displayQR(disp, lab)
		case modeGame:
			displayGame(disp, lab)
		}
		time.Sleep(frame * time.Millisecond)
	}
}
