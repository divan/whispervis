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

	changed bool
	title   string
	value   int
}

// NewRange builds new Range widget.
func NewRange(title string, value int) *Range {
	return &Range{
		title: title,
		value: value,
	}
}

func (r *Range) Render() vecty.ComponentOrHTML {
	value := fmt.Sprintf("%d", r.value)
	return elem.Div(
		vecty.Markup(
			vecty.Class("pure-markup-group", "pure-u-1"),
			vecty.Style("height", "32px"),
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("pure-u-1-3"),
			),
			vecty.Text(r.title),
		),
		elem.Input(
			vecty.Markup(
				prop.Value(value),
				prop.Type("range"),
				vecty.Attribute("min", "1"),
				vecty.Attribute("max", "1000"), // allow 1-1000 range for steps
				event.Input(r.onChange),
				vecty.Class("pure-input-1-3"),
				vecty.Style("height", "100%"),
			),
		),
		elem.Input(
			vecty.Markup(
				prop.Value(value),
				event.Input(r.onChange),
				vecty.Class("pure-input-1-4"),
				vecty.Style("float", "right"),
				vecty.Style("margin-right", "10px"),
				vecty.Style("text-align", "right"),
			),
		),
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
}
