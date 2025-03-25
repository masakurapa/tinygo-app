package main

import (
	"github.com/masakurapa/tinygo-app/games/marumaru"
	"github.com/sago35/koebiten"
	"github.com/sago35/koebiten/hardware"
)

func main() {
	koebiten.SetHardware(hardware.Device)
	if err := koebiten.RunGame(marumaru.New()); err != nil {
		panic(err)
	}
}
