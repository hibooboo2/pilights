package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
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

	for _, p := range pins {
		p.Output()
	}

	ctx := context.Background()
	ctx, c := context.WithCancel(ctx)
	cancels := []func(){c}
	for _, p := range pins {
		ctxp, c2 := context.WithCancel(ctx)
		cancels = append(cancels, c2)
		go func(p rpio.Pin, ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					p.Toggle()
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
				}
			}
		}(p, ctxp)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	for _ = range stop {
		for _, v := range pins {
			v.Low()
		}
		for _, can := range cancels {
			can()
		}
		rpio.Close()
		println("Cleaned up")
		return
	}
}

