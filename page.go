package main

import (
	"bytes"
	"fmt"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/graph"
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

	loaded bool

	loader      *widgets.Loader
	forceEditor *widgets.ForceEditor
	upload      *widgets.UploadWidget
	network     *NetworkSelector

	data *graph.Graph
}

// NewPage creates and inits new app page.
func NewPage() *Page {
	page := &Page{
		loader:      widgets.NewLoader(),
		forceEditor: widgets.NewForceEditor(),
		network:     NewNetworkSelector("net100"),
	}
	page.upload = widgets.NewUploadWidget(page.onUpload)
	page.webgl = NewWebGLScene()
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
				p.upload,
				elem.Div(
					vecty.Markup(
						vecty.MarkupIf(!p.loaded, vecty.Style("visibility", "hidden")),
					),
					p.forceEditor,
					p.updateButton(),
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
	go p.StartSimulation()
}

// UpdateNetworkGraph updates graph and scene with new data.
func (p *Page) UpdateNetworkGraph(json []byte) error {
	buf := bytes.NewBuffer(json)
	data, err := formats.FromD3JSONReader(buf)
	if err != nil {
		return fmt.Errorf("update graph: %v", err)
	}

	p.data = data
	config := p.forceEditor.Config()
	p.layout = layout.NewFromConfig(data, config.Config)
	go p.StartSimulation()
	return nil
}
