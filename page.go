package main

import (
	"fmt"

	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/status-im/whispervis/widgets"
)

// Page is our main page component.
type Page struct {
	vecty.Core

	layout *layout.Layout

	webgl *WebGLScene

	loaded       bool
	isSimulating bool

	loader           *widgets.Loader
	forceEditor      *widgets.ForceEditor
	network          *NetworkSelector
	simulationWidget *widgets.Simulation
	statsWidget      *widgets.Stats

	simulation *Simulation
}

// NewPage creates and inits new app page.
func NewPage() *Page {
	page := &Page{
		loader: widgets.NewLoader(),
	}
	page.forceEditor = widgets.NewForceEditor(page.onForcesApply)
	page.network = NewNetworkSelector(page.onNetworkChange)
	page.webgl = NewWebGLScene()
	page.simulationWidget = widgets.NewSimulation("http://localhost:8084", page.startSimulation, page.replaySimulation)
	page.statsWidget = widgets.NewStats()
	return page
}

// Render implements the vecty.Component interface.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("columns"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column", "is-narrow"),
					vecty.Style("width", "300px"),
				),
				p.header(),
				elem.Div(
					vecty.Markup(
						vecty.MarkupIf(p.isSimulating,
							// disable
							vecty.Style("pointer-events", "none"),
							vecty.Style("opacity", "0.4"),
						),
					),
					p.network,
					p.forceEditor,
				),
				elem.Div(
					vecty.Markup(
						vecty.MarkupIf(!p.loaded, vecty.Style("visibility", "hidden")),
					),
					p.simulationWidget,
					elem.Div(
						vecty.Markup(
							vecty.MarkupIf(p.isSimulating,
								// disable
								vecty.Style("pointer-events", "none"),
								vecty.Style("opacity", "0.4"),
							),
						),
						p.statsWidget,
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column"),
					/*
						we use display:none property to hide WebGL instead of mounting/unmounting,
						because we want to create only one WebGL context and reuse it. Plus,
						WebGL takes time to initialize, so it can do it being hidden.
					*/
					vecty.MarkupIf(!p.loaded,
						vecty.Style("visibility", "hidden"),
						vecty.Style("height", "0px"),
						vecty.Style("width", "0px"),
					),
				),
				p.webgl,
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("pure-u-4-5"),
				),
				vecty.If(!p.loaded, p.loader),
			),
		),
		vecty.Markup(
			event.KeyDown(p.KeyListener),
			event.VisibilityChange(p.VisibilityListener),
		),
	)
}

// onForcesApply executes when Force Editor click is fired.
func (p *Page) onForcesApply() {
	if !p.loaded {
		return
	}
	p.UpdateGraph()
}

func (p *Page) onNetworkChange(network *Network) {
	fmt.Println("Network changed:", network)
	config := p.forceEditor.Config()
	p.layout = layout.NewFromConfig(network.Data, config.Config)
	go p.UpdateGraph()
}

// startSimulation is called on the end of each simulation round.
func (p *Page) startSimulation() error {
	p.isSimulating = true
	vecty.Rerender(p)

	defer func() {
		p.isSimulating = false
		vecty.Rerender(p)
	}()

	backend := p.simulationWidget.Address()
	sim, err := p.runSimulation(backend)
	if err != nil {
		return err
	}

	// calculate stats and update stats widget
	stats := p.RecalculateStats(sim.plog)
	p.statsWidget.Update(stats)

	sim.stats = stats
	p.simulation = sim

	p.replaySimulation()
	return nil
}

// replaySimulation animates last simulation.
func (p *Page) replaySimulation() {
	if p.simulation == nil {
		return
	}
	p.webgl.AnimatePropagation(p.simulation.plog)
}

func (p *Page) header() *vecty.HTML {
	return elem.Section(
		elem.Heading2(
			vecty.Markup(
				vecty.Class("title", "has-text-weight-light"),
			),
			vecty.Text("Whisper Message Propagation"),
		),
		elem.Heading5(
			vecty.Markup(
				vecty.Class("subtitle", "has-text-weight-light"),
			),
			vecty.Text("This visualization represents message propagation in the p2p network."),
		),
	)
}
