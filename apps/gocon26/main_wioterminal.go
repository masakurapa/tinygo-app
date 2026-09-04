//go:build wioterminal

package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/drivers/lis3dh"
)

type note struct {
	tone     float64
	duration float64
}

var (
	rightButton     machine.Pin
	leftButton      machine.Pin
	downButton      machine.Pin
	upButton        machine.Pin
	stickPushButton machine.Pin
	opt1Button      machine.Pin
	opt2Button      machine.Pin
	opt3Button      machine.Pin

	accel      lis3dh.Device
	atanOffset float64
)

func initDisplay() displayer {
	btnX, btnY, btnZ, btnB, btnU := machine.SWITCH_X, machine.SWITCH_Y, machine.SWITCH_Z, machine.SWITCH_B, machine.SWITCH_U
	// ジョイスティックを押したとき
	btn1, btn2, btn3 := machine.BUTTON_1, machine.BUTTON_2, machine.BUTTON_3

	rightButton = btnZ
	btnZ.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	leftButton = btnY
	btnY.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	downButton = btnB
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	upButton = btnX
	btnX.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	stickPushButton = btnU
	btnU.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
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

	// 加速度情報の取得
	scl, sda := machine.SCL0_PIN, machine.SDA0_PIN
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{SCL: scl, SDA: sda})
	accel = lis3dh.New(i2c)
	accel.Address = lis3dh.Address0
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

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
func PressEnter() bool    { return !stickPushButton.Get() }

// Option 1（上部の左ボタン）
func PressOption1() bool { return !opt3Button.Get() }

// Option 2（上部の中央ボタン）
func PressOption2() bool { return !opt2Button.Get() }

// Option 3（上部の右）
func PressOption3() bool { return !opt1Button.Get() }

func CalibrateAtan() {
	_, y, _, _ := accel.ReadAcceleration()
	atanOffset = float64(y) / 1000000
}

// 画面の傾き（-1g〜1g の範囲）
func AtanY() float64 {
	_, y, _, _ := accel.ReadAcceleration()
	return float64(y)/1000000 - atanOffset
}

func SupportAtan() bool {
	return true
}
