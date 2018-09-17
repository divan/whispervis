package main

import "github.com/gopherjs/vecty"

// KeyListener implements listener for keydown events.
func (p *Page) KeyListener(e *vecty.Event) {
	key := e.Get("key").String()
	switch key {
	case "p":
		p.webgl.ToggleAutoRotation()
	}
}
