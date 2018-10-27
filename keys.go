package main

import (
	"github.com/gopherjs/vecty"
)

// KeyListener implements listener for keydown events.
func (p *Page) KeyListener(e *vecty.Event) {
	p.webgl.rt.EnableRendering()
	key := e.Get("key").String()
	switch key {
	case "p":
		p.webgl.ToggleAutoRotation()
	case "o":
		p.webgl.ToggleWobbling()
	case "f":
		p.ApplyForces()
	case "]":
		p.simulationWidget.StepForward()
	case "[":
		p.simulationWidget.StepBackward()
	}
}
