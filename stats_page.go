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
			vecty.Class("title", "has-text-centered"),
		),
		elem.Heading1(vecty.Text("Stats page")),
	)
}
