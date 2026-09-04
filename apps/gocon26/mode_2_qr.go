package main

import (
	"time"

	"github.com/masakurapa/tinygo-app/apps/gocon26/img"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	qrModeCorp uint8 = iota
	qrModeMe
)

var (
	currentQRData qrData
	currentQR     int
	qrMode        uint8
	qrSwitch      bool

	qrList = []qrData{
		{
			title:  "Technical PR on X",
			bitmap: img.QrTechnicalPROnXBitmap(),
			x:      44,
			y:      26,
			scale:  7,
		},
		{
			title:  "Corporate website",
			bitmap: img.QrCorporateWebsiteBitmap(),
			x:      44,
			y:      26,
			scale:  7,
		},
		{
			title:  " Casual  interview",
			bitmap: img.QrCasualInterviewBitmap(),
			x:      48,
			y:      30,
			scale:  6,
		},
	}
)

type qrData struct {
	title       string
	bitmap      [][]bool
	x, y, scale int16
}

func displayQR(disp displayer, lab *label) {
	// switch mode
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQRMode(disp, lab)
		return
	}
	if PressOption3() {
		switchGame()
		return
	}

	// qrModeMeでの再描画はしない
	if qrMode == qrModeMe {
		return
	}

	if PressKeyUp() || PressKeyRight() {
		switchCurrentQR(currentQR + 1)
		return
	}
	if PressKeyDown() || PressKeyLeft() {
		switchCurrentQR(currentQR - 1)
		return
	}

	if !qrSwitch {
		return
	}

	lab.FillScreen(white)
	drawBitmap(lab, currentQRData.bitmap, currentQRData.x, currentQRData.y, currentQRData.scale)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 62, 24, currentQRData.title, black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
	qrSwitch = false
}

func switchQR() {
	switchCurrentQR(0)
	currentMode = modeQR
	waitRelease(PressOption2)
}

func switchQRMode(disp displayer, lab *label) {
	lab.FillScreen(white)

	defer time.Sleep(frame * 10 * time.Millisecond)

	if qrMode == qrModeMe {
		switchCurrentQR(0)
		qrMode = qrModeCorp
		return
	}

	// qrModeMeへの切り替え時は、1回だけ画像描画する
	drawBitmap(lab, img.QrMyXBitmap(), 44, 26, 7)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 76, 24, "masakurapa (X)", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	qrMode = qrModeMe
}

func switchCurrentQR(i int) {
	defer time.Sleep(frame * 10 * time.Millisecond)

	if i < 0 {
		i = len(qrList) - 1
	}
	if i >= len(qrList) {
		i = 0
	}

	currentQR = i
	currentQRData = qrList[i]
	qrSwitch = true
}

func drawBitmap(l *label, bitmap [][]bool, offsetX, offsetY, scale int16) {
	for y, row := range bitmap {
		for x, isBlack := range row {
			var c uint16
			if !isBlack {
				c = 0xFFFF // 白
			}
			for dy := int16(0); dy < scale; dy++ {
				for dx := int16(0); dx < scale; dx++ {
					px := offsetX + int16(x)*scale + dx
					py := offsetY + int16(y)*scale + dy
					if px >= 0 && px < l.w && py >= 0 && py < l.h {
						l.buf[int(px)+int(py)*int(l.w)] = c
					}
				}
			}
		}
	}
}
