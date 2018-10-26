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

	changed     bool
	title       string
	description string
	value       int

	handler func(int)
}

// NewRange builds new Range widget.
func NewRange(title, description string, value int, handler func(int)) *Range {
	return &Range{
		title:       title,
		description: description,
		value:       value,
		handler:     handler,
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
					vecty.Attribute("min", "1"),
					vecty.Attribute("max", "1000"), // allow 1-1000 range for steps
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

	r.changed = true
	r.value = value
	vecty.Rerender(r)

	if r.handler != nil {
		r.handler(value)
	}
}
