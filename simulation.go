package main

import (
	"runtime"

	"github.com/gopherjs/vecty"
)

// StartSimulation starts graph layout simulation.
func (p *Page) StartSimulation() {
	p.loader.Reset()
	p.loaded = false
	vecty.Rerender(p)

	config := p.forceEditor.Config()
	p.loader.SetSteps(config.Steps)
	for i := 0; i < config.Steps; i++ {
		p.layout.UpdatePositions()
		p.loader.Inc()
		vecty.Rerender(p.loader)
		runtime.Gosched()
	}
	p.loaded = true
	p.webgl.CreateObjects(p.layout.Positions(), p.layout.Links())
	vecty.Rerender(p)
}
