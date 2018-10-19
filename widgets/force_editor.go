package widgets

import (
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// DefaultForcesConfig specifies default configuration for physics simulation.
var DefaultForcesConfig = ForcesConfig{
	Config: layout.DefaultConfig,
	Steps:  10,
}

// ForceEditor represents forces and physics simulation configuration widget.
type ForceEditor struct {
	vecty.Core

	config ForcesConfig

	repelling *ForceInput
	spring    *ForceInput
	drag      *ForceInput
	steps     *Range
}

// ForcesConfig represents physics simulation configuration.
type ForcesConfig struct {
	layout.Config
	Steps int
}

// Render implements vecty's Component interface for ForceEditor.
func (l *ForceEditor) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading3(
			vecty.Text("Layout forces:"),
		),
		elem.Form(
			vecty.Markup(
				vecty.Class("pure-form"),
			),
			l.repelling,
			l.spring,
			l.drag,
			l.steps,
		),
	)
}

// NewForceEditor creates a new ForceEditor widget.
func NewForceEditor() *ForceEditor {
	config := DefaultForcesConfig
	repelling := NewForceInput("Gravity force:", config.Repelling)
	spring := NewForceInput("Spring force:", config.SpringStiffness)
	drag := NewForceInput("Drag force:", config.DragCoeff)
	steps := NewRange("Steps:", config.Steps)
	return &ForceEditor{
		config:    config,
		repelling: repelling,
		spring:    spring,
		drag:      drag,
		steps:     steps,
	}
}

// Config returns current forces & simulation configuration, getting values from UI.
func (l *ForceEditor) Config() ForcesConfig {
	return ForcesConfig{
		Steps: l.steps.Value(),
		Config: layout.Config{
			Repelling:       l.repelling.Value(),
			SpringStiffness: l.spring.Value(),
			SpringLen:       l.config.SpringLen,
			DragCoeff:       l.drag.Value(),
		},
	}
}
