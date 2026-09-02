package main

import (
	"github.com/skip2/go-qrcode"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	currentQRData qrData
	currentQR     int
	qrSwitch      bool

	qrList = []qrData{
		qrData{title: "Technical PR on X", url: "https://x.com/kaonavi_devs", x: 44, y: 26, scale: 7},
		qrData{title: "Corporate website", url: "https://corp.kaonavi.jp/", x: 44, y: 26, scale: 7},
		qrData{title: "Casual interview", url: "https://recruit.kaonavi.jp/recruit-info", x: 48, y: 30, scale: 6},
	}
)

type qrData struct {
	title       string
	url         string
	x, y, scale int16
}

func displayQR(disp displayer, lab *label) {
	// switch mode
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption3() {
		switchGame()
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

	q, _ := qrcode.New(currentQRData.url, qrcode.Low)

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 62, 20, currentQRData.title, black)
	drawBitmap(lab, q.Bitmap(), currentQRData.x, currentQRData.y, currentQRData.scale)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
	qrSwitch = false
}

func switchQR() {
	switchCurrentQR(0)
	currentMode = modeQR
}

func switchCurrentQR(i int) {
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
