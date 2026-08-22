package main

import (
	"image/color"
	"time"

	"github.com/sago35/tinydisplay"
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var display *initdisplay.TinyDisplay

func initDisplay() *initdisplay.TinyDisplay {
	d, _ := tinydisplay.NewClient("host.docker.internal", 9812, 320, 240)

	d.FillScreen(color.RGBA{0, 0, 0, 255})
	time.Sleep(100 * time.Millisecond)
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	display = &initdisplay.TinyDisplay{
		Client: d,
	}

	return display
}

func PressKeyRight() bool { return display.GetPressedKey() == 0x106 }
func PressKeyLeft() bool  { return display.GetPressedKey() == 0x107 }
func PressKeyDown() bool  { return display.GetPressedKey() == 0x108 }
func PressKeyUp() bool    { return display.GetPressedKey() == 0x109 }

// キー押し込み
func PressEnter() bool { return display.GetPressedKey() == 0x101 }

// Option 1（上部の左ボタン）
func PressOption1() bool { return display.GetPressedKey() == '2' }

// Option 2（上部の中央ボタン）
func PressOption2() bool { return display.GetPressedKey() == '1' }

// Option 3（上部の右）
func PressOption3() bool { return display.GetPressedKey() == '0' }
