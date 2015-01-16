package main

import (
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
)

func gfxLoop(w window.Window, r gfx.Renderer) {

}

func main() {

	props := window.NewProps()

	window.Run(gfxLoop, props)
}
