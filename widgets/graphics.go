package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// Graphics represents graphics configuration widget.
type Graphics struct {
	vecty.Core

	rtSwitch    *Switch
	collapsable *Collapsable
}

// NewGraphics creates a new Graphics widget.
func NewGraphics() *Graphics {
	g := &Graphics{}
	g.rtSwitch = NewSwitch(true)
	g.collapsable = NewCollapsable("Graphics:", false,
		g.applyButton,
		g.rtSwitch,
	)
	return g
}

// Render implements vecty's Component interface for Graphics.
func (g *Graphics) Render() vecty.ComponentOrHTML {
	return Widget(
		g.collapsable,
	)
}

func (g *Graphics) applyButton() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("button", "is-info", "is-small"),
			event.Click(g.onApply),
		),
		vecty.Text("Apply"),
	)
}

func (g *Graphics) onApply(e *vecty.Event) {
	// TODO
}
