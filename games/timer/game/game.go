package game

import (
	"fmt"
	"time"

	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
)

var (
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

func New() koebiten.Game {
	return &game{}
}

type game struct {
	start     time.Time
	result    time.Duration
	isPressed bool
}

func (g *game) Update() error {
	if koebiten.IsKeyJustPressed(koebiten.Key3) {
		if g.result != 0 {
			g.result = time.Duration(0)
			return nil
		}

		g.start = time.Now()
		g.result = time.Duration(0)
		g.isPressed = true
		return nil
	}

	if koebiten.IsKeyJustReleased(koebiten.Key3) {
		if !g.isPressed {
			return nil
		}

		end := time.Now()
		g.result = end.Sub(g.start)
		g.isPressed = false
		return nil
	}

	return nil
}

func (g *game) Draw(screen *koebiten.Image) {
	if g.isPressed {
		koebiten.Println("Joystick is pressed")
		return
	}

	if g.result == 0 {
		koebiten.Println("Press the joystick\nto start!!")
		return
	}

	koebiten.Println("Your score is ...")
	koebiten.Println(fmt.Sprintf("%.3f sec!!", g.result.Seconds()))
	koebiten.Println()
	koebiten.Println()
	koebiten.Println()
	koebiten.Println("Press the joystick\nto reset")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideHeight, outsideHeight
}
