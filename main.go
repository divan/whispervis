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

	forceEditor := widgets.NewForceEditor()
	config := forceEditor.Config()
	l := layout.NewFromConfig(data, config)
	steps := 50
	page := &Page{
		layout:      l,
		loader:      widgets.NewLoader(steps),
		forceEditor: forceEditor,
	}

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/pure-min.css")
	vecty.AddStylesheet("css/controls.css")
	vecty.RenderBody(page)

	go func() {
		for i := 0; i < steps; i++ {
			l.UpdatePositions()
			page.loader.Inc()
			vecty.Rerender(page.loader)
			runtime.Gosched()
		}
		page.loaded = true
		vecty.Rerender(page)
	}()
}
