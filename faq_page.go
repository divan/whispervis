package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// FAQPage is FAQ/docs view component.
type FAQPage struct {
	vecty.Core

	width, height string
}

// NewFAQPage creates and inits new stats page.
func NewFAQPage() *FAQPage {
	width, height := PageViewSize()
	return &FAQPage{
		width:  fmt.Sprintf("%dpx", width),
		height: fmt.Sprintf("%dpx", height),
	}
}

// Render implements the vecty.Component interface.
func (s *FAQPage) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("width", s.width),
			vecty.Style("height", s.height),
			vecty.Class("title", "has-text-centered"),
		),
		elem.Heading1(vecty.Text("FAQ page")),
	)
}
