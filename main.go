package main

import (
	"github.com/gopherjs/vecty"
)

func main() {
	page := NewPage()

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("css/bulma.css")
	vecty.AddStylesheet("css/bulma-extensions.min.css")
	vecty.AddStylesheet("css/custom.css")
	vecty.RenderBody(page)
}
