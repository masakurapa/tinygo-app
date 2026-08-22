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
	modeStartTimer
	modeStopTimer
)

// 設定モード
const (
	settingMinutes = iota
	settingSeconds
)

var (
	black  = color.RGBA{0, 0, 0, 255}
	white  = color.RGBA{255, 255, 255, 255}
	yellow = color.RGBA{240, 200, 50, 255}
	red    = color.RGBA{200, 70, 70, 255}
	gray   = color.RGBA{100, 110, 130, 255}

	// timer setting
	minutes = 5
	seconds = 0
	// current mode
	mode    = modeSetting
	setMode = settingMinutes

	// timer mode setting
	deadlineMilliseconds int64
	halfLine             int64
	tenthLine            int64
	startTime            time.Time
)

type note struct {
	tone     float64
	duration float64
}

func main() {
	disp := initDisplay()
	lab := newLabel(320, 240)

	for {
		switch mode {
		case modeSetting:
			displaySettingMode(disp, lab)
		case modeStartTimer:
			displayStartTimerMode(disp, lab)
		case modeStopTimer:
			displayStopTimerMode(disp, lab)
		}
	}
}

func switchSettingMode() {
	mode = modeSetting
	fmt.Println("switch setting mode")
}

func switchStartTimerMode() {
	if mode != modeStopTimer {
		deadlineMilliseconds = int64(minutes*60+seconds) * 1000
		halfLine = deadlineMilliseconds / 2
		tenthLine = deadlineMilliseconds / 10
	}
	startTime = time.Now()
	mode = modeStartTimer
	fmt.Printf("switch start timer mode: %v\n", startTime)
}

func switchStopTimerMode() {
	deadlineMilliseconds = deadlineMilliseconds - time.Now().Sub(startTime).Milliseconds()
	if deadlineMilliseconds < 0 {
		deadlineMilliseconds = 0
	}
	mode = modeStopTimer
	fmt.Println("switch stop timer mode")
}

func displaySettingMode(disp *initdisplay.TinyDisplay, lab *label) {
	// switch mode
	if PressOption2() {
		switchStartTimerMode()
		return
	}

	defer func() {
		time.Sleep(100 * time.Millisecond)
	}()

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 20, "1:set / 2:start / 3:stop", black)

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

func displayStartTimerMode(disp *initdisplay.TinyDisplay, lab *label) {
	if PressOption3() {
		switchStopTimerMode()
		return
	}

	dur := time.Now().Sub(startTime)
	millisec := (deadlineMilliseconds - dur.Milliseconds())
	if millisec < 0 {
		millisec = 0
	}

	totalSeconds := millisec / 1000
	min := totalSeconds / 60
	sec := totalSeconds % 60

	if millisec <= tenthLine {
		lab.FillScreen(red)
	} else if millisec <= halfLine {
		lab.FillScreen(yellow)
	} else {
		lab.FillScreen(white)
	}

	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 20, "1:set / 2:start / 3:stop", black)
	tinyfont.WriteLine(lab, &freesans.Regular24pt7b, 90, 120, fmt.Sprintf("%02d : %02d", min, sec), black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

func displayStopTimerMode(disp *initdisplay.TinyDisplay, lab *label) {
	if PressOption1() {
		switchSettingMode()
		return
	}
	if PressOption2() {
		switchStartTimerMode()
		return
	}

	// show timer setting
	totalSeconds := deadlineMilliseconds / 1000
	min := totalSeconds / 60
	sec := totalSeconds % 60

	lab.FillScreen(gray)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 20, "1:set / 2:start / 3:stop", black)
	tinyfont.WriteLine(lab, &freesans.Regular24pt7b, 90, 120, fmt.Sprintf("%02d : %02d", min, sec), black)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 120, 200, "stopped", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
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
