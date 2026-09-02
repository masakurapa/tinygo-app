package main

import (
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func displayGame(disp displayer, lab *label) {
	// switch mode
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQR()
		return
	}

	lab.FillScreen(white)

	tinyfont.WriteLine(lab, &freesans.Regular24pt7b, 10, 40, "wip: play GAME", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

func switchGame() {
	currentMode = modeGame
}
