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
	Steps:  100,
}

// ForceEditor represents forces and physics simulation configuration widget.
type ForceEditor struct {
	vecty.Core

	config ForcesConfig

	repelling   *ForceInput
	spring      *ForceInput
	springLen   *ForceInput
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
	f.repelling = NewForceInput("Gravity:", RepellingForceDescription, f.config.Repelling)
	f.spring = NewForceInput("Spring:", SpringForceDescription, f.config.SpringStiffness)
	f.springLen = NewForceInput("Length:", SpringLengthDescription, f.config.SpringLen)
	f.drag = NewForceInput("Drag:", DragForceDescritption, f.config.DragCoeff)
	f.steps = NewRange("Steps:", StepsDescription, f.config.Steps)
	f.collapsable = NewCollapsable("Layout forces:", false,
		f.applyButton,
		f.repelling,
		f.spring,
		f.springLen,
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
			SpringLen:       f.springLen.Value(),
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

const (
	RepellingForceDescription = "Repelling force coefficient defines the force for nodes to repel from each other. It's a coefficent for Coulomb's law, and should be negative (positive is attraction)"
	SpringForceDescription    = "Spring force coefficient defines the force of attraction between linked nodes. It obeys Hooke's law. The larger the coefficient, the more stiffer the spring."
	SpringLengthDescription   = "Spring length defines a \"normal\" length of the spring. If link is shorted then this value, nodes start to repel, attract otherwise."
	DragForceDescritption     = "Drag force coefficient defines a drag force, that slows down nodes velocities after applying repelling and spring forces. Increase it if you see jitterness or too much movement on each iteration."
	StepsDescription          = "Number of physics simulation steps to run. Too big value may slowdown calculation without giving more benefit to the layout. Too little may not be enough to fully apply forces."
)
