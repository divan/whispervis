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

	conf SceneConfigurator
}

// NewGraphics creates a new Graphics widget. It needs to have
// access to scene configuration, as it configures mostly things from it.
func NewGraphics(conf SceneConfigurator) *Graphics {
	g := &Graphics{
		conf: conf,
	}
	g.rtSwitch = NewSwitch("Render throttler", true, conf.ToggleRenderThrottler)
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
