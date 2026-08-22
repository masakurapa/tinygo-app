package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/sago35/tinydisplay/examples/initdisplay"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

// 実行モード
const (
	// 時間の設定
	modeSetting = iota
	// タイマー
	modeTimer
)

// 設定モード
const (
	settingMinutes = iota
	settingSeconds
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}

	// timer setting
	minutes = 5
	seconds = 0
	// current mode
	mode    = modeSetting
	setMode = settingMinutes
)

func main() {
	disp := initDisplay()
	lab := newLabel(320, 240)

	for {
		lab.FillScreen(white)

		switch mode {
		case modeSetting:
			switchSettingMode(disp, lab)
		case modeTimer:
		}
	}
}

func switchSettingMode(disp *initdisplay.TinyDisplay, lab *label) {
	defer func() {
		time.Sleep(100 * time.Millisecond)
	}()

	// show timer setting
	tinyfont.WriteLine(lab, &freesans.Regular24pt7b, 90, 120, fmt.Sprintf("%02d : %02d", minutes, seconds), black)
	switch setMode {
	case settingMinutes:
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 90, 140, "---------", black)
	case settingSeconds:
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 178, 140, "---------", black)
	}

	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)

	// switch setMode
	if PressKeyLeft() {
		if setMode == settingSeconds {
			setMode = settingMinutes
		}
		return
	}
	if PressKeyRight() {
		if setMode == settingMinutes {
			setMode = settingSeconds
		}
		return
	}
	// set timer
	if PressKeyUp() {
		switch setMode {
		case settingMinutes:
			setAdjustedTimerSetting(minutes+1, seconds)
		case settingSeconds:
			setAdjustedTimerSetting(minutes, seconds+1)
		}
		return
	}
	if PressKeyDown() {
		switch setMode {
		case settingMinutes:
			setAdjustedTimerSetting(minutes-1, seconds)
		case settingSeconds:
			setAdjustedTimerSetting(minutes, seconds-1)
		}
		return
	}
}

func setAdjustedTimerSetting(newMin, newSec int) {
	// マイナス補正
	newMin = max(newMin, 0)
	newSec = max(newSec, 0)

	// 秒の最大値補正
	if newSec >= 60 {
		newMin++
		newSec -= 60
	}

	// 分の最大値補正
	newMin = min(newMin, 60)
	// 分の最大値の場合、秒のカウントアップをさせない
	if newMin == 60 {
		newSec = 0
	}

	minutes = newMin
	seconds = newSec
}

type label struct {
	buf  []uint16
	w, h int16
}

func newLabel(w, h int16) *label {
	return &label{
		buf: make([]uint16, int(w)*int(h)),
		w:   w,
		h:   h,
	}
}

func (l *label) Display() error                    { return nil }
func (l *label) Size() (int16, int16)              { return l.w, l.h }
func (l *label) SetPixel(x, y int16, c color.RGBA) { l.buf[int(x)+int(y)*int(l.w)] = rgbaTo565(c) }
func (l *label) FillScreen(c color.RGBA) {
	for i := range l.buf {
		l.buf[i] = rgbaTo565(c)
	}
}

func rgbaTo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) + ((g & 0xFC00) >> 5) + ((b & 0xF800) >> 11))
}
