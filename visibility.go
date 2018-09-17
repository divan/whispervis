package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

// VisibilityListener implements listener for visibilitychange events.
// We use it to pause animation when the page is hidden, because WebGL
// animation is pretty CPU expensive.
// See https://developer.mozilla.org/en-US/docs/Web/API/Page_Visibility_API for details.
func (p *Page) VisibilityListener(e *vecty.Event) {
	document := js.Global.Get("document")
	hidden := document.Get("hidden")
	fmt.Println("Page is hidden:", hidden)
}
