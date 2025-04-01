package main

import (
	"github.com/masakurapa/tinygo-app/games/dropcircle/game"
	"github.com/sago35/koebiten"
	"github.com/sago35/koebiten/hardware"
)

func main() {
	koebiten.SetHardware(hardware.Device)
	if err := koebiten.RunGame(game.New()); err != nil {
		panic(err)
	}
}
