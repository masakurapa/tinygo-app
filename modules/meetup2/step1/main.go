package main

import (
	"time"

	"github.com/masakurapa/tinygo-app/pkg/color"
	"github.com/masakurapa/tinygo-app/pkg/devices"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func main() {
	tick := time.Tick(16 * time.Millisecond)
	disp := devices.New()

	for {
		<-tick
		tinyfont.WriteLine(disp, &freesans.Regular18pt7b, 15, 140, "Push stick to start!!", color.Black)
	}
}
