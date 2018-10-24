package widgets

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// Radio is a wrapper around checkbox-like radio.
type Radio struct {
	vecty.Core

	title     string
	groupName string // name="" for radiogroup
	isChecked bool
	handler   func()

	domID string
}

// NewRadio creates and inits a new radio.
func NewRadio(title, groupName string, checked bool, handler func()) *Radio {
	rnd := rand.Int63n(math.MaxInt64)
	domID := fmt.Sprintf("idRadio%d", rnd)
	return &Radio{
		title:     title,
		isChecked: checked,
		handler:   handler,
		domID:     domID,
		groupName: groupName,
	}
}

// Render implements vecty's Component interface for Radio.
func (r *Radio) Render() vecty.ComponentOrHTML {
	return elem.Span(
		elem.Input(
			vecty.Markup(
				vecty.Class("is-checkradio", "is-rounded", "is-small", "is-horizontal"),
				prop.ID(r.domID),
				vecty.Attribute("name", r.groupName),
				prop.Type(prop.TypeRadio),
				event.Change(r.onToggle),
				vecty.MarkupIf(r.isChecked,
					vecty.Attribute("checked", "checked"),
				),
			),
		),
		elem.Label(
			vecty.Markup(
				// this is needed to properly handle click
				// see https://wikiki.github.io/form/checkradio/
				vecty.Attribute("for", r.domID),
			),
			vecty.Text(r.title),
		),
	)
}

// Checked returns the checked state of the radio.
func (r *Radio) Checked() bool {
	return r.isChecked
}

// onToggle changes the checked state of the radio.
func (r *Radio) onToggle(*vecty.Event) {
	r.isChecked = !r.isChecked
	vecty.Rerender(r)
	if r.handler != nil {
		go r.handler()
	}
}
