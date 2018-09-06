package widgets

import (
	"fmt"

	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type ForceEditor struct {
	vecty.Core

	inputs vecty.List
	config layout.Config
}

func (l *ForceEditor) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading3(
			vecty.Text("Layout forces:"),
		),
		elem.HorizontalRule(),
		elem.Form(
			vecty.Markup(
				vecty.Class("pure-form"),
			),
			vecty.List(l.inputs),
		),
		elem.Div(
			elem.Button(
				vecty.Markup(
					vecty.Class("pure-button"),
					vecty.Style("background", "rgb(28, 184, 65)"),
					vecty.Style("color", "white"),
					vecty.Style("border-radius", "4px"),
					event.Click(l.onUpdateClick),
				),
				vecty.Text("Update"),
			),
		),
	)
}

func NewForceEditor() *ForceEditor {
	config := layout.DefaultConfig
	inputs := vecty.List{
		NewForceInput("Gravity force:", config.Repelling),
		NewForceInput("Spring force:", config.SpringStiffness),
		NewForceInput("Drag force:", config.DragCoeff),
	}
	return &ForceEditor{
		inputs: inputs,
		config: config,
	}
}

func (l *ForceEditor) Config() layout.Config {
	return l.config
}

func (l *ForceEditor) onUpdateClick(e *vecty.Event) {
	fmt.Println("Clicked")
}
