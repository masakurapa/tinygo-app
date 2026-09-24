package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var black = color.RGBA{0, 0, 0, 255}

func main() {
	tick := time.Tick(16 * time.Millisecond)
	initDisplay()

	for {
		<-tick
		tinyfont.WriteLine(disp, &freesans.Regular18pt7b, 15, 140, "Push stick to start!!", black)
	}
}
