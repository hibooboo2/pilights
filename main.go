package main

import (
	"context"
	"math/rand"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	pins := []rpio.Pin{}
	pins = append(pins, rpio.Pin(26))
	pins = append(pins, rpio.Pin(11))
	pins = append(pins, rpio.Pin(2))
	pins = append(pins, rpio.Pin(3))
	pins = append(pins, rpio.Pin(9))

	defer func() {
		for _, v := range pins {
			defer v.Low()
		}

		rpio.Close()
		println("Cleaned up")
	}()

	for _, p := range pins {
		p.Output()
	}

	ctx := context.Background()
	for _, p := range pins {
		go func(p rpio.Pin, ctx context.Context) {
			for {
				p.Toggle()
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
			}
		}(p, ctx)
	}
	time.Sleep(time.Second * 5)
}

