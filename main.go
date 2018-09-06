package main

import (
	"bytes"
	"time"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/lngramos/three"
	"github.com/status-im/whispervis/widgets"
	"github.com/vecty/vthree"
)

func main() {
	buf := bytes.NewBuffer(inputJSON)
	data, err := formats.FromD3JSONReader(buf)
	if err != nil {
		panic(err)
	}

	l := layout.NewAuto(data)
	steps := 50
	page := &Page{
		layout: l,
		loader: widgets.NewLoader(steps),
	}

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/pure-min.css")
	vecty.AddStylesheet("css/controls.css")
	vecty.RenderBody(page)

	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < steps; i++ {
			l.UpdatePositions()
			page.loader.Inc()
			vecty.Rerender(page.loader)
		}
		page.loaded = true
		vecty.Rerender(page)
	}()
}

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

	loaded bool
	loader *widgets.Loader
}

// Render implements the vecty.Component interface.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("pure-g")),

			elem.Div(
				vecty.Markup(vecty.Class("pure-u-1-5")),
				elem.Heading1(vecty.Text("Whisper Message Propagation")),
				elem.Paragraph(vecty.Text("This visualization represents message propagation in the p2p network.")),
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
