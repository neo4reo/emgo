package main

import (
	"delay"

	"stm32/hal/gpio"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var leds *gpio.Port

const (
	led1 = gpio.Pin3
	led2 = gpio.Pin5
)

func init() {
	system.Setup192(8)
	systick.Setup()

	gpio.H.EnableClock(false)
	leds = gpio.H

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	leds.Setup(led1|led2, &cfg)
}

func wait() {
	//delay.Loop(1e7)
	delay.Millisec(500)
}

func main() {
	for {
		leds.ClearPins(led1)
		leds.SetPins(led2)
		wait()

		leds.ClearPins(led2)
		leds.SetPins(led1)
		wait()
	}
}