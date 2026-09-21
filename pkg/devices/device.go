package devices

import (
	"time"

	"tinygo.org/x/drivers"
)

type Displayer interface {
	drivers.Displayer
	DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error
}

func WaitRelease(pressed func() bool) {
	for pressed() {
		time.Sleep(16 * time.Millisecond)
	}
	time.Sleep(16 * 3 * time.Millisecond)
}
