package widgets

import (
	"github.com/divan/graphx/layout"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
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

	apply func()
}

// ForcesConfig represents physics simulation configuration.
type ForcesConfig struct {
	layout.Config
	Steps int
}

// NewForceEditor creates a new ForceEditor widget.
func NewForceEditor(apply func()) *ForceEditor {
	config := DefaultForcesConfig
	repelling := NewForceInput("Gravity:", config.Repelling)
	spring := NewForceInput("Spring:", config.SpringStiffness)
	drag := NewForceInput("Drag:", config.DragCoeff)
	steps := NewRange("Steps:", config.Steps)
	return &ForceEditor{
		config:    config,
		repelling: repelling,
		spring:    spring,
		drag:      drag,
		steps:     steps,
		apply:     apply,
	}
}

// Render implements vecty's Component interface for ForceEditor.
func (f *ForceEditor) Render() vecty.ComponentOrHTML {
	return Widget(
		Header("Layout forces:"),
		elem.Form(
			f.repelling,
			f.spring,
			f.drag,
			f.steps,
		),
		f.applyButton(),
	)
}

// Config returns current forces & simulation configuration, getting values from UI.
func (f *ForceEditor) Config() ForcesConfig {
	return ForcesConfig{
		Steps: f.steps.Value(),
		Config: layout.Config{
			Repelling:       f.repelling.Value(),
			SpringStiffness: f.spring.Value(),
			SpringLen:       f.config.SpringLen,
			DragCoeff:       f.drag.Value(),
		},
	}
}

func (f *ForceEditor) applyButton() *vecty.HTML {
	return elem.Div(
		elem.Button(
			vecty.Markup(
				vecty.Class("button", "is-info", "is-small"),
				event.Click(f.onClick),
			),
			vecty.Text("Apply"),
		),
	)
}

func (f *ForceEditor) onClick(e *vecty.Event) {
	go f.apply()
}
