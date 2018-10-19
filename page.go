package main

import (
	"bytes"
	"fmt"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/status-im/simulation/propagation"
	"github.com/status-im/whispervis/widgets"
)

// Page is our main page component.
type Page struct {
	vecty.Core

	layout *layout.Layout

	webgl *WebGLScene

	loaded bool

	loader         *widgets.Loader
	forceEditor    *widgets.ForceEditor
	network        *NetworkSelector
	simulationConf *widgets.Simulation
	statsWidget    *widgets.Stats
}

// NewPage creates and inits new app page.
func NewPage() *Page {
	page := &Page{
		loader:      widgets.NewLoader(),
		forceEditor: widgets.NewForceEditor(),
	}
	page.network = NewNetworkSelector(page.onNetworkChange)
	page.webgl = NewWebGLScene()
	page.simulationConf = widgets.NewSimulation("localhost:8084", page.CurrentNetwork, page.onSimulationFinish)
	page.statsWidget = widgets.NewStats()
	return page
}

// Render implements the vecty.Component interface.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("pure-g"),
				vecty.Style("height", "100%"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("pure-u-1-5")),
				elem.Heading1(vecty.Text("Whisper Message Propagation")),
				elem.Paragraph(vecty.Text("This visualization represents message propagation in the p2p network.")),
				p.network,
				elem.HorizontalRule(),
				elem.Div(
					vecty.Markup(
						vecty.MarkupIf(!p.loaded, vecty.Style("visibility", "hidden")),
					),
					p.simulationConf,
					elem.HorizontalRule(),
					p.forceEditor,
					p.updateButton(),
					elem.HorizontalRule(),
					p.statsWidget,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("pure-u-4-5"),
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

func (p *Page) updateButton() *vecty.HTML {
	return elem.Div(
		elem.Button(
			vecty.Markup(
				vecty.Class("pure-button"),
				vecty.Style("background", "rgb(28, 184, 65)"),
				vecty.Style("color", "white"),
				vecty.Style("border-radius", "4px"),
				event.Click(p.onUpdateClick),
			),
			vecty.Text("Update"),
		),
	)
}

func (p *Page) onUpdateClick(e *vecty.Event) {
	if !p.loaded {
		return
	}
	go p.UpdateGraph()
}

func (p *Page) onNetworkChange(network *Network) {
	fmt.Println("Network changed:", network)
	config := p.forceEditor.Config()
	p.layout = layout.NewFromConfig(network.Data, config.Config)
	go p.UpdateGraph()
}

// CurrentNetwork returns JSON encoded description of the current graph/network.
func (p *Page) CurrentNetwork() []byte {
	net := p.network.current.Data
	var buf bytes.Buffer
	err := formats.NewD3JSON(&buf, true).ExportGraph(net)
	if err != nil {
		fmt.Println("[ERROR] Can't export graph:", err)
		return nil
	}
	return buf.Bytes()
}

// onSimulationFinish is called on the end of each simulation round.
func (p *Page) onSimulationFinish(plog *propagation.Log) {
	net := p.network.current
	nodesCount := len(net.Data.Nodes())
	linksCount := len(net.Data.Links())
	p.statsWidget.Update(plog, nodesCount, linksCount)

	p.webgl.AnimatePropagation(plog)
}
