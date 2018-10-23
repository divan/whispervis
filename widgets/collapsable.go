package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// Collapsable represents widgets that has open and closed state.
type Collapsable struct {
	vecty.Core
	elems vecty.List

	title  string
	isOpen bool

	button func() vecty.ComponentOrHTML //FIXME: hack against "vecty: next child render must not equal previous child render " panic
}

// NewCollapsable creates new collapsable widget with a init state.
func NewCollapsable(title string, isOpen bool, button func() vecty.ComponentOrHTML, elems ...vecty.ComponentOrHTML) *Collapsable {
	return &Collapsable{
		title:  title,
		isOpen: isOpen,
		elems:  vecty.List(elems),
		button: button, // FIXME: hack
	}
}

// Render implements vecty's Component interface for Collapsable.
func (c *Collapsable) Render() vecty.ComponentOrHTML {
	btn := elem.Button(
		vecty.Markup(
			vecty.Class("button", "is-small", "outlined", "is-pulled-right", "is-clearfix"),
			event.Click(c.onClick),
		),
		vecty.If(c.isOpen, vecty.Text("-")),
		vecty.If(!c.isOpen, vecty.Text("+")),
	)

	return elem.Div(
		btn,
		Header(c.title),
		elem.Div(
			vecty.Markup(
				vecty.MarkupIf(!c.isOpen,
					vecty.Class("is-invisible"),
					vecty.Style("height", "0px"),
				),
			),
			elem.Break(),
			c.elems,
			c.button(), // FIXME: hack
		),
	)
}

func (c *Collapsable) onClick(e *vecty.Event) {
	c.isOpen = !c.isOpen
	vecty.Rerender(c)
}
