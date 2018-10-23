package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Widgets renders common sidebar widget layout.
func Widget(elements ...vecty.MarkupOrChild) *vecty.HTML {
	hr := elem.HorizontalRule(
		vecty.Markup(
			vecty.Class("has-background-grey-light"),
			vecty.Style("margin", "2px 0 10px 0"),
		),
	)
	markup := vecty.Markup(
		vecty.Class("box", "is-shadowless", "is-marginless", "has-background-light"),
		vecty.Style("padding", "1px"),
	)
	elems := []vecty.MarkupOrChild{markup}
	elems = append(elems, elements...)
	return elem.Div(
		hr,
		elem.Div(
			elems...,
		),
	)
}

// Header renders common header for widgets.
func Header(title string) *vecty.HTML {
	return elem.Heading4(
		vecty.Markup(
			vecty.Class("subtitle", "has-text-weight-light", "is-marginless"),
		),
		vecty.Text(title),
	)
}

// QuestionTooltip renders question mark with tooltip with given text.
func QuestionTooltip(description string) *vecty.HTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("is-small", "tooltip", "is-tooltip-multiline"),
			vecty.Attribute("data-tooltip", description),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("icon", "is-small", "has-text-grey-light"),
			),
			elem.Italic(
				vecty.Markup(
					vecty.Class("fas", "fa-question-circle"),
				),
			),
		),
	)
}
