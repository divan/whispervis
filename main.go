package main

import (
	"bytes"
	"runtime"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/status-im/whispervis/widgets"
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
		layout:      l,
		loader:      widgets.NewLoader(steps),
		forceEditor: widgets.NewForceEditor(),
	}

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/pure-min.css")
	vecty.AddStylesheet("css/controls.css")
	vecty.RenderBody(page)

	go func() {
		for i := 0; i < steps; i++ {
			l.UpdatePositions()
			page.loader.Inc()
			runtime.Gosched()
		}
		page.loaded = true
		vecty.Rerender(page)
	}()
}
