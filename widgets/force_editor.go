package widgets

import (
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type ForceEditor struct {
	vecty.Core

	config layout.Config

	repelling *ForceInput
	spring    *ForceInput
	drag      *ForceInput
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
			l.repelling,
			l.spring,
			l.drag,
		),
	)
}

func NewForceEditor() *ForceEditor {
	config := layout.DefaultConfig
	repelling := NewForceInput("Gravity force:", config.Repelling)
	spring := NewForceInput("Spring force:", config.SpringStiffness)
	drag := NewForceInput("Drag force:", config.DragCoeff)
	return &ForceEditor{
		config:    config,
		repelling: repelling,
		spring:    spring,
		drag:      drag,
	}
}

func (l *ForceEditor) Config() layout.Config {
	return layout.Config{
		Repelling:       l.repelling.Value(),
		SpringStiffness: l.spring.Value(),
		SpringLen:       l.config.SpringLen,
		DragCoeff:       l.drag.Value(),
	}
}
