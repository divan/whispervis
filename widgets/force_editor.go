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

	repelling   *ForceInput
	spring      *ForceInput
	drag        *ForceInput
	steps       *Range
	collapsable *Collapsable

	apply func()
}

// ForcesConfig represents physics simulation configuration.
type ForcesConfig struct {
	layout.Config
	Steps int
}

// NewForceEditor creates a new ForceEditor widget.
func NewForceEditor(apply func()) *ForceEditor {
	f := &ForceEditor{}
	f.config = DefaultForcesConfig
	f.repelling = NewForceInput("Gravity:", f.config.Repelling)
	f.spring = NewForceInput("Spring:", f.config.SpringStiffness)
	f.drag = NewForceInput("Drag:", f.config.DragCoeff)
	f.steps = NewRange("Steps:", f.config.Steps)
	f.collapsable = NewCollapsable("Layout forces:", true,
		f.applyButton,
		f.repelling,
		f.spring,
		f.drag,
		f.steps,
	)
	f.apply = apply
	return f
}

// Render implements vecty's Component interface for ForceEditor.
func (f *ForceEditor) Render() vecty.ComponentOrHTML {
	return Widget(
		f.collapsable,
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

func (f *ForceEditor) applyButton() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("button", "is-info", "is-small"),
			event.Click(f.onClick),
		),
		vecty.Text("Apply"),
	)
}

func (f *ForceEditor) onClick(e *vecty.Event) {
	go f.apply()
}
