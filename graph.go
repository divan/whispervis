package main

import (
	"runtime"

	"github.com/gopherjs/vecty"
)

// UpdateGraph starts graph layout simulation.
func (p *Page) UpdateGraph() {
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

	p.RecreateObjects()
}

// RecreateObjects removes (if any) objects from the WebGL scene,
// and creates all new objects based on current data and positions.
// It doesn't do any calculations of changes to the layout or graph data.
func (p *Page) RecreateObjects() {
	p.loaded = true
	p.loader.Reset()

	p.webgl.RemoveObjects()
	p.webgl.CreateObjects(p.layout.Positions(), p.layout.Links())

	vecty.Rerender(p)
}

// ApplyForces applies current forces to the objects, and runs
// a single simulation run to update positions.
func (p *Page) ApplyForces() {
	fc := p.forceEditor.Config()
	p.layout.SetConfig(fc.Config)
	p.layout.UpdatePositions()
	p.webgl.updatePositions()
	p.webgl.rt.Disable()
}
