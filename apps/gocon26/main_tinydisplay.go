//go:build !wioterminal

package main

import (
	"image/color"
	"time"

	"github.com/sago35/tinydisplay"
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var display *initdisplay.TinyDisplay

type tinyDisplay struct{}

func (*tinyDisplay) DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error {
	return display.DrawRGBBitmap(x, y, data, w, h)
}

func (*tinyDisplay) FillScreen(c color.RGBA) {
	display.FillScreen(c)
}

func initDisplay() displayer {
	d, _ := tinydisplay.NewClient("host.docker.internal", 9812, 320, 240)

	d.FillScreen(color.RGBA{0, 0, 0, 255})
	time.Sleep(100 * time.Millisecond)
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	display = &initdisplay.TinyDisplay{
		Client: d,
	}

	return &tinyDisplay{}
}

func PressKeyRight() bool { return display.GetPressedKey() == 0x106 }
func PressKeyLeft() bool  { return display.GetPressedKey() == 0x107 }
func PressKeyDown() bool  { return display.GetPressedKey() == 0x108 }
func PressKeyUp() bool    { return display.GetPressedKey() == 0x109 }
func PressEnter() bool    { return display.GetPressedKey() == 0x101 }

// Option 1（上部の左ボタン）
func PressOption1() bool { return display.GetPressedKey() == '2' }

// Option 2（上部の中央ボタン）
func PressOption2() bool { return display.GetPressedKey() == '1' }

// Option 3（上部の右）
func PressOption3() bool { return display.GetPressedKey() == '0' }

// 画面の傾きはサポート外
func CalibrateAccel()    {}
func AccelY() float64    { return 0 }
func SupportAccel() bool { return false }
