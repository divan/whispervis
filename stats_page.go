package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// StatsPage is stats view component.
type StatsPage struct {
	vecty.Core

	width, height string
}

// NewStatsPage creates and inits new stats page.
func NewStatsPage() *StatsPage {
	width, height := PageViewSize()
	return &StatsPage{
		width:  fmt.Sprintf("%dpx", width),
		height: fmt.Sprintf("%dpx", height),
	}
}

// Render implements the vecty.Component interface.
func (s *StatsPage) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("width", s.width),
			vecty.Style("height", s.height),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("title", "has-text-centered"),
			),
			vecty.Text("Stats page"),
		),
		// consult this tile madness here https://bulma.io/documentation/layout/tiles/
		elem.Div(vecty.Markup(vecty.Class("tile", "is-anscestor")),
			elem.Div(vecty.Markup(vecty.Class("tile", "is-parent", "is-4", "is-vertical")),
				elem.Div(vecty.Markup(vecty.Class("tile")),
					elem.Div(vecty.Markup(vecty.Class("tile", "is-child", "box")),
						vecty.Text("Part left"),
					),
				),
				elem.Div(vecty.Markup(vecty.Class("tile")),
					elem.Div(vecty.Markup(vecty.Class("tile", "is-child", "box")),
						vecty.Text("Part left 2"),
					),
				),
			),
			elem.Div(vecty.Markup(vecty.Class("tile", "is-parent")),
				elem.Div(vecty.Markup(vecty.Class("tile", "is-child", "box")),
					vecty.Text("Part right"),
				),
			),
		),
	)
}
