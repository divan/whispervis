package main

import (
	"bytes"
	"fmt"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/graph"
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/lngramos/three"
	"github.com/status-im/whispervis/widgets"
	"github.com/vecty/vthree"
)

// Page is our main page component.
type Page struct {
	vecty.Core

	layout *layout.Layout

	scene    *three.Scene
	camera   three.PerspectiveCamera
	renderer *three.WebGLRenderer
	graph    *three.Group
	nodes    *three.Group
	edges    *three.Group
	controls TrackBallControl

	autoRotate bool

	loaded      bool
	loader      *widgets.Loader
	forceEditor *widgets.ForceEditor
	upload      *widgets.UploadWidget
	data        *graph.Graph
}

// NewPage creates and inits new app page.
func NewPage(steps int) *Page {
	page := &Page{
		loader:      widgets.NewLoader(steps),
		forceEditor: widgets.NewForceEditor(),
	}
	page.upload = widgets.NewUploadWidget(page.onUpload)
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
				vecty.Markup(vecty.Class("pure-u-4-5")),
				vecty.If(p.loaded,
					vthree.WebGLRenderer(vthree.WebGLOptions{
						Init:     p.init,
						Shutdown: p.shutdown,
					}),
				),
				vecty.If(!p.loaded, p.loader),
			),
		),
		vecty.Markup(
			event.KeyDown(p.KeyListener),
		),
	)
}

func (p *Page) renderWebGLCanvas() vecty.Component {
	return vthree.WebGLRenderer(vthree.WebGLOptions{
		Init:     p.init,
		Shutdown: p.shutdown,
	})
}

func (p *Page) init(renderer *three.WebGLRenderer) {
	windowWidth := js.Global.Get("innerWidth").Float()*80/100 - 20
	windowHeight := js.Global.Get("innerHeight").Float() - 20

	p.renderer = renderer
	p.renderer.SetSize(windowWidth, windowHeight, true)

	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()
	p.renderer.SetPixelRatio(devicePixelRatio)

	p.InitScene(windowWidth, windowHeight)

	p.CreateObjects()

	p.animate()
}

func (p *Page) shutdown(renderer *three.WebGLRenderer) {
	p.scene = nil
	p.camera = three.PerspectiveCamera{}
	p.renderer = nil
	p.graph, p.nodes, p.edges = nil, nil, nil
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

func (p *Page) UpdateNetworkGraph(json []byte) error {
	buf := bytes.NewBuffer(json)
	data, err := formats.FromD3JSONReader(buf)
	if err != nil {
		return fmt.Errorf("update graph: %v", err)
	}

	p.data = data
	config := p.forceEditor.Config()
	p.layout = layout.NewFromConfig(data, config)
	go p.StartSimulation()
	return nil
}
