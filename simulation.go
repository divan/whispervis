package main

import (
	"runtime"

	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
)

// StartSimulation starts graph layout simulation.
func (p *Page) StartSimulation() {
	p.loader.Reset()
	p.loaded = false
	vecty.Rerender(p)

	config := p.forceEditor.Config()
	l := layout.NewFromConfig(p.data, config)
	p.layout = l
	for i := 0; i < p.loader.Steps(); i++ {
		p.layout.UpdatePositions()
		p.loader.Inc()
		vecty.Rerender(p.loader)
		runtime.Gosched()
	}
	p.loaded = true
	vecty.Rerender(p)
}
