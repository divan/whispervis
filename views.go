package main

import (
	"github.com/gopherjs/vecty"
)

// known views
const (
	View3D    = "3d"
	ViewStats = "stats"
	ViewFAQ   = "faq"
)

// onTabSwitch returns a proper tab switching function depending on the tab clicked.
func (p *Page) onTabSwitch(view string) func(e *vecty.Event) {
	if p.activeView == view {
		return nil
	}
	return func(e *vecty.Event) {
		p.switchView(view)
	}
}

func (p *Page) switchView(view string) {
	p.activeView = view
	vecty.Rerender(p)
}
