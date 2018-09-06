package widgets

import (
	"fmt"
	"strconv"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type ForceInput struct {
	vecty.Core

	changed bool
	title   string
	value   float64
}

func NewForceInput(title string, value float64) *ForceInput {
	return &ForceInput{
		title: title,
		value: value,
	}
}

func (f *ForceInput) Render() vecty.ComponentOrHTML {
	value := fmt.Sprintf("%.2f", f.value)
	return elem.Div(
		vecty.Markup(
			vecty.Class("pure-markup-group", "pure-u-1"),
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("pure-u-1-2"),
			),
			vecty.Text(f.title),
		),
		elem.Input(
			vecty.Markup(
				prop.Value(value),
				event.Input(f.onEditInput),
				vecty.Class("pure-input-1-3"),
				vecty.Style("float", "right"),
				vecty.Style("margin-right", "10px"),
				vecty.Style("text-align", "right"),
			),
		),
	)
}

func (f *ForceInput) onEditInput(event *vecty.Event) {
	value := event.Target.Get("value").String()
	fvalue, err := strconv.ParseFloat(value, 0)
	if err != nil {
		return
	}

	f.changed = true
	f.value = fvalue
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
