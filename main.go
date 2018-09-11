package main

import (
	"github.com/gopherjs/vecty"
)

func main() {
	page := NewPage(50)

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/pure-min.css")
	vecty.AddStylesheet("css/controls.css")
	vecty.RenderBody(page)

	page.UpdateNetworkGraph(inputJSON)
}
