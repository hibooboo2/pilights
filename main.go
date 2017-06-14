package main

import (
	"math/rand"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	pin := rpio.Pin(26)
	pin2 := rpio.Pin(11)
	defer pin.Low()
	defer pin2.Low()
	defer rpio.Close()

	pin.Output()
	pin2.Output()
	pin2.High()
	for {
		pin.Toggle()
		pin2.Toggle()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	}
}

