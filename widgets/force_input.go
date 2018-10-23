package widgets

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// ForceInput represents input widget for forces.
type ForceInput struct {
	vecty.Core

	changed     bool
	title       string
	description string
	value       float64
}

// NewForceInput creates new input.
func NewForceInput(title, description string, value float64) *ForceInput {
	f := &ForceInput{
		title:       title,
		description: description,
		value:       value,
	}
	return f
}

// Render implements vecty.Component interface for ForceInput.
func (f *ForceInput) Render() vecty.ComponentOrHTML {
	value := fmt.Sprintf("%.4f", f.value)
	return elem.Div(
		vecty.Markup(
			vecty.Class("field", "is-horizontal", "is-paddingless", "is-marginless"),
		),

		elem.Div(
			vecty.Markup(
				vecty.Class("field-label", "is-normal"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("label"),
				),
				vecty.Text(f.title),
			),
		),

		fieldControl(
			elem.Input(
				vecty.Markup(
					vecty.Class("input", "is-small"),
					vecty.Style("text-align", "right"),
					prop.Value(value),
					event.Input(f.onEditInput),
				),
			),
		),
		vecty.If(f.description != "", QuestionTooltip(f.description)),
	)
}

// helper for wrapping many divs
func fieldControl(element *vecty.HTML) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field-body"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("control"),
				),

				element,
			),
		),
	)
}

func (f *ForceInput) onEditInput(event *vecty.Event) {
	value := event.Target.Get("value").Float()

	f.changed = true
	f.value = value
}

// Value returns the current value.
func (f *ForceInput) Value() float64 {
	return f.value
}

// Changed returns if input value has been changed, and resets it's value to false.
func (f *ForceInput) Changed() bool {
	if f.changed {
		f.changed = false
		return true
	}

	return false
}
