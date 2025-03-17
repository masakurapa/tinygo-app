package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		fmt.Println("Hello World")
		led.Toggle()
		time.Sleep(1000 * time.Microsecond)
	}
}
