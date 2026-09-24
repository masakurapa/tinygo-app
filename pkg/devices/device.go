package devices

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers"
)

type Displayer interface {
	drivers.Displayer
	DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error
	FillScreen(c color.RGBA)
}

func WaitRelease(pressed func() bool) {
	for pressed() {
		time.Sleep(16 * time.Millisecond)
	}
	time.Sleep(16 * 3 * time.Millisecond)
}
