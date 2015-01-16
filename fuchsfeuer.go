package main

import (
	"fmt"
	"time"

	"azul3d.org/clock.v1"
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
)

const (
	targetFrameRate = 60.0
	secondsPerFrame = 1.0 / targetFrameRate

	secToNanoMultiplikator               = 1000000000
	durationFrame          time.Duration = 1 * secToNanoMultiplikator / 60 * time.Nanosecond
)

func gfxLoop(w window.Window, r gfx.Renderer) {

	clock := clock.New()
	delay := clock.Time().Seconds()
	last := delay

	for {
		clock.Tick()
		elapsedTime := (clock.Time().Seconds() - last)
		last = clock.Time().Seconds()
		delay += elapsedTime

		// Process input here
		fmt.Println("Input")

		for delay >= secondsPerFrame {
			delay -= secondsPerFrame
			// Update here
			fmt.Println("Update")
		}

		r.Clear(r.Bounds(), gfx.Color{R: 1, G: 1, B: 1, A: 1})

		// Render here
		fmt.Println("Render")

		r.Render()
	}

}

func main() {

	props := window.NewProps()

	window.Run(gfxLoop, props)
}
