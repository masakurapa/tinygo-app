package main

import (
	_ "embed"
	"fmt"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	//go:embed gopher.png
	gopher []byte

	gd gameData
)

var buffer [3 * 8 * 8 * 4]uint16

func displayGame(disp displayer, lab *label) {
	// switch mode
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQR()
		return
	}

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)

	lab.DrawBitmap(gopherBitmap, gopherW, gopherH, 140, 40)

	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	gd.countUp()

	time.Sleep(16 * time.Millisecond)
}

func switchGame() {
	gd = gameData{
		score: 0,
	}
	currentMode = modeGame
}

type gameData struct {
	score    int
	interval int8
}

func (gd *gameData) countUp() {
	// 10フレームに1回でステータス加算するイメージ
	if gd.interval > 10 {
		gd.score += 100
		gd.interval = 0
	}
	gd.interval++
}
