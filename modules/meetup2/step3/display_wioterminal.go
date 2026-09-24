//go:build wioterminal

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/drivers/lis3dh"
)

var (
	rightButton     machine.Pin
	leftButton      machine.Pin
	upButton        machine.Pin
	stickPushButton machine.Pin

	accel       lis3dh.Device
	accelOffset float64
)

var disp *ili9341.Device

func initDisplay() {
	btnX, btnY, btnZ, btnU := machine.SWITCH_X, machine.SWITCH_Y, machine.SWITCH_Z, machine.SWITCH_U

	rightButton = btnZ
	btnZ.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	leftButton = btnY
	btnY.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	upButton = btnX
	btnX.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	stickPushButton = btnU
	btnU.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	scl, sda := machine.SCL0_PIN, machine.SDA0_PIN
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{SCL: scl, SDA: sda})
	accel = lis3dh.New(i2c)
	accel.Address = lis3dh.Address0
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	sck, sdo, sdi := machine.LCD_SCK_PIN, machine.LCD_SDO_PIN, machine.LCD_SDI_PIN
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       sck,
		SDO:       sdo,
		SDI:       sdi,
		Frequency: 60000000,
	})

	spi3 := *machine.SPI3
	dc, cs, rst := machine.LCD_DC, machine.LCD_SS_PIN, machine.LCD_RESET
	disp = ili9341.NewSPI(spi3, dc, cs, rst)
	disp.Configure(ili9341.Config{Rotation: ili9341.Rotation270})
	disp.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.LCD_BACKLIGHT.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.LCD_BACKLIGHT.High()
}

func pressKeyRight() bool { return !rightButton.Get() }
func pressKeyLeft() bool  { return !leftButton.Get() }
func pressKeyUp() bool    { return !upButton.Get() }
func pressEnter() bool    { return !stickPushButton.Get() }

func calibrateAccel() {
	_, y, _, _ := accel.ReadAcceleration()
	accelOffset = float64(y) / 1000000
}

func accelY() float64 {
	_, y, _, _ := accel.ReadAcceleration()
	return float64(y)/1000000 - accelOffset
}

func supportAccel() bool { return true }

func waitRelease(pressed func() bool) {
	for pressed() {
		time.Sleep(16 * time.Millisecond)
	}
	time.Sleep(16 * 3 * time.Millisecond)
}
