package widgets

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// Range represents <input type="range..> widget.
type Range struct {
	vecty.Core

	min, max    int
	changed     bool
	title       string
	description string
	value       int

	handler func(int)
}

// NewRange builds new Range widget.
func NewRange(title, description string, value, min, max int, handler func(int)) *Range {
	return &Range{
		title:       title,
		description: description,
		value:       value,
		handler:     handler,
		min:         min,
		max:         max,
	}
}

func (r *Range) Render() vecty.ComponentOrHTML {
	value := fmt.Sprintf("%d", r.value)
	return elem.Div(
		vecty.Markup(
			vecty.Class("field", "is-horizontal", "is-paddingless", "is-marginless"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field-label"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("label"),
				),
				vecty.Text(r.title),
			),
		),

		fieldControl(
			elem.Input(
				vecty.Markup(
					vecty.Class("is-small"),
					vecty.Style("margin-right", "10px"),
					prop.Value(value),
					prop.Type("range"),
					vecty.Attribute("min", fmt.Sprintf("%d", r.min)),
					vecty.Attribute("max", fmt.Sprintf("%d", r.max)),
					event.Input(r.onChange),
				),
			),
		),
		fieldControl(
			elem.Input(
				vecty.Markup(
					prop.Value(value),
					event.Input(r.onChange),
					vecty.Class("input", "is-small"),
				),
			),
		),
		vecty.If(r.description != "", QuestionTooltip(r.description)),
	)
}

// Value returns the current value.
func (r *Range) Value() int {
	return r.value
}

// Changed returns if range value has been changed, and resets it's value to false.
func (r *Range) Changed() bool {
	if r.changed {
		r.changed = false
		return true
	}

	return false
}

func (r *Range) onChange(event *vecty.Event) {
	value := event.Target.Get("value").Int()

	r.SetValue(value)
}

func (r *Range) Inc() {
	r.SetValue(r.value + 1)
}

func (r *Range) Dec() {
	r.SetValue(r.value - 1)
}

// SetValue sets a new value for the range.
func (r *Range) SetValue(value int) {
	if value > r.max {
		value = r.max
	}
	if value < r.min {
		value = r.min
	}

	r.value = value
	vecty.Rerender(r)

	r.changed = true

	if r.handler != nil {
		r.handler(r.value)
	}
}
