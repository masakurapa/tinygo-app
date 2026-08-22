package main

import (
	"github.com/masakurapa/tinygo-app/apps/dropcircle/game"
	"github.com/sago35/koebiten"
	"github.com/sago35/koebiten/hardware"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			koebiten.Println("panic!!")
			koebiten.Println(r)
		}
	}()

	koebiten.SetHardware(hardware.Device)
	if err := koebiten.RunGame(game.New()); err != nil {
		panic(err)
	}
}
