package main

import (
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	titleX           int16
	currentTitle     int
	currentTitleData titleData

	titleList = []titleData{
		{title: "kaonavi", min: -170, max: 65},
		{title: "Face you, Face next.", min: -440, max: 65},
	}
)

type titleData struct {
	title    string
	min, max int16
}

func init() {
	switchCurrentTitle(0)
}

func displayTitle(disp displayer, lab *label) {
	// switch mode
	if PressOption2() {
		switchQR()
		return
	}
	if PressOption3() {
		switchGame()
		return
	}

	scaled := lab.Scale(5)
	scaled.FillScreen(white)
	tinyfont.WriteLine(scaled, &freesans.Regular24pt7b, titleX, 40, currentTitleData.title, black)
	disp.DrawRGBBitmap(0, 0, scaled.buf, scaled.w, scaled.h)

	time.Sleep(16 * time.Millisecond)

	titleX -= 2
	if titleX < currentTitleData.min {
		switchCurrentTitle(currentTitle + 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func switchTitle() {
	switchCurrentTitle(0)
	currentMode = modeTitle
}

func switchCurrentTitle(i int) {
	if i < 0 {
		i = len(titleList) - 1
	}
	if i >= len(titleList) {
		i = 0
	}

	currentTitle = i
	currentTitleData = titleList[i]
	titleX = currentTitleData.max

	currentMode = modeTitle
}
