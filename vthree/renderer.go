package vthree

import (
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type WebGLRenderer struct {
	vecty.Core
	opts   WebGLOptions          `vecty:"prop"`
	markup []vecty.MarkupOrChild `vecty:"prop"`

	canvas   *vecty.HTML
	renderer *three.WebGLRenderer
}

// Mount implements the vecty.Mounter interface.
func (r *WebGLRenderer) Mount() {
	r.renderer = newWebGLRenderer(&webGLRendererParameters{
		Canvas:       r.canvas.Node(),
		WebGLOptions: &r.opts,
	})
	r.opts.Init(r.renderer)
}

// Unmount implements the vecty.Unmounter interface.
func (r *WebGLRenderer) Unmount() {
	if r.opts.Shutdown != nil {
		r.opts.Shutdown(r.renderer)
	}
}

// Render implements the vecty.Component interface.
func (r *WebGLRenderer) Render() vecty.ComponentOrHTML {
	r.canvas = elem.Canvas(r.markup...)
	return r.canvas
}

// WebGLOptions represent options for the WebGLRenderer component.
type WebGLOptions struct {
	// Init is called when the three.js WebGLRenderer has been created.
	//
	// This can happen multiple times during the lifecycle of an application
	// if the Vecty WebGLRenderer component was unmounted and mounted again,
	// e.g. due to navigating to a different page and back again.
	Init func(r *three.WebGLRenderer)

	// Shutdown is called before the canvas associated with the three.js
	// WebGLRenderer will be destroyed. For example, when your Vecty
	// application no longer renders the WebGLRenderer component and it is
	// being unmounted.
	Shutdown func(r *three.WebGLRenderer)

	// TODO(slimsag): allow specifying other parameters like context, precision,
	// etc. from three.js WebGLRenderer constructor here:
	// https://threejs.org/docs/#api/renderers/WebGLRenderer
	Antialias bool
}

// NewWebGLRenderer returns a Vecty component that initializes a three.js WebGL renderer for
// use in a Vecty application.
func NewWebGLRenderer(opts WebGLOptions, markup ...vecty.MarkupOrChild) *WebGLRenderer {
	if opts.Init == nil {
		panic("vthree: Renderer: must specify opts.Init")
	}
	return &WebGLRenderer{
		opts:   opts,
		markup: markup,
	}
}

type webGLRendererParameters struct {
	Canvas *js.Object
	*WebGLOptions
}

// Note: We can't use three.NewWebGLRenderer because it doesn't allow
// specifying any parameters yet. Easy enough to just call ourself, though.
func newWebGLRenderer(parameters *webGLRendererParameters) *three.WebGLRenderer {
	return &three.WebGLRenderer{
		Object: js.Global.Get("THREE").Get("WebGLRenderer").New(map[string]interface{}{
			"canvas":    parameters.Canvas,
			"antialias": parameters.Antialias,
		}),
	}
}
