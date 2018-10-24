package widgets

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// RadioGroup is a group of checkradio widgets.
// FIXME: currently using the same values and text as ints, because
// the only use case of the widget is FPS setting. Change it to string
// and add proper separation when needed.
type RadioGroup struct {
	vecty.Core

	title   string
	value   int
	handler func(int)

	radios []*Radio

	domID string
}

// NewRadioGroup creates and inits a new radio.
func NewRadioGroup(title string, value int, handler func(int), values []int) *RadioGroup {
	rnd := rand.Int63n(math.MaxInt64)
	domID := fmt.Sprintf("idRadioGroup%d", rnd)

	rg := &RadioGroup{
		title:   title,
		handler: handler,
		domID:   domID,
		radios:  make([]*Radio, len(values)),
	}

	// create radios
	for i, val := range values {
		str := fmt.Sprintf("%d", val)
		checked := value == val
		radio := NewRadio(str, title, checked, rg.onChange(val))
		rg.radios[i] = radio
	}

	return rg
}

// Render implements vecty's Component interface for RadioGroup.
func (s *RadioGroup) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field"),
		),
		elem.Span(
			vecty.Text(s.title),
		),
		s.renderRadios(),
	)
}

func (s *RadioGroup) renderRadios() *vecty.HTML {
	elems := []vecty.MarkupOrChild{}
	for i := range s.radios {
		elems = append(elems, s.radios[i])
	}
	return elem.Div(
		elems...,
	)
}

func (s *RadioGroup) onChange(value int) func() {
	return func() {
		s.value = value
		if s.handler != nil {
			go s.handler(value)
		}
	}
}
