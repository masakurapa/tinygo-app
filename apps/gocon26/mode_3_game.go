package main

import (
	"fmt"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	gd gameData
)

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

	switch {
	case gd.waiting():
		tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 45, 120, "Press to start!!", black)

		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 45, 200, "[<-] Move left / Move right [->]", black)

	case gd.playing():

		gd.countUp()
	case gd.finished():
		tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 65, 120, "Game Over!!", black)
	}

	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
	time.Sleep(16 * time.Millisecond)
}

func switchGame() {
	gd = gameData{
		mode: gameModeWaiting,
	}
	currentMode = modeGame
}

type gameData struct {
	mode int8

	score    int
	interval int8
}

const (
	gameModeWaiting int8 = iota
	gameModePlaying
	gameModeFinished
)

func (gd *gameData) countUp() {
	if !gd.playing() {
		return
	}

	// 10フレームに1回でステータス加算するイメージ
	if gd.interval > 10 {
		gd.score += 100
		gd.interval = 0
	}
	gd.interval++
}

func (gd *gameData) waiting() bool {
	return gd.mode == gameModeWaiting
}

func (gd *gameData) playing() bool {
	return gd.mode == gameModePlaying
}

func (gd *gameData) finished() bool {
	return gd.mode == gameModeFinished
}
