//go:build !wioterminal

package main

import (
	"cmp"
	"image/color"
	"os"
	"time"

	"github.com/sago35/tinydisplay"
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var display *initdisplay.TinyDisplay

type tinyDisplay struct{}

func (t *tinyDisplay) Size() (x, y int16)            { return display.Size() }
func (t *tinyDisplay) SetPixel(x, y int16, c color.RGBA) { display.SetPixel(x, y, c) }
func (t *tinyDisplay) Display() error                { return display.Display() }
func (t *tinyDisplay) DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error {
	return display.DrawRGBBitmap(x, y, data, w, h)
}
func (t *tinyDisplay) FillScreen(c color.RGBA) { display.FillScreen(c) }

var disp = &tinyDisplay{}

func initDisplay() {
	d, _ := tinydisplay.NewClient(cmp.Or(os.Getenv("TINYDISPLAY_HOST"), "localhost"), 9812, 320, 240)
	d.FillScreen(color.RGBA{0, 0, 0, 255})
	time.Sleep(100 * time.Millisecond)
	d.FillScreen(color.RGBA{255, 255, 255, 255})
	display = &initdisplay.TinyDisplay{Client: d}
}

func pressKeyRight() bool { return display.GetPressedKey() == 0x106 }
func pressKeyLeft() bool  { return display.GetPressedKey() == 0x107 }
func pressKeyUp() bool    { return display.GetPressedKey() == 0x109 }
func pressEnter() bool    { return display.GetPressedKey() == 0x101 }

func calibrateAccel()    {}
func accelY() float64    { return 0 }
func supportAccel() bool { return false }

func waitRelease(pressed func() bool) {
	for pressed() {
		time.Sleep(16 * time.Millisecond)
	}
	time.Sleep(16 * 3 * time.Millisecond)
}
