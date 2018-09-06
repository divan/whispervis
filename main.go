package main

import (
	"bytes"

	"github.com/divan/graphx/formats"
	"github.com/gopherjs/vecty"
)

func main() {
	buf := bytes.NewBuffer(inputJSON)
	data, err := formats.FromD3JSONReader(buf)
	if err != nil {
		panic(err)
	}

	page := NewPage(data, 50)

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/pure-min.css")
	vecty.AddStylesheet("css/controls.css")
	vecty.RenderBody(page)

	page.StartSimulation()
}
