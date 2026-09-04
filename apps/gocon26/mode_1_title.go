package main

import (
	"time"

	"github.com/masakurapa/tinygo-app/apps/gocon26/img"
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
)

type titleData struct {
	title    string
	min, max int16
}

func init() {
	switchCurrentTitle(0)
	currentMode = modeTitle
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

	titleX -= 3
	if titleX < currentTitleData.min {
		switchCurrentTitle(currentTitle + 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func switchTitle() {
	switchCurrentTitle(0)
	titleMode = titleModeCorp
	currentMode = modeTitle
	waitRelease(PressOption1)
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
	drawMyGopher(lab)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	titleMode = titleModeMe
}

func drawMyGopher(lab *label) {
	for sy := int16(0); sy < img.MyGopherH; sy++ {
		for sx := int16(0); sx < img.MyGopherW; sx++ {
			i := (int(sy)*int(img.MyGopherW) + int(sx)) * 2
			pixel := uint16(img.MyGopherRaw[i]) | uint16(img.MyGopherRaw[i+1])<<8
			px, py := int16(60)+sx, int16(40)+sy
			if px >= 0 && px < lab.w && py >= 0 && py < lab.h {
				lab.buf[int(py)*int(lab.w)+int(px)] = pixel
			}
		}
	}
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
}
