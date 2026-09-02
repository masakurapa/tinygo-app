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

	switch {
	case gd.waiting():
		displayGameWaiting(disp, lab)
	case gd.playing():
		displayGamePlaying(disp, lab)
	case gd.finished():
		displayGameFinished(disp, lab)
	}

	time.Sleep(16 * time.Millisecond)
}

func switchGame() {
	gd = gameData{
		mode: gameModeFinished,
	}
	currentMode = modeGame
}

func displayGameWaiting(disp displayer, lab *label) {
	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	lab.DrawBitmap(gopherBitmap, gopherW, gopherH, 140, 40)
	tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 45, 120, "Press to start!!", black)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 45, 200, "[<-] Move left / Move right [->]", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

func displayGamePlaying(disp displayer, lab *label) {
	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	lab.DrawBitmap(gopherBitmap, gopherW, gopherH, 140, 40)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	gd.countUp()
}

func displayGameFinished(disp displayer, lab *label) {
	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	lab.DrawBitmap(gopherBitmap, gopherW, gopherH, 140, 40)
	tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 65, 120, "Game Over!!", black)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 200, "Press the [top-right button] to go back", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
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
