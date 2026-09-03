package main

import (
	"bytes"
	_ "embed"
	"time"

	"tinygo.org/x/drivers/image/png"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	titleModeCorp uint8 = iota
	titleModeMe
)

var (
	titleX           int16
	currentTitle     int
	currentTitleData titleData

	titleMode uint8

	titleList = []titleData{
		{title: "kaonavi", min: -170, max: 65},
		{title: "Face you, Face next.", min: -440, max: 65},
	}

	//go:embed image.png
	myGopher []byte
	buf      [3 * 9 * 9 * 4]uint16
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
	if PressOption1() {
		switchTitleMode(disp, lab)
		return
	}
	if PressOption2() {
		switchQR()
		return
	}
	if PressOption3() {
		switchGame()
		return
	}

	// titleModeMeでの再描画はしない
	if titleMode == titleModeMe {
		return
	}

	scaled := lab.Scale(5)
	scaled.FillScreen(white)
	tinyfont.WriteLine(scaled, &freesans.Regular24pt7b, titleX, 40, currentTitleData.title, black)
	disp.DrawRGBBitmap(0, 0, scaled.buf, scaled.w, scaled.h)

	titleX -= 2
	if titleX < currentTitleData.min {
		switchCurrentTitle(currentTitle + 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func switchTitle() {
	switchCurrentTitle(0)
	titleMode = titleModeCorp
	currentMode = modeTitle
}

func switchTitleMode(disp displayer, lab *label) {
	lab.FillScreen(white)
	disp.FillScreen(white)

	defer time.Sleep(frame * 10 * time.Millisecond)

	if titleMode == titleModeMe {
		switchCurrentTitle(0)
		titleMode = titleModeCorp
		return
	}

	// titleModeMeへの切り替え時は、1回だけ画像描画する
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 90, 25, "masakurapa", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	img := bytes.NewReader(myGopher)
	png.SetCallback(buf[:], func(data []uint16, x, y, w, h, width, height int16) {
		disp.DrawRGBBitmap(x+60, y+40, data[:w*h], w, h)
	})
	png.Decode(img)

	titleMode = titleModeMe
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
