package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Header renders common header for widgets.
func Header(title string) *vecty.HTML {
	return elem.Heading4(
		vecty.Markup(
			vecty.Class("subtitle", "has-text-weight-light", "is-marginless"),
		),
		vecty.Text(title),
	)
}
