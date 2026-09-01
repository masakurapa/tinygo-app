//go:build wioterminal

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/ili9341"
)

type note struct {
	tone     float64
	duration float64
}

var (
	rightButton machine.Pin
	leftButton  machine.Pin
	downButton  machine.Pin
	upButton    machine.Pin
	opt1Button  machine.Pin
	opt2Button  machine.Pin
	opt3Button  machine.Pin

	bzr         buzzer.Device
	buzzerNotes = []note{
		{buzzer.C4, buzzer.Quarter},
		{buzzer.F4, buzzer.Half},
		{buzzer.F4, buzzer.Eighth},
		{buzzer.F4, buzzer.Quarter},
		{buzzer.A4, buzzer.Quarter},
		{buzzer.G4, buzzer.Half},
		{buzzer.F4, buzzer.Eighth},
		{buzzer.G4, buzzer.Quarter},
		{buzzer.A4, buzzer.Quarter},
		{buzzer.F4, buzzer.Half},
		{buzzer.F4, buzzer.Eighth},
		{buzzer.A4, buzzer.Quarter},
		{buzzer.C5, buzzer.Quarter},
		{buzzer.D5, buzzer.Half},
	}
)

func initDisplay() displayer {
	bzrPin := machine.WIO_BUZZER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr = buzzer.New(bzrPin)

	btnX, btnY, btnZ, btnB := machine.SWITCH_X, machine.SWITCH_Y, machine.SWITCH_Z, machine.SWITCH_B
	btn1, btn2, btn3 := machine.BUTTON_1, machine.BUTTON_2, machine.BUTTON_3

	rightButton = btnZ
	btnZ.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	leftButton = btnY
	btnY.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	downButton = btnB
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	upButton = btnX
	btnX.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	opt1Button = btn1
	opt1Button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	opt2Button = btn2
	opt2Button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	opt3Button = btn3
	opt3Button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	sck, sdo, sdi := machine.LCD_SCK_PIN, machine.LCD_SDO_PIN, machine.LCD_SDI_PIN
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       sck,
		SDO:       sdo,
		SDI:       sdi,
		Frequency: 60000000,
	})

	spi3 := *machine.SPI3
	dc, cs, rst := machine.LCD_DC, machine.LCD_SS_PIN, machine.LCD_RESET
	d := ili9341.NewSPI(spi3, dc, cs, rst)
	d.Configure(ili9341.Config{
		Rotation: ili9341.Rotation270,
	})
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.LCD_BACKLIGHT.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.LCD_BACKLIGHT.High()

	return d
}

func PressKeyRight() bool { return !rightButton.Get() }
func PressKeyLeft() bool  { return !leftButton.Get() }
func PressKeyDown() bool  { return !downButton.Get() }
func PressKeyUp() bool    { return !upButton.Get() }

// Option 1（上部の左ボタン）
func PressOption1() bool { return !opt3Button.Get() }

// Option 2（上部の中央ボタン）
func PressOption2() bool { return !opt2Button.Get() }

// Option 3（上部の右）
func PressOption3() bool { return !opt1Button.Get() }

func Buzzer() {
	for _, n := range buzzerNotes {
		bzr.Tone(n.tone, n.duration)
		time.Sleep(10 * time.Millisecond)
	}
}
